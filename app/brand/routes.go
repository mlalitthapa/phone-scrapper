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
		brandLink, _ := link.Attr("href")
		brand := &Brand{
			Name:    link.Clone().Children().Remove().End().Text(),
			Slug:    brandLink,
			Devices: link.Find("span").Text(),
		}
		brands = append(brands, brand)
	})

	app.SuccessResponse(c, brands)
}
