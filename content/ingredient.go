package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Ingredient struct {
	item.Item

	Name        string `json:"name"`
	Description string `json:"description"`
	Molecule    string `json:"molecule"`
}

// MarshalEditor writes a buffer of html to edit a Ingredient within the CMS
// and implements editor.Editable
func (i *Ingredient) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(i,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Ingredient field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", i, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", i, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: editor.File("Molecule", i, map[string]string{
				"label":       "Image",
				"placeholder": "Upload an image here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Ingredient editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Ingredient"] = func() interface{} { return new(Ingredient) }
}

// String defines how a Ingredient is printed. Update it using more descriptive
// fields from the Ingredient struct type
func (i *Ingredient) String() string {
	return fmt.Sprintf("%s", i.Name)
}
