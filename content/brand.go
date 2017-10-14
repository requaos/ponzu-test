package content

import (
	"fmt"

	"github.com/bosssauce/reference"
	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Brand struct {
	item.Item

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Established int      `json:"established"`
	Products    []string `json:"products"`
	Logo        string   `json:"logo"`
}

// MarshalEditor writes a buffer of html to edit a Brand within the CMS
// and implements editor.Editable
func (b *Brand) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(b,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Brand field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", b, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", b, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.Input("Established", b, map[string]string{
				"label":       "Established",
				"type":        "text",
				"placeholder": "Enter the Established here",
			}),
		},
		editor.Field{
			View: editor.File("Logo", b, map[string]string{
				"label":       "Logo",
				"placeholder": "Upload the Logo here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Products", b, map[string]string{
				"label": "Products",
			},
				"Product",
				`{{ .name }}: {{ .price }}`,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Brand editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Brand"] = func() interface{} { return new(Brand) }
}

// String defines how a Brand is printed. Update it using more descriptive
// fields from the Brand struct type
func (b *Brand) String() string {
	return fmt.Sprintf("%s", b.Name)
}
