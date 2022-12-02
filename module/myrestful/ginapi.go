package myrestful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinRouter(router *RestfulDriver) {
	routerPath := "/api/DataService"

	router.Handle(http.MethodPost, routerPath+"/POST_Request", ginrequest)
	router.Handle(http.MethodGet, routerPath+"/GET_Request", ginrequest)
}

func ginrequest(ctx *gin.Context) {

	var err error
	var res []byte
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.Data(http.StatusOK, "application/json; charset=utf-8", res)
	}
}
