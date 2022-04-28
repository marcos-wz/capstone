package entity

// FruitCSVRecord entity with string values, and validation rules for csv parser
type FruitCSVRecord struct {
	Id           string `validate:"required,numeric,ne=0"`
	Name         string `validate:"required,printascii,gt=2"`
	Description  string `validate:"required,printascii,gt=2"`
	Color        string `validate:"required,printascii,gt=2"`
	Unit         string `validate:"oneof=KG LB"`
	Price        string `validate:"required,numeric,ne=0.00,ne=0"`
	Currency     string `validate:"oneof=MXN BRL CAD USD"`
	Stock        string `validate:"required,numeric"`
	CaducateDays string `validate:"required,numeric,ne=0"`
	Country      string `validate:"oneof=MEXICO BRAZIL CANADA USA"`
	CreateTime   string `validate:"required,numeric"`
	UpdateTime   string `validate:"omitempty,numeric"`
}
