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

//AppConfiguration stores global gonfiguration settings
var AppConfiguration Configuration

//ReadConfig reads Global Configuration
func ReadConfig() error {
	file, err := os.Open("../conf.json")
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfiguration)
}
