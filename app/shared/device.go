package shared

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/mlalitthapa/phone-scrapper/app"
)

func GetDeviceList(doc *goquery.Document) []*app.Device {
	var devices []*app.Device
	doc.Find("#review-body ul li a").Each(func(i int, d *goquery.Selection) {
		name := d.Find("strong span").Text()
		slug, _ := d.Attr("href")
		image := d.Find("img")
		imageLink, _ := image.Attr("src")
		imageAlt, _ := image.Attr("title")
		devices = append(devices, &app.Device{
			Name: name,
			Slug: slug,
			Image: app.DeviceImage{
				Src: imageLink,
				Alt: imageAlt,
			},
		})
	})
	return devices
}
