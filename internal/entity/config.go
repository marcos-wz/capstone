package entity

// ServerConfig entity configuration file with validation rules
type ServerConfig struct {
	Host        string `json:"host" validate:"omitempty,hostname"`
	Port        int    `json:"port" validate:"required,min=1"`
	ExternalAPI string `json:"external-api" validate:"required,url"`
	DebugLevel  uint32 `json:"debug-level"`

	// DATA
	JSONFile string `json:"json-file" validate:"required,file"`
	CSVFile  string `json:"csv-file" validate:"required,file"`

	// SSL CERTIFICATES
	SSLCert string `json:"ssl-cert" validate:"required,file"`
	SSLKey  string `json:"ssl-key" validate:"required,file"`
}

// ClientConfig entity configuration file with validation rules
type ClientConfig struct {
	ServerHost string `json:"server-host" validate:"omitempty,hostname"`
	ServerPort int    `json:"server-port" validate:"required,min=1"`
	Debug      uint32 `json:"debug"`
	SSLCert    string `json:"ssl-cert" validate:"file"`
}
