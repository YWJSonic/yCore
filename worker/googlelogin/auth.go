package googlelogin

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"ycore/driver/encryption/jwt"
	"ycore/module/myrestful"

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
	router.LoadHTMLGlob("*.html")

	router.Handle(http.MethodGet, "/", Home)
	router.Handle(http.MethodGet, "/ouath/google/url", GoogleAccsess)
	router.Handle(http.MethodGet, "/ouath/google/login", GoogleRedirectLogin)
	router.Handle(http.MethodPost, "/ouath/google/login", GoogleLogin)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GoogleAccsess(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, oauthURL())
}

func oauthURL() string {
	u := "%s?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"
	scopes := []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}

	return fmt.Sprintf(u, config.Auth_uri, config.Client_id, strings.Join(scopes, "+"), config.Redirect_uris[0])
}

func GoogleRedirectLogin(c *gin.Context) {
	fmt.Println(c.Request.URL)
	querys, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	c.JSON(http.StatusOK, querys)

}

func GoogleLogin(c *gin.Context) {
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

	jwtClimDto := &JwtClimDto{}
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

var mockHtml = `<html>
<body>
<script src="https://accounts.google.com/gsi/client" async defer></script>
<div id="g_id_onload"
data-client_id="14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com"
data-context="signin"
data-ux_mode="popup"
data-login_uri="http://localhost/ouath/google/login"
data-auto_prompt="false">
</div>
<div class="g_id_signin"
data-type="standard"
data-size="medium"
data-theme="filled_blue"
data-text="signin_with"
data-shape="pill"
data-callback="onSinbgIn" 
data-logo_alignment="left">
</div>
<script>
function onSinbgIn(googleUser){
  var profile = googleUser.getBasicProfile();
  print(googleUser.getBasicProfile().getEmail())
}
</script>
</body>
</html>`
