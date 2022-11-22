package myrestful

import (
	"ycore/driver/connect/restful"
)

func New() *restful.RestfulDriver {
	router := restful.New()
	return router
}

// func respontHtml(c *gin.Context, name, payload string) {
// 	c.HTML(http.StatusOK, name, payload)
// }

// func respontJson(c *gin.Context, name, payload string) {
// 	c.JSON(http.StatusOK, payload)
// }
