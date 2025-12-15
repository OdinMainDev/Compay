package efwjwt

type JwtClaimsData map[string]string

const (
	EncryptedKeySuffix = "_efw"
	GinJwtClaimsData   = "_efw/jwt/claims_data"
)
