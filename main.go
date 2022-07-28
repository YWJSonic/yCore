package main

import (
	"flag"
	"net/http"
	"ycore/config"
	"ycore/driver/connect/restful"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	config.Init(*configPath)
	router := restful.New()
	router.Handle(http.MethodGet, "/", oauth)
	router.Run(":8088")
}

func oauth(c *gin.Context) {

	respontJson(c, CreateGoogleOAuthURL())
}

func respontJson(c *gin.Context, payload string) {
	c.JSON(http.StatusOK, payload)
}

func CreateGoogleOAuthURL() string {
	// 使用 lib 產生一個特定 config instance
	config := &oauth2.Config{
		//憑證的 client_id
		ClientID: "14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com",
		//憑證的 client_secret
		ClientSecret: "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",
		//當 Google auth server 驗證過後，接收從 Google auth server 傳來的資訊
		RedirectURL: "http://localhost:8080/google-login-callback",
		//告知 Google auth server 授權範圍，在這邊是取得用戶基本資訊和Email，Scopes 為 Google 提供
		Scopes: []string{
			"openid",
		},
		//指的是 Google auth server 的 endpoint，用 lib 預設值即可
		Endpoint: google.Endpoint,
	}

	//產生出 config instance 後，就可以使用 func AuthCodeURL 建立請求網址
	return config.AuthCodeURL("state")
}
