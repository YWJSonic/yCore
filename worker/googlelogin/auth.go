package googlelogin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/YWJSonic/ycore/driver/encryption/jwt"
	"github.com/YWJSonic/ycore/module/myrestful"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/sheets/v4"
)

var config *gConfig = &gConfig{
	Web: Web{
		Client_id:                   "14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com",
		Project_id:                  "online-mrcf",
		Auth_uri:                    "https://accounts.google.com/o/oauth2/auth",
		Token_uri:                   "https://oauth2.googleapis.com/token",
		Auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
		Redirect_uris:               []string{"http://localhost/ouath/google/login"},
		Javascript_origins:          []string{"https://localhost:12345"},
	},
}

func New() {

	// if err := yamlloader.LoadYaml(configPath, config); err != nil {
	// 	return
	// }

	ctx := context.Background()
	client, err := sheets.NewService(ctx)
	fmt.Println(client, err)

	// Restful
	restfulRouter := myrestful.GinNew()
	Router(restfulRouter)
	restfulRouter.Run(":80")
}

func Router(router *myrestful.RestfulDriver) {
	// router.LoadHTMLFiles("*.html")
	router.LoadHTMLGlob("./worker/googlelogin/*.html")

	router.Handle(http.MethodGet, "/", Home)

	// 以不對外提供密鑰方式登入, 第一次取得權杖(code),使用權帳存取資料
	router.Handle(http.MethodGet, "/login", HttpOptionLogin)                       // 跳轉登入
	router.Handle(http.MethodGet, "/ouath/google/login", GoogleLoginRedirectGET)   // 前端 Google 工具登入回傳
	router.Handle(http.MethodPost, "/ouath/google/login", GoogleLoginRedirectPOST) // Server 跳轉登入後返回資料
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func oauthURL() string {
	u := "%s?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	redirectUr := url.QueryEscape(config.Redirect_uris[0])
	return fmt.Sprintf(u, config.Auth_uri, config.Client_id, strings.Join(scopes, "+"), redirectUr)
}

func HttpOptionLogin(c *gin.Context) {

	//     https://accounts.google.com/o/oauth2/auth?					client_id=14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com&redirect_uri=http:%2F%2Flocalhost%2Fouath%2Fgoogle%2Flogin&response_type=code&scope=https://www.googleapis.com/auth/userinfo.profile+https://www.googleapis.com/auth/userinfo.email
	// url := `https://accounts.google.com/o/oauth2/auth?client_id=14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com&redirect_uri=http%3A%2F%2Flocalhost%2Fouath%2Fgoogle%2Flogin&scope=openid%20email%20profile&response_type=code&state=abcdef1234567890`
	c.Redirect(http.StatusFound, oauthURL())
}

var (
	clientSecret     = "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" // 替換為你的 Google 客戶端密鑰
	tokenEndpoint    = "https://oauth2.googleapis.com/token"
	userInfoEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"
)

func GoogleLoginRedirectGET(c *gin.Context) {
	code := c.Query("code")

	// 請求訪問令牌
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", config.Client_id)
	data.Set("client_secret", clientSecret)
	data.Set("redirect_uri", config.Redirect_uris[0])
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", tokenEndpoint, bytes.NewBufferString(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "創建請求時出錯"})
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "發送請求時出錯"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "讀取回應時出錯"})
		return
	}

	var tokenResponse map[string]interface{}
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析訪問令牌回應時出錯"})
		return
	}

	accessToken, ok := tokenResponse["access_token"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "無法獲取訪問令牌"})
		return
	}

	// 使用訪問令牌請求用戶資訊
	userInfoResp, err := http.Get(fmt.Sprintf("%s?access_token=%s", userInfoEndpoint, accessToken))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "請求用戶資訊時出錯"})
		return
	}
	defer userInfoResp.Body.Close()

	userInfoBody, err := io.ReadAll(userInfoResp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "讀取用戶資訊回應時出錯"})
		return
	}

	var userInfo UserInfoDto
	if err := json.Unmarshal(userInfoBody, &userInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "解析用戶資訊回應時出錯"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userInfo": userInfo})
}

func GoogleLoginRedirectPOST(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}
	values, err := url.ParseQuery(string(b))
	if err != nil {
		c.HTML(http.StatusOK, "login", err)
		return
	}

	quarys, ok := values["credential"]
	if !ok {
		return
	}

	certMap, err := GetCerts()
	if err != nil {
		return
	}

	jwtClimDto := &UserInfoDto{}
	if err = jwt.Decode(quarys[0], certMap, jwtClimDto); err != nil {
		return
	}

	c.JSON(http.StatusOK, jwtClimDto)
}

func GetCerts() (map[string]string, error) {

	req, err := http.NewRequest(http.MethodGet, config.Auth_provider_x509_cert_url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	certs := map[string]string{}
	err = json.Unmarshal(body, &certs)
	if err != nil {
		log.Fatal(err)
	}
	return certs, nil
}
