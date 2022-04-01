package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marcos-wz/capstone/internal/controller"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
)

func main() {
	// Port
	// NOTE: should i use environment variable, USE VIPER
	port := 8080
	// Repo csv file path
	const csvFile string = "data/fruits.csv"
	// const csvFile string = "data/test/csv/fruits-test-error.csv"
	// NOTE: should i use environment variable, USE VIPER
	// var csvFile = os.Get("CSV_FILE")

	e := echo.New()
	e.Use(middleware.Logger())

	// Creators
	readerRepo := repository.NewReaderRepo(csvFile)
	filterService := service.NewFilterService(readerRepo)
	filterController := controller.NewFilterController(filterService)

	// Handlers
	e.GET("/v1/fruit/:filter/:value", filterController.FilterFruit)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
