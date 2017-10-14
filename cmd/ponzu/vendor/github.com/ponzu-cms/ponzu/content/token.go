package content

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
	"github.com/requaos/access"
)

type Token struct {
	item.Item

	Email string `json:"email"`
}

// MarshalEditor writes a buffer of html to edit a Token within the CMS
// and implements editor.Editable
func (t *Token) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(t,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Token field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Email", t, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Token editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Token"] = func() interface{} { return new(Token) }
}

// String defines how a Token is printed. Update it using more descriptive
// fields from the Token struct type
func (t *Token) String() string {
	return fmt.Sprintf("Token: %s", t.UUID)
}

func (t *Token) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}
func (t *Token) Hide(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (t *Token) BeforeAPICreate(res http.ResponseWriter, req *http.Request) error {
	// create an access configuration including the duration after which the
	// token will expire, the ResponseWriter to write the token to, and which
	// of the req.Header or req.Cookie{}
	cfg := &access.Config{
		ExpireAfter:    time.Hour * 24 * 7,
		ResponseWriter: res,
		TokenStore:     req.Header,
	}

	user, pass, ok := req.BasicAuth()
	if !ok {
		res.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("Basic authentication missing or invalid")
	}
	fmt.Printf("Username: %s\nPassword: %s\n", user, pass)

	// Grant access to the user based on the request
	grant, err := access.Login(user, pass, cfg)
	if err != nil {
		res.Header().Del("authorization")
		res.WriteHeader(http.StatusBadRequest)
		return err
	}

	fmt.Printf(
		"The access token for user (%s) is: %s\n",
		grant.Key, grant.Token,
	)

	return fmt.Errorf("User %s has been authenticated", t.Email)
}
