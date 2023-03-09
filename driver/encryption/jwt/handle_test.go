package jwt

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

func TestDecode(t *testing.T) {
	// sample token string taken from the New example
	tokenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImQyNWY4ZGJjZjk3ZGM3ZWM0MDFmMDE3MWZiNmU2YmRhOWVkOWU3OTIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE2Nzc2ODc2MDQsImF1ZCI6IjE0MDk5NTk5NDA3LW43OHE5Y244ZXNsaHQxa3NjdWx1YnVpNm91am40bWF2LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTA1MTMzNzAyODQxMjI3NjA0NDYxIiwiZW1haWwiOiJzb255Nzk0MTBAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF6cCI6IjE0MDk5NTk5NDA3LW43OHE5Y244ZXNsaHQxa3NjdWx1YnVpNm91am40bWF2LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwibmFtZSI6IualiuaWh-WTsu-8iFNvbmlj77yJIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FHTm15eGFsbjR5cTZmT3dnSmJseWN4UTV3eDlSeUo3RG1pcExLWGtrYndOPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IuaWh-WTsiIsImZhbWlseV9uYW1lIjoi5qWKIiwiaWF0IjoxNjc3Njg3OTA0LCJleHAiOjE2Nzc2OTE1MDQsImp0aSI6IjhmMzI2OWNlNTA3MDhkZWZmYWVhMzYwMmI4NGRhZWNkYTJhYTAyODkifQ.qZLt3UKaKrd9ZWaadt-5BXb3KuO5xqBjWkGbeJRrll7LIJZqfsXjAwngW_pxWpydVc0LAKwDqhCrdx1Dw28JFh-UDOKzU3cSAvGpPKdZCIAEqqZOFZ0qjME9dgXMRQ57EUVl2vIzpOPiKUmebxo8goR24c2RZ8E9_7U6jbp2DO4oRCPs0-_YhosmFnbmYTaL2jePPHiLAA78f0ISIXKNlkw3hKBiUYajZdZoTeq0URWLGq7ui8ilrjQknVoQUNo8XfTn6kATTwyB5GJz2OnlFNkiIzqmlpfjokeblo3YmQkm1kdwbwKUGDbgB1dxk071BXjkYbcLAKtuXDLmeYSs7Q"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return token.Raw, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}

func TestXxx(t *testing.T) {

	// 假设以下是从请求中获取到的 JWT
	tokenString := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImQyNWY4ZGJjZjk3ZGM3ZWM0MDFmMDE3MWZiNmU2YmRhOWVkOWU3OTIiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJuYmYiOjE2Nzc2ODc2MDQsImF1ZCI6IjE0MDk5NTk5NDA3LW43OHE5Y244ZXNsaHQxa3NjdWx1YnVpNm91am40bWF2LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTA1MTMzNzAyODQxMjI3NjA0NDYxIiwiZW1haWwiOiJzb255Nzk0MTBAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImF6cCI6IjE0MDk5NTk5NDA3LW43OHE5Y244ZXNsaHQxa3NjdWx1YnVpNm91am40bWF2LmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwibmFtZSI6IualiuaWh-WTsu-8iFNvbmlj77yJIiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FHTm15eGFsbjR5cTZmT3dnSmJseWN4UTV3eDlSeUo3RG1pcExLWGtrYndOPXM5Ni1jIiwiZ2l2ZW5fbmFtZSI6IuaWh-WTsiIsImZhbWlseV9uYW1lIjoi5qWKIiwiaWF0IjoxNjc3Njg3OTA0LCJleHAiOjE2Nzc2OTE1MDQsImp0aSI6IjhmMzI2OWNlNTA3MDhkZWZmYWVhMzYwMmI4NGRhZWNkYTJhYTAyODkifQ.qZLt3UKaKrd9ZWaadt-5BXb3KuO5xqBjWkGbeJRrll7LIJZqfsXjAwngW_pxWpydVc0LAKwDqhCrdx1Dw28JFh-UDOKzU3cSAvGpPKdZCIAEqqZOFZ0qjME9dgXMRQ57EUVl2vIzpOPiKUmebxo8goR24c2RZ8E9_7U6jbp2DO4oRCPs0-_YhosmFnbmYTaL2jePPHiLAA78f0ISIXKNlkw3hKBiUYajZdZoTeq0URWLGq7ui8ilrjQknVoQUNo8XfTn6kATTwyB5GJz2OnlFNkiIzqmlpfjokeblo3YmQkm1kdwbwKUGDbgB1dxk071BXjkYbcLAKtuXDLmeYSs7Q"

	token, err := jwt.ParseWithClaims(tokenString, nil, func(token *jwt.Token) (interface{}, error) {
		alg := token.Method.Alg()
		if isRs(alg) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(`-----BEGIN CERTIFICATE-----MIIDJzCCAg+gAwIBAgIJAIF1A77L4wb2MA0GCSqGSIb3DQEBBQUAMDYxNDAyBgNVBAMMK2ZlZGVyYXRlZC1zaWdub24uc3lzdGVtLmdzZXJ2aWNlYWNjb3VudC5jb20wHhcNMjMwMjE5MTUyMjM1WhcNMjMwMzA4MDMzNzM1WjA2MTQwMgYDVQQDDCtmZWRlcmF0ZWQtc2lnbm9uLnN5c3RlbS5nc2VydmljZWFjY291bnQuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxQJXLHo7DQfiKztmmDFheKx2JqZZsEoJh/3P+CX4X1W//wpEDbYFVKTwgBboKzjCbdj0N4m4Sg2k8a980jorfSXgdtrT9m5ug3MG2qy50l7/ofmiYys5EPbqEvcCC+1EyzxTMHUdP7Jlf/tPZvsqaTTogLwUIRmUeza4qCE80nZ+/1bUj8njLhZx0wM8H8Q6gSfSaJ/YOeg6L3bBViYWANyKS6hHf4RgVIkOlY+VXRdAGJjzCmespI1/gOmA9jRvrCBD4lWoUJXGM3+lX5qwCgcWCWQDuncBhISsOhGYZZcW3Ji1Af0JCnwdgwFsfaJyuub/blkcwY5ERXaPy+44mQIDAQABozgwNjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIHgDAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOCAQEAVKPm2PqrpDo0athNogM4Ef39qIsHmM2ZqzJaBfq4XoesX7YvoP2AaH7T1Zs+NezwIGCGxokUdvlgSMAuMsqdKne+bsItJgOh1YYZpK7BmrPyJpGxdZHr8Rmpxfk68J28UFd88Urz1GYLlG+PXptKqWVfBspWB1swGq6rVxB2zFHwl6lLoqnlDUH/DXOPUU7j7vWk8ZpIkkq449pZkERv6NtRElXHj3zlgCWejgcG5jzekocrhT6hc3mbttaHwNuy6t2YLqN5/rX583Lc+qxuyfjlFjU1eQ+Gqbt2nVJKJUr+uRxQ0uKaFc2O89XGpONPO88AfQuB7+YgaIMks2QVPg==-----END CERTIFICATE-----`))
		} else if isEs(alg) {
			return jwt.ParseRSAPublicKeyFromPEM([]byte("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"))
		}
		return nil, errors.New("error")
	})

	fmt.Println(err)
	// 验证 JWT 是否有效
	if !token.Valid {
		fmt.Println("Invalid JWT")
		return
	}

	// 获取 JWT 中的自定义字段
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("JWT is valid")
		fmt.Println("Username:", claims["username"])
		fmt.Println("Expires At:", time.Unix(int64(claims["exp"].(float64)), 0))
	}
}
