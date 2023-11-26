package main

import (
	"fmt"
	"net/http"
	in "test/internal/app"
	mw "test/internal/middleware"

	"github.com/gin-gonic/gin"
)

func getHandler() *gin.Engine {
	r := gin.Default()
	r.Use(mw.RoleCheck())
	r.GET("", dateHandler)
	return r
}

func main() {
	router := getHandler()
	router.Run("localhost:8080")
}

func dateHandler(ctx *gin.Context) {
	date := in.TimeCounter()
	ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("Days left: %d", date))
}

