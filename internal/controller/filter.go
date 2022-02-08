package controller

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/marcos-wz/capstone/internal/entity"
)

// DOMAIN **************************************

type iFilterController interface {
	// Get fruits filtered from the service
	FilterFruit(c echo.Context) error
}

type filterService interface {
	GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitsFilterError)
}

type filterController struct {
	service filterService
}

func NewFilterController(svc filterService) iFilterController {
	return &filterController{svc}
}

// IMPLEMENTATION ********************************

func (fc *filterController) FilterFruit(c echo.Context) error {
	// Input Validation
	filter := &entity.FruitsFilterParams{
		Filter: c.Param("filter"),
		Value:  c.Param("value"),
	}
	// Entity Error response: Unprocessable Entity
	if err := validator.New().Struct(filter); err != nil {
		log.Println("ERROR Controller: entity validation - ", err)
		return c.JSON(
			http.StatusUnprocessableEntity,
			map[string]string{"message": err.Error()},
		)
	}
	fruits, err := fc.service.GetFilteredFruits(filter)
	// Error validations
	if err != nil {
		switch err.Type {
		// Repository File Error response: internal server error
		case "Repo.FileError":
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{"message": err.Error.Error()},
			)
		// Repository parser error response : partial content data with errors
		case "Repo.ParserError":
			return c.JSON(
				http.StatusPartialContent,
				map[string]interface{}{"fruits": fruits, "parser_errors": err.ParserErrors},
			)
		default:
			// Default error response
			return c.JSON(
				http.StatusBadRequest,
				map[string]string{"message": err.Error.Error()},
			)
		}
	}
	// Successful response
	return c.JSON(
		http.StatusOK,
		map[string]interface{}{"fruits": fruits},
	)
}
