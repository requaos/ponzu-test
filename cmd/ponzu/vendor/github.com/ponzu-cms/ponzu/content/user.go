package content

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bosssauce/reference"
	"github.com/requaos/access"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type User struct {
	item.Item

	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Email             string   `json:"email"`
	Password          string   `json:"password"`
	PositionTitle     string   `json:"position_title"`
	AccountStatus     string   `json:"account_status"`
	PurchasedProducts []string `json:"purchased_products"`
}

// MarshalEditor writes a buffer of html to edit a User within the CMS
// and implements editor.Editable
func (u *User) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(u,
		// Take note that the first argument to these Input-like functions
		// is the string version of each User field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("FirstName", u, map[string]string{
				"label":       "FirstName",
				"type":        "text",
				"placeholder": "Enter the FirstName here",
			}),
		},
		editor.Field{
			View: editor.Input("LastName", u, map[string]string{
				"label":       "LastName",
				"type":        "text",
				"placeholder": "Enter the LastName here",
			}),
		},
		editor.Field{
			View: editor.Input("Email", u, map[string]string{
				"label":       "Email",
				"type":        "text",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.Input("Password", u, map[string]string{
				"label":       "Password",
				"type":        "password",
				"placeholder": "Enter the Email here",
			}),
		},
		editor.Field{
			View: editor.Input("PositionTitle", u, map[string]string{
				"label":       "PositionTitle",
				"type":        "text",
				"placeholder": "Enter the PositionTitle here",
			}),
		},
		editor.Field{
			View: editor.Input("AccountStatus", u, map[string]string{
				"label":       "AccountStatus",
				"type":        "text",
				"placeholder": "Enter the AccountStatus here",
			}),
		},
		editor.Field{
			View: reference.SelectRepeater("PurchasedProducts", u, map[string]string{
				"label": "PurchasedProducts",
			},
				"Product",
				`Product: {{ .name }}`,
			),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render User editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["User"] = func() interface{} { return new(User) }
}

// String defines how a User is printed. Update it using more descriptive
// fields from the User struct type
func (u *User) String() string {
	return fmt.Sprintf("%s, %s", u.LastName, u.FirstName)
}

// Create 'Approve' and 'Reject' buttons
func (u *User) Approve(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (u *User) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (u *User) Update(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (u *User) Hide(res http.ResponseWriter, req *http.Request) error {
	if !access.IsOwner(req, req.Header, u.Email) {
		return item.ErrAllowHiddenItem
	}

	return nil
}

func (u *User) BeforeAPICreate(res http.ResponseWriter, req *http.Request) error {
	// Check if email address is already in use based on the request
	err := access.Check(u.Email)
	if err == nil {
		if err = access.Pending(u.Email); err != nil {
			return fmt.Errorf("error adding user to pending bucket: %v", err)
		}
	}

	return err
}

func (u *User) BeforeAPIUpdate(res http.ResponseWriter, req *http.Request) error {
	if !access.IsOwner(req, req.Header, u.Email) {
		return fmt.Errorf(
			"grant provided is not owner of user, from %s",
			req.RemoteAddr,
		)
	}

	// request contains proper, valid token
	return nil
}

func (u *User) AfterApprove(res http.ResponseWriter, req *http.Request) error {
	// create an access configuration including the duration after which the
	// token will expire, the ResponseWriter to write the token to, and which
	// of the req.Header or req.Cookie{}
	cfg := &access.Config{
		ExpireAfter:    time.Hour * 24 * 7,
		ResponseWriter: res,
		TokenStore:     req.Header,
	}

	// Grant access to the user based on the request
	grant, err := access.Grant(u.Email, u.Password, cfg)
	if err != nil {
		res.Header().Del("authorization")
		res.WriteHeader(http.StatusBadRequest)
		return err
	}

	fmt.Printf(
		"The access token for user (%s) is: %s\n",
		grant.Key, grant.Token,
	)
	// Since this request is going to be made by an admin,
	// the admin doesn't need the user's access token.
	// We are just lodging the 'grant' for when the user logs in
	res.Header().Del("authorization")

	return nil
}

func (u *User) AfterReject(res http.ResponseWriter, req *http.Request) error {
	return access.ClearPending(u.Email)
}

func (u *User) AfterAdminDelete(res http.ResponseWriter, req *http.Request) error {
	return access.ClearPending(u.Email)
}

func (u *User) AfterDelete(res http.ResponseWriter, req *http.Request) error {
	return access.ClearPending(u.Email)
}
