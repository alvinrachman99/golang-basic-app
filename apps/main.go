package main

import (
	"golang-basic/apps/config"
	"golang-basic/apps/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.ConnectDb()
	if err != nil {
		panic(err)
	}

	router := gin.New()

	router.Use(gin.Logger())

	authController := controller.AuthController{}

	router.GET("/ping", Ping)
	router.POST("/v1/auth/register", authController.Register)

	router.Run(":4444")
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}
