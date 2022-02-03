package entity

// PARAMS
type FruitsFilterParams struct {
	Filter string `param:"filter" validate:"oneof=id name color country all"`
	// Filter string `param:"filter" validate:"required,alpha"`
	Value string `param:"value" validate:"omitempty,alphanum"`
}

// RESPONSES
type FruitsFilterResponse struct {
	Fruits      []Fruit `json:"fruits"`
	ParserError string  `json:"parser_error"`
}
