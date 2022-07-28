package myrestful

import (
	"net/http"
	"ycore/driver/connect/restful"

	"github.com/gin-gonic/gin"
)

func New() {

	router := restful.New()
	Router(router)
	router.Run(":80")
}

func Router(router *restful.RestfulDriver) {
	// router.Handle("/SignUp/", SingUp)
}

func respontHtml(c *gin.Context, name, payload string) {
	c.HTML(http.StatusOK, name, payload)
}

func respontJson(c *gin.Context, name, payload string) {
	c.JSON(http.StatusOK, payload)
}
