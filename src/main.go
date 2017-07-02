package main

import (
	"storage/postgres/support"
	"app/rest"
	"storage/postgres"
	"services"
	"app"
)

func main() {
	conf := app.Conf{}
	conf.ReadConfig()

	pgDb := support.PostgresDB{}
	pgDb.InitDB(conf.Service.Postgres.USERNAME,
		conf.Service.Postgres.PASSWORD,
		conf.Service.Postgres.DB_NAME)
	productStorage := postgres.NewProductStorage(&pgDb)
	productService := services.NewProductService(productStorage)

	a := rest.NewApp(productService)
	a.Run(":8080")
}