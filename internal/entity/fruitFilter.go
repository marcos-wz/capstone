package entity

// PARAMS
type FruitsFilterParams struct {
	Filter string `validate:"oneof=id name color country"`
	Value  string `validate:"required,alphanum"`
}

// ERRORS
type FruitFilterError struct {
	Type         string
	Error        error
	ParserErrors []ParseFruitRecordCSVError
}

// REPONSES
type FruitFilterResponse struct {
	Fruits       []Fruit                    `json:"fruits"`
	ParserErrors []ParseFruitRecordCSVError `json:"parser_errors,omitempty"`
}
