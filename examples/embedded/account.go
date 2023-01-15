package embedded

import "github.com/gzuidhof/tygo/examples/abstract"

type (
	omit *struct{}

	User struct {
		abstract.StructBar
		Email string
	}

	Account struct {
		Name    string
		Balance int64 `json:"balance"`
	}

	AccountResponse struct {
		*Account
		*User
		Balance omit `json:"balance,omitempty"`
	}
)
