package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/mlalitthapa/phone-scrapper/utils"
	"os"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", os.Getenv("DB_URL"))
	if err != nil {
		return db, err
	}
	utils.Dump("Database connection successful")

	// Enable log mode
	db.LogMode(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "dvc_" + defaultTableName
	}

	return db, nil
}
