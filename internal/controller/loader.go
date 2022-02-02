package controller

import (
	"net/http"
)

// DOMAIN **************************************

type iLoaderController interface {
	// Add new fruits to the service from external API
	LoadFruits(w http.ResponseWriter, r *http.Request)
}

type loaderSvc interface {
	LoadAPIFruits() (int, error)
}

type loaderController struct {
	service loaderSvc
}

func NewLoaderController(svc loaderSvc) iLoaderController {
	return &loaderController{svc}
}

// IMPLEMENTATION ********************************

func (*loaderController) LoadFruits(w http.ResponseWriter, r *http.Request) {
	// total, err := fc.service.LoadFruits()
	// if err != nil {
	// 	sendJson(w, http.StatusBadRequest, err)
	// }
	// // RESPONSE
	// sendJson(w, http.StatusOK, total)
}
