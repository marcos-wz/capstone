package service

import "github.com/marcos-wz/capstone/internal/entity"

//	Error range: 20 - 29 service errors
const (
	ErrUndefined entity.ErrorType = 0

	// ErrFilterFactory general factory error
	ErrFilterFactory entity.ErrorType = 20

	// ErrFilterFactoryID error factory on the ID
	ErrFilterFactoryID entity.ErrorType = 21
)

var ErrDesc = map[entity.ErrorType]string{
	ErrUndefined:       "undefined error",
	ErrFilterFactory:   "service filter factory error",
	ErrFilterFactoryID: "service filter factory ID error",
}
