package controller

import "github.com/marcos-wz/capstone/internal/entity"

// NOTE: Where should put this code ? swagger reference
// utils.go, controller.go, another package, into the same controller ?

// REQUESTS ****************************

// RESPONSES ****************************

// return error response
type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GET FRUIT FILTER ***************
type getFruitsFilterResp struct {
	Code        int            `json:"code"`
	Fruits      []entity.Fruit `json:"fruits"`
	ParserError string         `json:"parser_error"`
}

type getFruitsFilterParams struct {
	// Filter string `param:"filter" validate:"id|name|color|country|all"`
	Filter string `param:"filter" validate:"required"`
	Value  string `param:"value" validate:"required,alphanum"`
}

// RESPONSES ****************************

// // returns json response, and format it if error body object
// func sendJson(w http.ResponseWriter, status int, body interface{}) {
// 	switch v := body.(type) {
// 	case error:
// 		log.Printf("┤ ERROR(%d): %v", status, v)
// 		body = &errorResponse{
// 			Code:    status,
// 			Message: v.Error(),
// 		}
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	err := json.NewEncoder(w).Encode(body)
// 	if err != nil {
// 		log.Printf("error while encoding, %v", err)
// 	}
// }
