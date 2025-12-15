package efwjwt

import "fmt"

var (
	ErrorJwtNilAuthorizationHeader  = fmt.Errorf("%s", "[EFW: JWT] NilAuthorizationHeader")
	ErrorJwtInvalidToken            = fmt.Errorf("%s", "[EFW: JWT] InvalidToken")
	ErrorJwtUnexpectedSigningMethod = fmt.Errorf("%s", "[EFW: JWT] UnexpectedSigningMethod")
	ErrorJwtUnexpectedParsingError  = func(err error) error {
		return fmt.Errorf("%s: %s", "[EFW: JWT] UnexpectedParsingError", err.Error())
	}
	ErrorJwtUnexpectedClaimsType = fmt.Errorf("%s", "[EFW: JWT] UnexpectedClaimsType")
	ErrorJwtInvalidJwtClaims     = fmt.Errorf("%s", "[EFW: JWT] InvalidJwtClaims")
)
