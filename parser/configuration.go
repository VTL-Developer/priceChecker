package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configuration struct {
	Database string
	Password string
}

func GetConfiguration(filename string) (*Configuration, error) {
	var configuration Configuration
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &configuration)

	if err != nil {
		return nil, err
	}

	return &configuration, err
}
