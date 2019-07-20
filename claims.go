package jwt

// StandardClaims https://tools.ietf.org/html/rfc7519#section-4.1
type StandardClaims struct {
	// Audience claim identifies the recipients that the JWT is intended for.
	Audience []string `json:"aud,omitempty"`

	// ExpiresAt claim identifies the expiration time on or after which the JWT MUST NOT be accepted for processing.
	ExpiresAt int64 `json:"exp,omitempty"`

	// ID claim provides a unique identifier for the JWT.
	ID string `json:"jti,omitempty"`

	// IssuedAt claim identifies the time at which the JWT was issued.
	// This claim can be used to determine the age of the JWT.
	IssuedAt int64 `json:"iat,omitempty"`

	// Issuer claim identifies the principal that issued the JWT.
	Issuer string `json:"iss,omitempty"`

	// NotBefore claim identifies the time before which the JWT MUST NOT be accepted for processing.
	NotBefore int64 `json:"nbf,omitempty"`

	// Subject claim identifies the principal that is the subject of the JWT.
	Subject string `json:"sub,omitempty"`
}