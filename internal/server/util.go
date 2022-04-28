package server

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"io/ioutil"
	"log"
)

// LoadServerConfig reads configuration from file and validates the values
func LoadServerConfig(file string) (*entity.ServerConfig, error) {
	// File
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("SERVER-ERROR: reading config file: %q : %v", file, err)
		return nil, err
	}
	// Decode
	config := &entity.ServerConfig{}
	if err := json.Unmarshal(f, config); err != nil {
		log.Printf("SERVER-ERROR: decoding config file: %q : %v", file, err)
		return nil, err
	}
	// Validate Input data
	if err := validator.New().Struct(config); err != nil {
		log.Println("SERVER-ERROR: input validator: ")
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("Field: %v, Value: %v, Tag: %v, Param: %v", e.StructField(), e.Value(), e.Tag(), e.Param())
		}
		return nil, err
	}
	return config, nil
}
