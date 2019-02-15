package app

import "github.com/jinzhu/gorm"

var DB *gorm.DB

type Page struct {
	Page uint   `json:"page"`
	Link string `json:"link"`
}

func Initialize(db *gorm.DB) {
	DB = db
}
