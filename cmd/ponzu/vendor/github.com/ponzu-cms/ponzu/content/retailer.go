package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Retailer struct {
	item.Item

	Name        string   `json:"name"`
	Description string   `json:"description"`
	Location    []string `json:"location"`
	Images      []string `json:"images"`
}

// MarshalEditor writes a buffer of html to edit a Retailer within the CMS
// and implements editor.Editable
func (r *Retailer) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Retailer field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", r, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", r, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("Location", r, map[string]string{
				"label": "Location",
			},
				"Location",
				`{{ .street1 }} {{ .street2 }} {{ .city }}, {{ .state }} {{ .zip }}`,
			),
		},
		editor.Field{
			View: editor.FileRepeater("Images", r, map[string]string{
				"label":       "Images",
				"placeholder": "Upload the Images here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Retailer editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Retailer"] = func() interface{} { return new(Retailer) }
}

// String defines how a Retailer is printed. Update it using more descriptive
// fields from the Retailer struct type
func (r *Retailer) String() string {
	return fmt.Sprintf("%s", r.Name)
}
