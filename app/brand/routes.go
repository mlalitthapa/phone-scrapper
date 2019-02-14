package brand

import (
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"net/http"
)

func Register(r *gin.RouterGroup) {
	app.DB.AutoMigrate(&Brand{})

	r.GET("/brand", GetBrands)
}

func GetBrands(c *gin.Context) {
	c.JSON(http.StatusOK, "Brands")
}
