package main

import (
	"github.com/joho/godotenv"
	"github.com/mlalitthapa/phone-scrapper/app"
	"github.com/mlalitthapa/phone-scrapper/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app.Initialize(db)

	InitializeRoute()
}
