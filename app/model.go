package app

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Page struct {
	Page uint   `json:"page"`
	Link string `json:"link"`
}

type Pages []*Page

func (p Pages) Len() int {
	return len(p)
}

func (p Pages) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pages) Less(i, j int) bool {
	return p[i].Page < p[j].Page
}

type DeviceImage struct {
	Src string `json:"src"`
	Alt string `json:"alt"`
}

type Device struct {
	Name  string      `json:"name"`
	Slug  string      `json:"slug"`
	Image DeviceImage `json:"image"`
}

func Initialize(db *gorm.DB) {
	DB = db
}
