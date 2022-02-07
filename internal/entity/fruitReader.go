package entity

// ERRORS *****************

type ParseFruitFieldCSVError struct {
	Field      string
	Value      string
	Validation string
	Error      string
}

type ParseFruitRecordCSVError struct {
	Record int                       `json:"record"`
	Errors []ParseFruitFieldCSVError `json:"errors"`
}

type ReadFruitsError struct {
	Type         string
	Error        error
	ParserErrors []ParseFruitRecordCSVError
}
