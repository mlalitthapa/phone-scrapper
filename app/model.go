package app

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func Initialize(db *gorm.DB) {
	DB = db
}
