package main

import (
	"encoding/json"
	"io/ioutil"
)

var configuration Configuration

type Configuration struct {
	Database string
	Password string
	LogLevel logLevel
}

func LoadConfiguration(filename string) {
	configuration.LogLevel = infoLevel
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		logFatal("Unable to read the configuration file \"%v\": %v", filename, err)
		panic(err)
	}

	err = json.Unmarshal(content, &configuration)

	if err != nil {
		logFatal("Unable to parse the configuration file \"%v\": %v", filename, err)
		panic(err)
	}

	if len(configuration.Database) == 0 {
		logFatal("Database was not defined.")
		panic("Database was not defined.")
	}

	if len(configuration.Password) == 0 {
		logFatal("Password was not defined.")
		panic("Password was not defined.")
	}

	if configuration.LogLevel == noLevel {
		configuration.LogLevel = infoLevel
	}
}
