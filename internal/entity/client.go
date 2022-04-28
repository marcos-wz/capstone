package entity

// ClientConfig entity configuration file with validation rules
type ClientConfig struct {
	ServerHost string `json:"server-host" validate:"omitempty,hostname"`
	ServerPort int    `json:"server-port" validate:"required,min=1"`
	Debug      bool   `json:"debug"`
}
