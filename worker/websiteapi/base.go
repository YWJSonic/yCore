package websiteapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/YWJSonic/ycore/module/mydb"
	"github.com/YWJSonic/ycore/module/myrestful"

	"github.com/gin-gonic/gin"
)

var db *mydb.Manager
var apiMap = []string{}

func New() {
	// DB
	dbManager, err := mydb.NewArangoDB("http://34.81.111.226:8529", "", "", "WebData")
	if err != nil {
		return
	}
	db = dbManager

	// Restful

	restfulRouter := myrestful.GinNew()
	Router(restfulRouter)
	restfulRouter.Run(":8080")
}

func Router(router *myrestful.RestfulDriver) {
	apiMap = append(apiMap, "GET '/'")
	apiMap = append(apiMap, "GET '/types'")
	apiMap = append(apiMap, "GET '/types/:typeId'")
	apiMap = append(apiMap, "GET '/types/:typeId/:namelike'")

	router.Handle(http.MethodGet, "/", getAllApi)
	router.Handle(http.MethodGet, "/types", getTypes)
	router.Handle(http.MethodGet, "/types/:typeId", getHistoryTypeId)
	router.Handle(http.MethodGet, "/types/:typeId/:namelike", getHistoryNameLike)
}

func getAllApi(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, apiMap)
}

func getTypes(ctx *gin.Context) {
	types, err := dbGetTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	var context string
	for _, typeInfo := range types {
		context += fmt.Sprintf("%v %v\n", typeInfo.TypeId, typeInfo.TypeName)
	}

	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(context))
}

func getHistoryTypeId(ctx *gin.Context) {
	typesIdStr := ctx.Param("typeId")
	typesId, err := strconv.Atoi(typesIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	names, err := dbGetItems(ctx, typesId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	var context string
	for _, name := range names {
		context += fmt.Sprintf("%v\n", name)
	}

	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(context))
}

func getHistoryNameLike(ctx *gin.Context) {

	typesIdStr := ctx.Param("typeId")
	typesId, err := strconv.Atoi(typesIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	namelike := ctx.Param("namelike")

	names, err := dbGetNameLike(ctx, typesId, namelike)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	var context string
	for _, name := range names {
		context += fmt.Sprintf("%v\n", name)
	}

	ctx.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(context))
}
