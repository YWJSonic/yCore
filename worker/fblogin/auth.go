package fblogin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

func New() {
	// 設定 Facebook OAuth2 認證相關參數
	conf := &oauth2.Config{
		ClientID:     "1886847135011158",
		ClientSecret: "ea30df510efeba6b47d8b08536823527",
		RedirectURL:  "http://localhost/callback",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}

	// 建立 Gin 引擎
	r := gin.Default()
	r.LoadHTMLGlob("*.html")

	// 設定路由
	r.GET("/", Home)

	r.GET("/auth", func(c *gin.Context) {
		// 重定向到 Facebook 登入頁面
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusTemporaryRedirect, url)
	})
	r.GET("/callback", func(c *gin.Context) {
		// 取得授權代碼
		code := c.Query("code")

		// 使用授權代碼來取得存取權杖
		token, err := conf.Exchange(c.Request.Context(), code)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 使用存取權杖來存取 Facebook API，例如取得用戶的基本資訊
		client := conf.Client(c.Request.Context(), token)
		resp, err := client.Get("https://graph.facebook.com/v16.0/me?fields=id,name,email")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()

		// 解析 API 響應
		var user struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Email string `json:"email"`
		}
		err = json.NewDecoder(resp.Body).Decode(&user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 在這裡使用用戶的資訊來進行登入或註冊

		// 回應用戶
		c.String(http.StatusOK, fmt.Sprintf("Welcome, %s!", user.Name))
	})
	r.POST("/logincallback", func(c *gin.Context) {
		// 解析前端完成登入後得回傳資料
		var authRes struct {
			AuthResponse struct {
				AccessToken                 string `json:"accessToken"`
				UserID                      string `json:"userID"`
				ExpiresIn                   int    `json:"expiresIn"`
				SignedRequest               string `json:"signedRequest"`
				GraphDomain                 string `json:"graphDomain"`
				Data_access_expiration_time int    `json:"data_access_expiration_time"`
			} `json:"authResponse"`
			Status string `json:"status"`
		}
		err := json.NewDecoder(c.Request.Body).Decode(&authRes)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// 在這裡使用用戶的資訊來進行登入或註冊

		// 回應用戶
		c.String(http.StatusOK, fmt.Sprintf("Welcome, %s!", authRes.AuthResponse.UserID))
	})

	// 啟動應用程式
	r.Run(":80")
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
