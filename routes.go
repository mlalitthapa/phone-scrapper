package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app/brand"
	"github.com/mlalitthapa/phone-scrapper/app/device"
	"os"
)

func InitializeRoute() {
	r := gin.Default()

	api := r.Group("/api")
	v1 := api.Group("/v1")

	brand.Register(v1)
	device.Register(v1)

	r.Run(":" + os.Getenv("PORT"))
}