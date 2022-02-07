package entity

// PARAMS
type FruitsFilterParams struct {
	Filter string `param:"filter" validate:"oneof=id name color country all"`
	Value  string `param:"value" validate:"required,alphanum"`
}

// ERRORS
type FruitsFilterError struct {
	Type         string
	Error        error
	ParserErrors []ParseFruitRecordCSVError
}
