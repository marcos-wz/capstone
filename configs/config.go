package configs

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/marcos-wz/capstone/internal/repository"
	"github.com/marcos-wz/capstone/internal/service"
	"io/ioutil"
	"log"
)

// LoadServerConfig reads configuration from file and validates the values
func LoadServerConfig(file string) (*entity.ServerConfig, error) {

	// File
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("SERVER-ERROR: reading configs file: %q : %v", file, err)
		return nil, err
	}
	// Decode
	config := &entity.ServerConfig{}
	if err := json.Unmarshal(f, config); err != nil {
		log.Printf("SERVER-ERROR: decoding configs file: %q : %v", file, err)
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
	// Set debugging verbose level
	if config.DebugLevel > 0 {
		log.Printf("Debugging level %d...", config.DebugLevel)
		service.DebugLevel = config.DebugLevel
		repository.DebugLevel = config.DebugLevel
	}
	return config, nil
}

// LoadClientConfig reads configuration from file and validates the values
func LoadClientConfig(file string) (*entity.ClientConfig, error) {
	// File
	f, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("CLIENT-ERROR: reading configs file: %q : %v", file, err)
		return nil, err
	}
	// Decode
	config := &entity.ClientConfig{}
	if err := json.Unmarshal(f, config); err != nil {
		log.Printf("CLIENT-ERROR: decoding configs file: %q : %v", file, err)
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
