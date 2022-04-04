package entity

// ERRORS *****************

type ParseFruitFieldCSVError struct {
	Field      string `json:"field"`
	Value      string `json:"value"`
	Validation string `json:"validation"`
	Error      string `json:"error"`
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
