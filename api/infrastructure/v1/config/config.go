package config

import (
	"encoding/json"
	"os"
)

//Configuration stores configuration
type Configuration struct {
	DbServer string
	DbName   string
}

//ReadConfig reads Global Configuration
func ReadConfig() (Configuration, error) {
	var config Configuration
	file, err := os.Open("conf.json")
	if err != nil {
		return config, err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	return config, err
}
