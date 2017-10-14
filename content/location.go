package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Location struct {
	item.Item

	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     int    `json:"zip"`
	Mapurl  string `json:"mapurl"`
	Phone   string `json:"phone"`
}

// MarshalEditor writes a buffer of html to edit a Location within the CMS
// and implements editor.Editable
func (l *Location) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(l,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Location field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Street1", l, map[string]string{
				"label":       "Street Address Line 1",
				"type":        "text",
				"placeholder": "Enter the Street Address here",
			}),
		},
		editor.Field{
			View: editor.Input("Street2", l, map[string]string{
				"label":       "Street Address Line 2",
				"type":        "text",
				"placeholder": "Enter the Street Address here",
			}),
		},
		editor.Field{
			View: editor.Input("City", l, map[string]string{
				"label":       "City",
				"type":        "text",
				"placeholder": "Enter the City here",
			}),
		},
		editor.Field{
			View: editor.Input("State", l, map[string]string{
				"label":       "State",
				"type":        "text",
				"placeholder": "Enter the State here",
			}),
		},
		editor.Field{
			View: editor.Input("Zip", l, map[string]string{
				"label":       "Zip Code",
				"type":        "text",
				"placeholder": "Enter the Zip Code here",
			}),
		},
		editor.Field{
			View: editor.Input("Mapurl", l, map[string]string{
				"label":       "Map URL",
				"type":        "text",
				"placeholder": "Enter the Map URL here",
			}),
		},
		editor.Field{
			View: editor.Input("Phone", l, map[string]string{
				"label":       "Phone Number",
				"type":        "text",
				"placeholder": "Enter the Phone Number here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Location editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Location"] = func() interface{} { return new(Location) }
}

// String defines how a Location is printed. Update it using more descriptive
// fields from the Location struct type
func (l *Location) String() string {
	return fmt.Sprintf("%s %s %s, %s %d", l.Street1, l.Street2, l.City, l.State, l.Zip)
}
