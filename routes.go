package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func InitializeRoute() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "HELLO WORLD")
	})

	r.Run(":" + os.Getenv("PORT"))
}