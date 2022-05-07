package configs

import (
	"fmt"
	"github.com/marcos-wz/capstone/internal/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServer_Config(t *testing.T) {
	var testCases = []struct {
		name     string
		file     string
		response *entity.ServerConfig
		err      string
	}{
		{
			"should load configs file successfully",
			"./test/server_ok.json",
			&entity.ServerConfig{
				Host: "", Port: 50051, ExternalAPI: "localhost:50051/v1/externalapi/fruits", DebugLevel: 1,
				JSONFile: "./../data/fruits.json",
				CSVFile:  "./../data/fruits.csv",
				SSLCert:  "./../certs/server.crt",
				SSLKey:   "./../certs/server.pem",
			},
			"<nil>",
		},
		{
			"should return json decode errors",
			"./test/server_err-decode.json",
			nil,
			"json: cannot unmarshal string into Go struct field ServerConfig.port of type int",
		},
		{
			"should return validation errors",
			"./test/server_err-validation.json",
			nil,
			"Key: 'ServerConfig.Host' Error:Field validation for 'Host' failed on the 'hostname' tag\nKey: 'ServerConfig.Port' Error:Field validation for 'Port' failed on the 'required' tag\nKey: 'ServerConfig.ExternalAPI' Error:Field validation for 'ExternalAPI' failed on the 'url' tag\nKey: 'ServerConfig.JSONFile' Error:Field validation for 'JSONFile' failed on the 'file' tag\nKey: 'ServerConfig.CSVFile' Error:Field validation for 'CSVFile' failed on the 'file' tag\nKey: 'ServerConfig.SSLCert' Error:Field validation for 'SSLCert' failed on the 'required' tag\nKey: 'ServerConfig.SSLKey' Error:Field validation for 'SSLKey' failed on the 'required' tag",
		},
	}
	// ------------------------------
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := LoadServerConfig(tc.file)
			assert.Equal(t, tc.response, resp)
			assert.Equal(t, tc.err, fmt.Sprintf("%v", err))
		})
	}
}
