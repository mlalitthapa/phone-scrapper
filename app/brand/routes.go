package brand

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"github.com/mlalitthapa/phone-scrapper/utils"
	"sort"
	"strconv"
)

func Register(r *gin.RouterGroup) {
	//app.DB.AutoMigrate(&Brand{})

	r.GET("/brand", GetBrands)
	r.GET("/brand/:slug", GetBrandDevices)
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

func GetBrandDevices(c *gin.Context) {
	slug := c.Param("slug")

	doc, err := app.Scrape(slug)
	if err != nil {
		app.ErrorResponse(c, err)
		return
	}

	var devices []*device
	var pages app.Pages

	doc.Find("#review-body ul li a").Each(func(i int, d *goquery.Selection) {
		name := d.Find("strong span").Text()
		slug, _ := d.Attr("href")
		image := d.Find("img")
		imageLink, _ := image.Attr("src")
		imageAlt, _ := image.Attr("title")
		devices = append(devices, &device{
			Name: name,
			Slug: slug,
			Image: deviceImage{
				Src: imageLink,
				Alt: imageAlt,
			},
		})
	})

	navPages := doc.Find(".review-nav .nav-pages")

	navPages.Find("a").Each(func(i int, p *goquery.Selection) {
		link, _ := p.Attr("href")
		page, err := strconv.ParseInt(p.Text(), 10, 64)
		if err != nil {
			utils.Dump(err)
		} else {
			pages = append(pages, &app.Page{
				Page: uint(page),
				Link: link,
			})
		}
	})

	currentPage, err := strconv.ParseInt(navPages.Find("strong").First().Text(), 10, 64)
	if err == nil {
		pages = append(pages, &app.Page{
			Page: uint(currentPage),
		})
	}
	sort.Sort(app.Pages(pages))

	app.SuccessResponse(c, map[string]interface{}{
		"devices": devices,
		"pages":   pages,
	})
}
