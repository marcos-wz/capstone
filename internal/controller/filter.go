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
	GetFruitsFilter(c echo.Context) error
}

type filterSvc interface {
	FilterFruits(filter, value string) ([]entity.Fruit, error)
}

type filterController struct {
	service filterSvc
}

func NewFilterController(svc filterSvc) iFilterController {
	return &filterController{svc}
}

// IMPLEMENTATION ********************************

func (fc *filterController) GetFruitsFilter(c echo.Context) error {
	// Input params validation
	p := &getFruitsFilterParams{
		Filter: c.Param("filter"),
		Value:  c.Param("value"),
	}
	// NOTE: How to bind by PATH
	// if err := c.Bind(p); err != nil {
	// if err := c.PathParamsBinder(p); err != nil {
	// Validator Framework as a Workaround
	if err := validator.New().Struct(p); err != nil {
		log.Println("ERROR CONTROLLER:", err)
		return c.JSON(http.StatusUnprocessableEntity, errorResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: err.Error(),
		})
	}

	fruits, err := fc.service.FilterFruits(p.Filter, p.Value)
	// ERROR RESPONSE VALIDATIONS
	if err != nil {
		// REPOSITORY ERRORS

		// REPOSITORY FILE ERROR
		if strings.HasSuffix(err.Error(), " no such file or directory") {
			return c.JSON(http.StatusInternalServerError, &errorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		// REPOSITORY PARSER ERROR
		// NOTE: fruits with parser errors
		// should return fruit valid fruit list, with default values and excluding invalid records?
		// should i create a parser ERROR to type validation, inteast of string ?
		if strings.HasPrefix(err.Error(), "parser error:") {
			return c.JSON(http.StatusPartialContent, &getFruitsFilterResp{
				Code:        http.StatusPartialContent,
				Fruits:      fruits,
				ParserError: err.Error(),
			})
		}

		// DEFAULT RESPONSE
		return c.JSON(http.StatusBadRequest, &errorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, &getFruitsFilterResp{
		Code:   http.StatusOK,
		Fruits: fruits,
	})
}
