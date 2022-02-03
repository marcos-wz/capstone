package entity

// ERRORS *****************

type ParseCSVFruitFieldError struct {
	// the index position field on the csv record
	Index int `json:"index"`
	// Field struct to be evaluated
	Field string `json:"field"`
	// Error validation response
	Error string `json:"error"`
	// Set field as required
	Required bool
}

type ParseCVSFruitRecordError struct {
	Record int                       `json:"record"`
	Errors []ParseCSVFruitFieldError `json:"errors"`
}
