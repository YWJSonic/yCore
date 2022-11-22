package restful

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinRouter(router *RestfulDriver) {
	routerPath := "/api/DataService"

	router.Handle(http.MethodPost, routerPath+"/Token", ginToken)
	router.Handle(http.MethodGet, routerPath+"/MatchById", ginMatchById)
	router.Handle(http.MethodGet, routerPath+"/CompetitionById", ginCompetitionById)
	router.Handle(http.MethodGet, routerPath+"/Region", ginRegion)
}

func ginToken(ctx *gin.Context) {
	// req := adaptor.TokenRequestDto{}

	// if res, err := adaptor.Token(req); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// } else {
	// 	ctx.Data(http.StatusOK, "application/json; charset=utf-8", res)
	// }
}

func ginMatchById(ctx *gin.Context) {
	// req := adaptor.MatchByIdRequestDto{}

	// token, ok := ctx.GetQuery("token")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query token not find"))
	// 	return
	// }

	// matchId, ok := ctx.GetQuery("MatchId")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query MatchId not find"))
	// 	return
	// }

	// includeMatchStats, ok := ctx.GetQuery("IncludeMatchStats")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query IncludeMatchStats not find"))
	// 	return
	// }

	// req.Token = token
	// req.MatchId = matchId
	// req.IncludeMatchStats = includeMatchStats

	// if res, err := adaptor.MatchById(req); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// } else {
	// 	ctx.Data(http.StatusOK, "application/json; charset=utf-8", res)
	// }
}

func ginCompetitionById(ctx *gin.Context) {
	// req := adaptor.CompetitionByIdRequestDto{}

	// token, ok := ctx.GetQuery("token")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query token not find"))
	// 	return
	// }

	// competitionId, ok := ctx.GetQuery("CompetitionId")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query CompetitionId not find"))
	// 	return
	// }

	// req.Token = token
	// req.CompetitionId = competitionId

	// if res, err := adaptor.CompetitionById(req); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// } else {
	// 	ctx.Data(http.StatusOK, "application/json; charset=utf-8", res)
	// }
}

func ginRegion(ctx *gin.Context) {
	// req := adaptor.RegionRequestDto{}

	// token, ok := ctx.GetQuery("token")
	// if !ok {
	// 	ctx.JSON(http.StatusInternalServerError, errors.New("Query token not find"))
	// 	return
	// }

	// req.Token = token

	// if res, err := adaptor.Region(req); err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, err)
	// } else {
	// 	ctx.Data(http.StatusOK, "application/json; charset=utf-8", res)
	// }
}
