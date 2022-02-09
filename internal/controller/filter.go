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

type FilterService interface {
	GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, *entity.FruitFilterError)
}

type filterController struct {
	service FilterService
}

func NewFilterController(svc FilterService) iFilterController {
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
		return c.JSON(http.StatusUnprocessableEntity, &entity.ErrorResponse{
			Message: err.Error(),
		})
	}
	fruits, err := fc.service.GetFilteredFruits(filter)
	// Error validations
	if err != nil {
		switch err.Type {
		// Repository File Error response: internal server error
		case "Repo.FileError":
			return c.JSON(http.StatusInternalServerError, &entity.ErrorResponse{
				Message: err.Error.Error(),
			})
		// Repository parser error response : partial fruits with parser errors
		case "Repo.ParserError":
			return c.JSON(http.StatusPartialContent, &entity.FruitFilterResponse{
				Fruits:       fruits,
				ParserErrors: err.ParserErrors,
			})
		default:
			// Default error response
			return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
				Message: err.Error.Error(),
			})
		}
	}
	// Successful response
	return c.JSON(http.StatusOK, &entity.FruitFilterResponse{
		Fruits: fruits,
	})
}
