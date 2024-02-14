/*
Package config contains the config file reader and the config struct.
*/
package config

/*
   return stuct of a yaml config file
*/

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

/*
	Example of a config file
	server:
		port: ":8888"
*/

// AppConf contains all main structs.
type AppConf struct {
	Server  ServerConfig `yaml:"server"`
	CintAPI CintAPI      `yaml:"cintAPI"`
}

// ServerConfig contains the data for the micro servise server.
type ServerConfig struct {
	Port    string `yaml:"port"`
	TimeOut int    `yaml:"timeOut"`
}

// Dalle contains the data for the dalle client.
type CintAPI struct {
	URL string `yaml:"url"`
	Key string
}

// GetAppConfig reads a spefic file and return the yaml format of it
// return ServerConfig struct yaml format of the config file.
func GetAppConfig(path string) (*AppConf, error) {
	var appConfig AppConf

	yamlFile, openfileError := os.ReadFile(filepath.Clean(path))
	if openfileError != nil {
		return nil, fmt.Errorf("internal_config_GetAppConfig_open_file %w", openfileError)
	}

	err := yaml.Unmarshal(yamlFile, &appConfig)
	if err != nil {
		return nil, fmt.Errorf("internal_config_GetAppConfig_yaml_unmarshal %w", err)
	}

	appConfig.CintAPI.Key = os.Getenv("CINT_API_KEY")

	if appConfig.CintAPI.Key == "" {
		return nil, fmt.Errorf("CINT_API_KEY environment variable is not set")
	}

	return &appConfig, nil
}
