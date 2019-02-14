package brand

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"github.com/mlalitthapa/phone-scrapper/utils"
)

func Register(r *gin.RouterGroup) {
	//app.DB.AutoMigrate(&Brand{})

	r.GET("/brand", GetBrands)
}

func GetBrands(c *gin.Context) {
	doc, err := app.Scrape(utils.BrandUrl)
	if err != nil {
		app.ErrorResponse(c, err)
	}

	var brands []*Brand

	doc.Find("div.st-text table tbody a").Each(func(i int, link *goquery.Selection) {
		brand := &Brand{
			Name:    link.Clone().Children().Remove().End().Text(),
			Devices: link.Find("span").Text(),
		}

		brandLink, exists := link.Attr("href")
		if exists {
			brand.Slug = brandLink
		}

		brands = append(brands, brand)
	})

	app.SuccessResponse(c, brands)
}
