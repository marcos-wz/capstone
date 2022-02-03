package entity

// ERRORS *****************

type ParseFruitCSVError struct {
	// the index position field on the csv record
	Index int
	// Field struct to be evaluated
	Field string
	// Error validation response
	Error error
}
