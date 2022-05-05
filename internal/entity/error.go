package entity

import "fmt"

type ErrorType uint32

type FruitError struct {
	Type     ErrorType
	TypeDesc string
	Message  string
}

// Error contract method for Error interface, returns type description and message,
func (fe *FruitError) Error() string {
	return fmt.Sprintf("%v: %v", fe.TypeDesc, fe.Message)
}
