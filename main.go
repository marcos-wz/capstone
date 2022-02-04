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
	// const csvFile string = "data/fruits-test-error.csv"
	// NOTE: should i use environment variable, USE VIPER
	// var csvFile = os.Get("CSV_FILE")

	e := echo.New()
	e.Use(middleware.Logger())

	// Get Fruits Filter handler
	// NOTE: where should i put creators ?

	repo := repository.NewReaderRepo(csvFile)
	svc := service.NewFilterService(repo)
	ctrl := controller.NewFilterController(svc)
	e.GET("/v1/fruit/:filter/:value", ctrl.FilterFruit)

	e.GET("/v1/fruit/:filter/:value", func(c echo.Context) error {
		repo := repository.NewReaderRepo(csvFile)
		svc := service.NewFilterService(repo)
		ctrl := controller.NewFilterController(svc)
		return ctrl.FilterFruit(c)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
