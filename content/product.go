package content

import (
	"fmt"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Product struct {
	item.Item

	Name        string   `json:"name"`
	Price       float64  `json:"price"`
	Description string   `json:"description"`
	Brand       string   `json:"brand"`
	Ingredients []string `json:"ingredients"`
	Retailers   []string `json:"retailers"`
}

// MarshalEditor writes a buffer of html to edit a Product within the CMS
// and implements editor.Editable
func (p *Product) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Product field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", p, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.Input("Price", p, map[string]string{
				"label":       "Price",
				"type":        "text",
				"placeholder": "Enter the Price here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Description", p, map[string]string{
				"label":       "Description",
				"placeholder": "Enter the Description here",
			}),
		},
		editor.Field{
			View: reference.Select("Brand", p, map[string]string{
				"label": "Brand",
			},
				"Brand",
				`{{ .name }}`,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Ingredients", p, map[string]string{
				"label": "Ingredients",
			},
				"Ingredient",
				`{{ .name }}`,
			),
		},
		editor.Field{
			View: reference.SelectRepeater("Retailers", p, map[string]string{
				"label": "Retailers",
			},
				"Retailer",
				`{{ .name }}`,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Product editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Product"] = func() interface{} { return new(Product) }
}

// String defines how a Product is printed. Update it using more descriptive
// fields from the Product struct type
func (p *Product) String() string {
	return fmt.Sprintf("%s", p.Name)
}
