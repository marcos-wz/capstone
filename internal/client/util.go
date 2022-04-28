package client

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"io/ioutil"
	"log"
)

// LoadClientConfig reads configuration from file and validates the values
func LoadClientConfig(file string) (*entity.ClientConfig, error) {
	// File
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("CLIENT-ERROR: reading config file: %q : %v", file, err)
		return nil, err
	}
	// Decode
	config := &entity.ClientConfig{}
	if err := json.Unmarshal(f, config); err != nil {
		log.Printf("CLIENT-ERROR: decoding config file: %q : %v", file, err)
		return nil, err
	}

	// Validate Input data
	if err := validator.New().Struct(config); err != nil {
		log.Println("CLIENT-ERROR: input validator: ")
		for _, e := range err.(validator.ValidationErrors) {
			log.Printf("Field: %v, Value: %v, Tag: %v, Param: %v", e.StructField(), e.Value(), e.Tag(), e.Param())
		}
		return nil, err
	}
	return config, nil
}
