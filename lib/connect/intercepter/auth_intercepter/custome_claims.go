package auth_intercepter

import "context"

// auth0 validator.CustomClaimsの実装
// https://pkg.go.dev/github.com/webmsi/go-jwt-middleware/validator#ValidatedClaims

// CustomClaims contains custom data we want from the token.
type CustomClaims struct {
	Permissions []string `json:"permissions"`
}

// Validate does nothing for this example, but we need
// it to satisfy validator.CustomClaims interface.
func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}
