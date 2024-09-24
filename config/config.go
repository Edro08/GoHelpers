package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	mapStringName    = "mapString"
	mapIntName       = "mapInt"
	mapInterfaceName = "mapInterface"
)

// IConfig interface defines methods for accessing configuration values
type IConfig interface {
	GetString(keys string) (string, bool)
	GetInt(keys string) (int, bool)
	GetBool(keys string) (bool, bool)
	GetMapInterface(keys string) (map[string]interface{}, bool)
	GetMapString(keys string) (map[string]string, bool)
	GetMapInt(keys string) (map[string]int, bool)
}

type Config struct {
	data map[string]interface{}
}

// NewConfig
// Read and deserialize the YAML file into a variable of type map[string]interface{}
func NewConfig(file string) *Config {
	//
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		panic(err)
	}

	return &Config{data: data}
}
