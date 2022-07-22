package main

import (
	"net/http"
	"strconv"
	"ycore/driver/connect/restful"

	"github.com/gin-gonic/gin"
)

var balance = 1000

func main() {

	router := restful.New()
	router.Handle(http.MethodGet, "/balance/", getBalance)

	_ = router.Run(":9123")
}

func getBalance(context *gin.Context) {
	var msg = "您的帳戶內有:" + strconv.Itoa(balance) + "元"
	context.JSON(http.StatusOK, gin.H{
		"amount":  balance,
		"status":  "ok",
		"message": msg,
	})
}
