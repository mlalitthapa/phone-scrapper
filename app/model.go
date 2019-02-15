package app

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Page struct {
	Label string `json:"label"`
	Link  string `json:"link"`
}

func Initialize(db *gorm.DB) {
	DB = db
}
