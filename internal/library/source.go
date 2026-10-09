package library

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Source reads a directory of eras as they are written by hand and returns
// what Build takes. Each era is a directory named by its slug:
//
//	<era>/era.json      the era's name, code and description
//	<era>/01, 02, ...   its packs, one layout with its own tokens each
//	<era>/tokens/NN     token sets without a layout (optional)
func Source(root string) (packs []string, opts Options, err error) {
	opts.Eras = map[string]Era{}
	eraFiles, err := filepath.Glob(filepath.Join(root, "*", "era.json"))
	if err != nil {
		return nil, opts, err
	}
	for _, path := range eraFiles {
		dir := filepath.Dir(path)
		raw, err := os.ReadFile(path)
		if err != nil {
			return nil, opts, err
		}
		var era Era
		if err := json.Unmarshal(raw, &era); err != nil {
			return nil, opts, fmt.Errorf("%s: %w", path, err)
		}
		opts.Eras[filepath.Base(dir)] = era
		layouts, _ := filepath.Glob(filepath.Join(dir, "[0-9][0-9]"))
		if len(layouts) == 0 {
			return nil, opts, fmt.Errorf("%s: an era needs at least one pack", dir)
		}
		packs = append(packs, layouts...)
		sets, _ := filepath.Glob(filepath.Join(dir, "tokens", "[0-9][0-9]"))
		opts.Variants = append(opts.Variants, sets...)
	}
	if len(packs) == 0 {
		return nil, opts, fmt.Errorf("%s: no eras found", root)
	}
	return packs, opts, nil
}
