package controller

import (
	"log"
	"net/http"
	"strings"

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
	GetFilteredFruits(filter *entity.FruitsFilterParams) ([]entity.Fruit, error)
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
	if err := validator.New().Struct(filter); err != nil {
		log.Println("ERROR CONTROLLER:", err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	fruits, err := fc.service.GetFilteredFruits(filter)
	// Error response validation
	if err != nil {
		// Repository file error : internal server
		if strings.HasSuffix(err.Error(), "no such file or directory") {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		// Return valid fruits and the parser error
		// Repository parser error: partial content (lost data)
		if strings.HasPrefix(err.Error(), "parser error: ") {
			return c.JSON(http.StatusPartialContent, &entity.FruitsFilterResponse{
				Fruits:      fruits,
				ParserError: err.Error(),
			})
		}
		// Default response
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, fruits)
}
