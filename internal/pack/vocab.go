package pack

// The vocabularies below are the frozen part of the pack format. Every pack
// defines exactly these tokens and shows exactly these components, which is
// what lets mix mode swap a palette, type or surface part between styles.
// docs/pack-format.md documents them; a test keeps the two in step.

var Archetypes = []string{"forum", "blog", "portal", "shop", "landing", "app", "docs"}

// Part is one swappable section of tokens.css.
type Part struct {
	Name   string
	Tokens []string
}

var Parts = []Part{
	{"palette", []string{
		"--color-page",
		"--color-canvas",
		"--color-surface",
		"--color-surface-alt",
		"--color-surface-strong",
		"--color-bar",
		"--color-bar-text",
		"--color-bar-alt",
		"--color-bar-alt-text",
		"--color-inverse",
		"--color-inverse-text",
		"--color-text",
		"--color-text-muted",
		"--color-heading",
		"--color-heading-alt",
		"--color-link",
		"--color-link-quiet",
		"--color-link-visited",
		"--color-link-hover",
		"--color-link-active",
		"--color-border",
		"--color-border-strong",
		"--color-border-muted",
		"--color-accent",
		"--color-accent-text",
		"--color-accent-alt",
		"--color-fill-1",
		"--color-fill-2",
		"--color-fill-3",
		"--color-fill-4",
		"--color-button",
		"--color-button-text",
		"--color-button-hover-text",
		"--color-button-secondary",
		"--color-button-secondary-text",
		"--color-input",
		"--color-input-text",
		"--color-input-border",
		"--color-disabled",
		"--color-disabled-text",
		"--color-danger",
		"--color-danger-surface",
		"--color-success",
		"--color-warning",
		"--color-notice",
		"--color-notice-text",
		"--color-focus",
		"--color-shadow",
		"--color-overlay",
	}},
	{"type", []string{
		"--font-body",
		"--font-heading",
		"--font-ui",
		"--font-mono",
		"--text-base",
		"--text-small",
		"--text-ui",
		"--text-large",
		"--text-display",
		"--text-h1",
		"--text-h2",
		"--text-h3",
		"--line-body",
		"--line-heading",
		"--line-display",
		"--weight-body",
		"--weight-bold",
		"--weight-heading",
		"--weight-display",
		"--weight-ui",
		"--heading-transform",
		"--heading-tracking",
		"--display-tracking",
		"--figure-shift",
		"--ui-transform",
		"--link-decoration",
		"--link-decoration-hover",
		"--link-decoration-quiet",
	}},
	{"surface", []string{
		"--border-width",
		"--border-width-strong",
		"--border-style",
		"--radius-control",
		"--radius-panel",
		"--radius-pill",
		"--radius-page",
		"--shadow-panel",
		"--shadow-control",
		"--shadow-control-hover",
		"--shadow-control-pressed",
		"--shadow-dialog",
		"--shadow-text",
		"--fill-page",
		"--fill-bar",
		"--fill-bar-alt",
		"--fill-inverse",
		"--fill-accent",
		"--fill-panel",
		"--fill-button",
		"--fill-button-hover",
		"--fill-button-secondary",
		"--fill-input",
		"--focus-ring",
		"--transition",
		"--backdrop-blur",
	}},
}

// Component is one entry of the specimen checklist.
type Component struct {
	ID    string // value of data-component in specimen.html
	Class string // base class in components.css
}

var Components = []Component{
	{"shell", "ds-page"},
	{"brand", "ds-brand"},
	{"nav", "ds-nav"},
	{"hero", "ds-hero"},
	{"page-header", "ds-page-header"},
	{"typography", "ds-prose"},
	{"links", "ds-link"},
	{"buttons", "ds-button"},
	{"form", "ds-form"},
	{"table", "ds-table"},
	{"list", "ds-list"},
	{"panel", "ds-panel"},
	{"stats", "ds-stat"},
	{"grid", "ds-grid"},
	{"tabs", "ds-tabs"},
	{"badge", "ds-badge"},
	{"sidebar", "ds-sidebar"},
	{"notice", "ds-notice"},
	{"pagination", "ds-pagination"},
	{"breadcrumb", "ds-breadcrumb"},
	{"dialog", "ds-dialog"},
	{"empty", "ds-empty"},
	{"footer", "ds-footer"},
}

// Variants are classes every pack defines and shows besides the base
// classes: the states an application needs whatever the era.
var Variants = []string{
	"ds-button--secondary",
	"ds-button--danger",
	"ds-badge--success",
	"ds-badge--warning",
	"ds-badge--danger",
}

// StyleSections are the level-2 headings of STYLE.md, in order.
var StyleSections = []string{
	"Summary",
	"Layout",
	"Typography and colour roles",
	"Components",
	"Never",
	"Extending",
}

// PresenceMetrics are the Never metrics that are either there or not: a
// rule on them says "= none" or "!= none".
var PresenceMetrics = []string{"box-shadow", "text-shadow", "transition", "animation"}

// NeverMetrics are the measurements a "Never" rule may constrain. Each one
// is something the extractor can measure on a rendered page.
var NeverMetrics = []string{
	"border-radius",
	"border-width",
	"box-shadow",
	"box-shadow-blur",
	"text-shadow",
	"gradient-fills",
	"font-size",
	"font-weight",
	"font-families",
	"line-height",
	"underlined-links",
	"letter-spacing",
	"uppercase-text",
	"row-gap",
	"block-gap",
	"content-width",
	"palette-colours",
	"transition",
	"animation",
}
