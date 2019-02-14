package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func InitializeRoute() {
	r := gin.Default()

	r.Run(":" + os.Getenv("PORT"))
}