package device

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"strings"
)

func Register(r *gin.RouterGroup) {
	r.GET("/device/:slug", Detail)
}

func Detail(c *gin.Context) {
	slug := c.Param("slug")
	doc, err := app.Scrape(slug)
	if err != nil {
		app.ErrorResponse(c, err)
		return
	}

	device := &Device{}
	device.Name = doc.Find(".review-header h1.specs-phone-name-title").Text()
	device.Specs = make(map[string]Specs)

	doc.Find("#specs-list table").Each(func(i int, group *goquery.Selection) {
		groupName := strings.ToLower(group.Find("th").Text())
		rows := group.Find("tr")
		specs := make(Specs, rows.Length())
		rows.Each(func(i int, row *goquery.Selection) {
			name := row.Find("td.ttl a").Text()
			if name == "" {
				name = strings.Title(groupName)
			}
			specs[i] = &Spec{
				Name:  name,
				Value: row.Find("td.nfo").Text(),
			}
		})
		device.Specs[groupName] = specs
	})

	app.SuccessResponse(c, device)
}
