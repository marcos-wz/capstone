package entity

import "fmt"

//	Error range: 10 - 19 repository errors

// ErrorType holds the error type
type ErrorType uint32

// FruitError system default error definition
type FruitError struct {
	Type ErrorType
	Desc string
	Err  error
}

// Error contract method for Error interface, returns type description and message,
func (fe *FruitError) Error() string {
	return fmt.Sprintf("%v: %v", fe.Desc, fe.Err)
}
