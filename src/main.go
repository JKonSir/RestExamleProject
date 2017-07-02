package main

import (
	"os"
	"storage/postgres/support"
	"app/rest"
	"storage/postgres"
	"services"
)

func main() {
	pgDb := support.PostgresDB{}
	pgDb.InitDB(os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	productStorage := postgres.NewProductStorage(&pgDb)
	productService := services.NewProductService(productStorage)

	a := rest.NewApp(productService)
	a.Run(":8080")
}