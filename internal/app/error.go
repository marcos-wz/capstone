package app

import "fmt"

type ErrorType uint32

type FruitError struct {
	eType   ErrorType
	message string
}

// NewFruitError return a FruitError instance
func NewFruitError(eType ErrorType, message string) error {
	err := &FruitError{}
	err.setType(eType)
	err.setMessage(message)
	return err
}

// Fruit type errors definitions
// Error range:
//	- 0 - 9 System
//	- 10 - 19 repository errors
//	- 20 - 29 service errors
//	- 30 - 39 server errors
const (
	ErrUndefined     ErrorType = 0
	ErrParseFruitCSV ErrorType = 1

	ErrRepoFileCSV  ErrorType = 10
	ErrRepoFileJSON ErrorType = 11

	ErrSVCFilterFactory   ErrorType = 20
	ErrSVCFilterFactoryID ErrorType = 21
)

// Description maps for ErrorType
var errDesc = map[ErrorType]string{
	ErrUndefined:     "undefined error",
	ErrParseFruitCSV: "parse fruit error",

	ErrRepoFileCSV:  "csv file error",
	ErrRepoFileJSON: "json file error",

	ErrSVCFilterFactory:   "service filter factory error",
	ErrSVCFilterFactoryID: "service filter ID factory error",
}

// Type returns the type of the error
func (fe *FruitError) Type() ErrorType {
	return fe.eType
}

// TypeDesc returns the type description, usually used for building errors
func (fe *FruitError) TypeDesc() string {
	return errDesc[fe.eType]
}

// Error returns a string message of key and message
func (fe *FruitError) Error() string {
	return fmt.Sprintf("%v: %v", fe.TypeDesc(), fe.message)
}

// setType set the key param into the struct. If not found set undefined
func (fe *FruitError) setType(eType ErrorType) {
	switch eType {
	case ErrRepoFileCSV, ErrRepoFileJSON, ErrParseFruitCSV:
		fe.eType = eType
	default:
		fe.eType = ErrUndefined
	}
}

// setValue set the error message
func (fe *FruitError) setMessage(message string) {
	fe.message = message
}
