package creds

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

// JwtClaims are custom claims for emrys JWTs
type JwtClaims struct {
	Scope []string `json:"scope"`
	jwt.StandardClaims
}

// Valid determines if JwtClaim j is valid
func (j *JwtClaims) Valid() error {
	vErr := new(jwt.ValidationError)

	for _, s := range j.Scope {
		switch s {
		case "user", "miner":
		default:
			vErr.Inner = fmt.Errorf("token scope contains something other than user or miner")
			vErr.Errors |= jwt.ValidationErrorClaimsInvalid // Generic claims validation error
			return vErr
		}
	}
	return j.StandardClaims.Valid()
}
