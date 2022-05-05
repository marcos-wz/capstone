package service

import "github.com/marcos-wz/capstone/internal/entity"

//	Error range: 20 - 29 service errors
const (
	ErrUndefined entity.ErrorType = 0

	// ErrFilterFactory general factory error
	ErrFilterFactory     entity.ErrorType = 20
	ErrFilterFactoryDesc string           = "service filter factory error"

	// ErrFilterFactoryID error factory on the ID
	ErrFilterFactoryID     entity.ErrorType = 21
	ErrFilterFactoryIDDesc string           = "service filter factory ID error"
)
