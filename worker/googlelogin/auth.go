package googlelogin

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"ycore/module/myrestful"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/sheets/v4"
)

var config *gConfig

func New() {

	// if err := yamlloader.LoadYaml(configPath, config); err != nil {
	// 	return
	// }

	config = &gConfig{
		Web: Web{
			Client_id:                   "14099599407-n78q9cn8eslht1ksculubui6oujn4mav.apps.googleusercontent.com",
			Project_id:                  "online-mrcf",
			Auth_uri:                    "https://accounts.google.com/o/oauth2/auth",
			Token_uri:                   "https://oauth2.googleapis.com/token",
			Auth_provider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
			Redirect_uris:               []string{"https://localhost"},
			Javascript_origins:          []string{"https://localhost:12345"},
		},
	}

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
	router.Handle(http.MethodGet, "/ouath/google/login", GoogleLogin)
	router.Handle(http.MethodPost, "/ouath/google/login", GoogleLogin)
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GoogleAccsess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"url": oauthURL(),
	})
}

func oauthURL() string {
	u := "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=code&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(u, config.Client_id, "https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email", "https://accounts.google.com/o/oauth2/v2/auth")
}
func GoogleLogin(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))
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
