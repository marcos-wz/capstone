package entity

// PARAMS
type FruitsFilterParams struct {
	Filter string `validate:"required,alpha"`
	Value  string `validate:"required,alpha"`
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
