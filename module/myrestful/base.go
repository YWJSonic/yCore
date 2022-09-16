package myrestful

import (
	"net/http"
	"ycore/driver/connect/restful"

	"github.com/gin-gonic/gin"
)

func New() *restful.RestfulDriver {
	router := restful.New()
	return router
}

func respontHtml(c *gin.Context, name, payload string) {
	c.HTML(http.StatusOK, name, payload)
}

func respontJson(c *gin.Context, name, payload string) {
	c.JSON(http.StatusOK, payload)
}
