package entity

// ServerConfig entity configuration file with validation rules
type ServerConfig struct {
	Host        string `json:"host" validate:"omitempty,hostname"`
	Port        int    `json:"port" validate:"required,min=1"`
	ExternalAPI string `json:"external-api" validate:"required,url"`
	Debug       bool   `json:"debug"`

	// DATA
	JSONFile string `json:"json-file" validate:"required,file"`
	CSVFile  string `json:"csv-file" validate:"required,file"`
}
