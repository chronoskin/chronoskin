package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/chronoskin/chronoskin/internal/library"
	"github.com/chronoskin/chronoskin/internal/pack"
)

// MCP is hand-written JSON-RPC 2.0: initialize, tools/list, tools/call and
// ping. The same dispatcher serves POST /mcp and stdio.

// mcpVersions, newest first; a client that asks for another gets the newest.
var mcpVersions = []string{"2025-06-18", "2025-03-26", "2024-11-05"}

// JSON-RPC 2.0 error codes.
const (
	rpcParseError     = -32700
	rpcMethodNotFound = -32601
	rpcInvalidParams  = -32602
)

const (
	// A message carries the CSS for lint_css.
	maxRPCBytes = 1 << 20
	// A line on stdin starts with a buffer of stdioLineBytes, which grows
	// to maxStdioLineBytes at most.
	stdioLineBytes    = 64 << 10
	maxStdioLineBytes = 4 << 20
)

type rpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type rpcResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      json.RawMessage `json:"id"`
	Result  any             `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type obj = map[string]any

type tool struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema obj    `json:"inputSchema"`
}

func str(description string) obj {
	return obj{"type": "string", "description": description}
}

var tools = []tool{
	{"list_eras", "List the eras of web design that styles can be generated for, with their years, archetypes and descriptions.",
		obj{"type": "object", "properties": obj{}}},
	{"generate_style", "Pick a style and return its Style ID, a summary and a preview URL. Call again with a different seed for another one. To regenerate from an existing style while keeping some of its parts, pass its ID as `from` and the parts to keep as `lock`.",
		obj{"type": "object", "properties": obj{
			"era":       str("Era slug from list_eras, or \"any\" (default)."),
			"archetype": str("Page type: " + strings.Join(pack.Archetypes, ", ") + ". Optional."),
			"mode":      obj{"type": "string", "enum": []string{"pure", "mix"}, "description": "pure returns one synthesized style; mix combines the structure of one with the palette, type and surface of its siblings. Default pure."},
			"seed":      obj{"type": "integer", "minimum": 0, "description": "The same seed and inputs give the same style. Random when omitted."},
			"from":      str("Style ID to regenerate from."),
			"lock":      obj{"type": "array", "items": obj{"type": "string", "enum": []string{"structure", "palette", "type", "surface"}}, "description": "Parts of `from` to keep."},
			"density":   obj{"type": "string", "enum": []string{"compact", "normal", "roomy"}, "description": "Scale the layout's spacing. Default: the layout's own, or what `from` has."},
			"colours":   obj{"type": "object", "additionalProperties": obj{"type": "string"}, "description": "Palette colours to set by hand, by token name, such as {\"--color-accent\": \"#2b55e0\"}. Values are #rrggbb or rgba(). Pass with `from` and all four parts locked to recolour the current style. Optional."},
		}}},
	{"get_style", "Return the five files of a style pack. Write them unchanged into .design/ at the repository root.",
		obj{"type": "object", "required": []string{"id"}, "properties": obj{"id": str("Style ID.")}}},
	{"lint_css", "Check CSS against a style pack. Returns every colour, font, size, radius and shadow that does not come from the pack's tokens.",
		obj{"type": "object", "required": []string{"id", "css"}, "properties": obj{
			"id":  str("Style ID, from .design/style.json."),
			"css": str("The CSS to check."),
		}}},
}

// rpc handles one message. It returns nil for notifications.
func (s *server) rpc(req rpcRequest, base string) *rpcResponse {
	if req.ID == nil {
		return nil
	}
	resp := &rpcResponse{JSONRPC: "2.0", ID: req.ID}
	switch req.Method {
	case "initialize":
		var p struct {
			ProtocolVersion string `json:"protocolVersion"`
		}
		json.Unmarshal(req.Params, &p)
		version := mcpVersions[0]
		if slices.Contains(mcpVersions, p.ProtocolVersion) {
			version = p.ProtocolVersion
		}
		resp.Result = obj{
			"protocolVersion": version,
			"capabilities":    obj{"tools": obj{}},
			"serverInfo":      obj{"name": "chronoskin", "version": s.lib.Version},
			"instructions":    "Style packs that make UI follow the look of a chosen era of web design. Generate a style, write its files into .design/, read .design/STYLE.md and .design/specimen.html before any UI work, and check new CSS with lint_css.",
		}
	case "ping":
		resp.Result = obj{}
	case "tools/list":
		resp.Result = obj{"tools": tools}
	case "tools/call":
		var p struct {
			Name      string          `json:"name"`
			Arguments json.RawMessage `json:"arguments"`
		}
		if err := json.Unmarshal(req.Params, &p); err != nil {
			resp.Error = &rpcError{rpcInvalidParams, "invalid params"}
			break
		}
		content, err := s.callTool(p.Name, p.Arguments, base)
		if err == errUnknownTool {
			resp.Error = &rpcError{rpcInvalidParams, "unknown tool: " + p.Name}
			break
		}
		if err != nil {
			resp.Result = obj{"content": []obj{{"type": "text", "text": err.Error()}}, "isError": true}
			break
		}
		blocks := make([]obj, len(content))
		for i, text := range content {
			blocks[i] = obj{"type": "text", "text": text}
		}
		resp.Result = obj{"content": blocks}
	default:
		resp.Error = &rpcError{rpcMethodNotFound, "method not found: " + req.Method}
	}
	return resp
}

var errUnknownTool = errors.New("unknown tool")

// toolArgs is the arguments of every tool together; each reads its own.
type toolArgs struct {
	Era       string            `json:"era"`
	Archetype string            `json:"archetype"`
	Mode      string            `json:"mode"`
	Seed      *uint64           `json:"seed"`
	From      string            `json:"from"`
	Lock      []string          `json:"lock"`
	ID        string            `json:"id"`
	Density   string            `json:"density"`
	Colours   map[string]string `json:"colours"`
	CSS       string            `json:"css"`
}

func (s *server) callTool(name string, raw json.RawMessage, base string) ([]string, error) {
	var args toolArgs
	if len(raw) > 0 {
		if err := json.Unmarshal(raw, &args); err != nil {
			return nil, fmt.Errorf("invalid arguments: %v", err)
		}
	}
	switch name {
	case "list_eras":
		return s.listEras()
	case "generate_style":
		return s.generateStyle(args, base)
	case "get_style":
		return s.getStyle(args.ID, base)
	case "lint_css":
		return s.lintCSS(args.ID, args.CSS, base)
	}
	return nil, errUnknownTool
}

func jsonText(v any) ([]string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	return []string{string(b)}, err
}

func (s *server) listEras() ([]string, error) {
	type era struct {
		library.Era
		Archetypes []string `json:"archetypes"`
	}
	var eras []era
	for _, e := range s.lib.Eras {
		eras = append(eras, era{e.Era, s.lib.Archetypes(e.Slug)})
	}
	return jsonText(eras)
}

func (s *server) generateStyle(args toolArgs, base string) ([]string, error) {
	req := library.Request{
		Era:       args.Era,
		Archetype: args.Archetype,
		Mode:      args.Mode,
		From:      args.From,
		Lock:      args.Lock,
		Density:   args.Density,
		Colours:   args.Colours,
		Seed:      randomSeed(),
	}
	if args.Seed != nil {
		req.Seed = *args.Seed
	}
	id, err := s.lib.Generate(req)
	if err != nil {
		return nil, err
	}
	return jsonText(s.describe(id, base))
}

func (s *server) packOf(text, base string) (pack.ID, map[string]string, error) {
	id, ok := s.lib.Parse(text)
	if !ok {
		return id, nil, fmt.Errorf("%q is not a Style ID", text)
	}
	files, err := s.packFiles(id, base)
	return id, files, err
}

func (s *server) getStyle(idText, base string) ([]string, error) {
	id, files, err := s.packOf(idText, base)
	if err != nil {
		return nil, err
	}
	out := []string{fmt.Sprintf("Style %s. Write the following %d files unchanged into %s/ at the repository root. Each block below starts with a FILE line that is not part of the file.",
		id, len(pack.Files), pack.Dir)}
	for _, name := range pack.Files {
		out = append(out, "FILE: "+pack.Dir+"/"+name+"\n"+files[name])
	}
	return out, nil
}

func (s *server) lintCSS(idText, css, base string) ([]string, error) {
	_, files, err := s.packOf(idText, base)
	if err != nil {
		return nil, err
	}
	violations, err := pack.LintCSS(css, files["tokens.css"])
	if err != nil {
		return nil, fmt.Errorf("could not parse the CSS: %v", err)
	}
	if len(violations) == 0 {
		return []string{"ok: no violations"}, nil
	}
	return []string{fmt.Sprintf("%d violations:\n- %s", len(violations), strings.Join(violations, "\n- "))}, nil
}

// mcpHTTP is the Streamable HTTP transport, stateless: one JSON-RPC message
// per POST, answered with one JSON body.
func (s *server) mcpHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "this endpoint accepts POST only", http.StatusMethodNotAllowed)
		return
	}
	// A browser page on another site has no business calling this endpoint.
	if origin := r.Header.Get("Origin"); origin != "" && origin != "http://"+r.Host && origin != "https://"+r.Host {
		http.Error(w, "cross-origin requests are not accepted", http.StatusForbidden)
		return
	}
	var req rpcRequest
	body := http.MaxBytesReader(w, r.Body, maxRPCBytes)
	if err := json.NewDecoder(body).Decode(&req); err != nil || req.JSONRPC != "2.0" {
		writeJSON(w, parseFailure("expected one JSON-RPC 2.0 message"))
		return
	}
	resp := s.rpc(req, s.base(r))
	if resp == nil {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	writeJSON(w, resp)
}

// parseFailure answers a message that could not be read, and so has no id.
func parseFailure(message string) rpcResponse {
	return rpcResponse{
		JSONRPC: "2.0",
		ID:      json.RawMessage("null"),
		Error:   &rpcError{rpcParseError, message},
	}
}

// serveStdio is the stdio transport: one JSON-RPC message per line.
func (s *server) serveStdio(in io.Reader, out io.Writer) error {
	base := s.baseURL
	if base == "" {
		base = defaultSite
	}
	scanner := bufio.NewScanner(in)
	scanner.Buffer(make([]byte, stdioLineBytes), maxStdioLineBytes)
	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var req rpcRequest
		if err := json.Unmarshal([]byte(line), &req); err != nil {
			enc.Encode(parseFailure("parse error"))
			continue
		}
		if resp := s.rpc(req, base); resp != nil {
			if err := enc.Encode(resp); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}
