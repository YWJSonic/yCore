package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(serverHash string, data map[string]interface{}, expTime time.Time) (string, error) {
	dataMap := jwt.MapClaims{
		// "jti": "",                     // token 編號
		// "iss": "",                     // token 發行者
		"exp": jwt.NewNumericDate(expTime), //  token 過期時間
		// "iat": time.Now().UnixMilli(), // token 發行時間
		// "nbf": time.Now().UnixMilli(), // token 可使啟用時間
	}

	for k, v := range data {
		dataMap[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataMap)
	tokenString, err := token.SignedString([]byte(serverHash))
	return tokenString, err
}

func ParseToken(serverHash string, tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(serverHash), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return claims, err
	}
}
