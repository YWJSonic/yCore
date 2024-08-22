package googlelogin

import "github.com/golang-jwt/jwt/v5"

type JwtClimDto struct {
	Email          string `json:"email"`          // "sony79410@gmail.com"
	Email_verified bool   `json:"email_verified"` // true
	Azp            string `json:"azp"`            // "14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com"
	Name           string `json:"name"`           // "楊文哲（Sonic）"
	Picture        string `json:"picture"`        // "https://lh3.googleusercontent.com/a/AGNmyxaln4yq6fOwgJblycxQ5wx9RyJ7DmipLKXkkbwN=s96-c"
	Given_name     string `json:"given_name"`     // "文哲"
	Family_name    string `json:"family_name"`    // "楊"

	jwt.RegisteredClaims
}
