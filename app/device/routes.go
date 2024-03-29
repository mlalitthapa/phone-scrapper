package device

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/mlalitthapa/phone-scrapper/app"
	"net/http"
	"strings"
)

func Register(r *gin.RouterGroup) {
	r.GET("/latest", LatestDevices)
	r.GET("/device/:slug", Detail)
}

func Detail(c *gin.Context) {
	slug := c.Param("slug")
	doc, err := app.Scrape(slug)
	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	images := make(chan []*Image)
	go Images(images, doc)

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

	device.Images = <-images

	app.SuccessResponse(c, device)
}

func Images(img chan []*Image, doc *goquery.Document) {
	imageWrapper := doc.Find(".article-info .specs-photo-main")
	var images []*Image
	mainImg := imageWrapper.Find("img")
	images = append(images, &Image{
		Name: mainImg.AttrOr("alt", ""),
		Src:  mainImg.AttrOr("src", ""),
	})

	imgLink := imageWrapper.Find("a").First().AttrOr("href", "")
	if imgLink != "" {
		imagesDoc, err := app.Scrape(imgLink)
		if err == nil {
			imagesDoc.Find("#pictures-list img").Each(func(i int, image *goquery.Selection) {
				dataSrc := image.AttrOr("data-src", "")
				images = append(images, &Image{
					Name: image.AttrOr("alt", ""),
					Src:  image.AttrOr("src", dataSrc),
				})
			})
		}
	}

	img <- images
}

// Get latest devices
func LatestDevices(c *gin.Context) {
	// Browse the homepage and get the document
	doc, err := app.Scrape("")
	if err != nil {
		app.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	sidebar := doc.Find("aside.sidebar").First()
	latestModule := sidebar.Find("div.module-phones").First()
	var devices []*app.Device
	latestModule.Find("a.module-phones-link").Each(func(i int, device *goquery.Selection) {
		image := device.Find("img").First()
		devices = append(devices, &app.Device{
			Name: device.Text(),
			Slug: device.AttrOr("href", ""),
			Image: app.DeviceImage{
				Alt: device.Text(),
				Src: image.AttrOr("src", ""),
			},
		})
	})

	app.SuccessResponse(c, devices)
}
