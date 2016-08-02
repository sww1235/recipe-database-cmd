package main

import (
	"encoding/json"
	"io/ioutil"
)

//Configuration stores the configuration that is read in and out from a file
type Configuration struct {
	IPConfig  string `json:"ipconfig"`
	RecipeDir string `json:"recipedir"`
	//not stored, only used internally
	configPath         string
	configFileNotFound bool
}

func readConfig(filename string) (Configuration, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}
	var config Configuration
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		return Configuration{}, err
	}
	return config, nil
}

func writeConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, bytes, 0)
}