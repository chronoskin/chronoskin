// Package pack defines the style pack: the five files in .design/ that are
// the only contract between the server and the agent.
package pack

// Dir is the directory a pack is written to, at the repository root.
const Dir = ".design"

// In the order they are served and archived.
var Files = []string{
	"style.json",
	"STYLE.md",
	"tokens.css",
	"components.css",
	"specimen.html",
}

// Manifest is style.json.
type Manifest struct {
	ID        string            `json:"id"`
	Library   string            `json:"library"`
	Era       Era               `json:"era"`
	Archetype string            `json:"archetype"`
	Mode      string            `json:"mode"`   // "pure" or "mix"
	Markup    string            `json:"markup"` // "modern" in release 1
	Viewport  Viewport          `json:"viewport"`
	Source    string            `json:"source"`
	Short     string            `json:"short,omitempty"`   // the short form of the ID
	Density   string            `json:"density,omitempty"` // "compact" or "roomy" when the spacing was adjusted
	Colours   map[string]string `json:"colours,omitempty"` // palette tokens set by hand, by name
}

type Era struct {
	Slug  string `json:"slug"`
	Years [2]int `json:"years"`
}

type Viewport struct {
	Width int  `json:"width"`
	Fluid bool `json:"fluid"`
	// Responsive says the layout rearranges itself for a phone. Without it
	// the page has one width, as pages had before phones.
	Responsive bool `json:"responsive,omitempty"`
}
