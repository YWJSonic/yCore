package jwt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func Decode(tokenString string, signatures map[string]string, jwtDto jwt.Claims) error {
	// sample token string taken from the New example

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.ParseWithClaims(tokenString, jwtDto, func(token *jwt.Token) (interface{}, error) {
		alg := token.Method.Alg()
		kid, ok := token.Header["kid"]
		if !ok {
			return nil, errors.New("")
		}
		signature, ok := signatures[kid.(string)]
		if !ok {
			return nil, errors.New("")
		}

		if isNone(alg) {
			return jwt.UnsafeAllowNoneSignatureType, nil
		} else if isEs(alg) {
			return jwt.ParseECPublicKeyFromPEM([]byte(signature))
		} else if isRs(alg) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(signature))
		} else if isEd(alg) {
			return jwt.ParseEdPublicKeyFromPEM([]byte(signature))
		}
		return nil, errors.New("error")
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalid JWT")
	}

	return nil
}
func isNone(alg string) bool {
	return alg == "none" || alg == ""
}
func isEs(alg string) bool {
	return strings.HasPrefix(alg, "ES")
}

func isRs(alg string) bool {
	return strings.HasPrefix(alg, "RS") || strings.HasPrefix(alg, "PS")
}

func isEd(alg string) bool {
	return alg == "EdDSA"
}
