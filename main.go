package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marcos-wz/capstone/internal/controller"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
)

func main() {

	// Repo csv file path
	const csvFile string = "data/fruits.csv"
	// const csvFile string = "data/test/csv/fruits-test-error.csv"
	// var csvFile = os.Get("CSV_FILE") // NOTE: should i use environment variable, USE VIPER

	e := echo.New()
	e.Use(middleware.Logger())

	// Creators
	readerRepo := repository.NewReaderRepo(csvFile)
	filterService := service.NewFilterService(readerRepo)
	filterController := controller.NewFilterController(filterService)

	// Handlers
	e.GET("/v1/fruit/:filter/:value", filterController.FilterFruit)

	e.Logger.Fatal(e.Start(":8080"))
}
