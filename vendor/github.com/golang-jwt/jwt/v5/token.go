package jwt

import (
	"encoding/base64"
	"encoding/json"
)

// Keyfunc will be used by the Parse methods as a callback function to supply
// the key for verification.  The function receives the parsed, but unverified
// Token.  This allows you to use properties in the Header of the token (such as
// `kid`) to identify which key to use.
type Keyfunc func(*Token) (interface{}, error)

// Token represents a JWT Token.  Different fields will be used depending on
// whether you're creating or parsing/verifying a token.
type Token struct {
	Raw       string                 // Raw contains the raw token.  Populated when you [Parse] a token
	Method    SigningMethod          // Method is the signing method used or to be used
	Header    map[string]interface{} // Header is the first segment of the token in decoded form
	Claims    Claims                 // Claims is the second segment of the token in decoded form
	Signature []byte                 // Signature is the third segment of the token in decoded form.  Populated when you Parse a token
	Valid     bool                   // Valid specifies if the token is valid.  Populated when you Parse/Verify a token
}

// New creates a new [Token] with the specified signing method and an empty map
// of claims. Additional options can be specified, but are currently unused.
func New(method SigningMethod, opts ...TokenOption) *Token {
	return NewWithClaims(method, MapClaims{}, opts...)
}

// NewWithClaims creates a new [Token] with the specified signing method and
// claims. Additional options can be specified, but are currently unused.
func NewWithClaims(method SigningMethod, claims Claims, opts ...TokenOption) *Token {
	return &Token{
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}

// SignedString creates and returns a complete, signed JWT. The token is signed
// using the SigningMethod specified in the token.
func (t *Token) SignedString(key interface{}) (string, error) {
	sstr, err := t.SigningString()
	if err != nil {
		return "", err
	}

	sig, err := t.Method.Sign(sstr, key)
	if err != nil {
		return "", err
	}

	return sstr + "." + t.EncodeSegment(sig), nil
}

// SigningString generates the signing string.  This is the most expensive part
// of the whole deal.  Unless you need this for something special, just go
// straight for the SignedString.
func (t *Token) SigningString() (string, error) {
	h, err := json.Marshal(t.Header)
	if err != nil {
		return "", err
	}

	c, err := json.Marshal(t.Claims)
	if err != nil {
		return "", err
	}

	return t.EncodeSegment(h) + "." + t.EncodeSegment(c), nil
}

// EncodeSegment encodes a JWT specific base64url encoding with padding
// stripped. In the future, this function might take into account a
// [TokenOption]. Therefore, this function exists as a method of [Token], rather
// than a global function.
func (*Token) EncodeSegment(seg []byte) string {
	return base64.RawURLEncoding.EncodeToString(seg)
}
