package main

import (
	"testing"
)

var origConfiguration Configuration

func configurationTestSetup() func() {
	origConfiguration = configuration
	return func() {
		configuration = origConfiguration
	}
}

func TestGetConfiguration_BasicRead(t *testing.T) {
	teardown := configurationTestSetup()
	defer teardown()

	LoadConfiguration("./test/test_configuration.json")

	AssertEqual(configuration.Database, "testDatabase1", t)
	AssertEqual(configuration.Password, "testPassword1", t)
	AssertEqual(configuration.LogLevel, infoLevel, t)
}

func TestGetConfiguration_ConfigurationWithLogLevel(t *testing.T) {
	teardown := configurationTestSetup()
	defer teardown()

	LoadConfiguration("./test/test_configuration_with_log_level.json")

	AssertEqual(configuration.Database, "testDatabase2", t)
	AssertEqual(configuration.Password, "testPassword2", t)
	AssertEqual(configuration.LogLevel, errorLevel, t)
}

func TestGetConfiguration_Suite(t *testing.T) {
	cases := []struct {
		name         string
		fileName     string
		errorMessage string
	}{
		{"Configuration File Not Found", "./test/no_configuration.json", "open ./test/no_configuration.json:"},
		{"Malformed Configuration File", "./test/malformed_configuration.json", "invalid character '}' looking for beginning of object key string"},
		{"No Database in configuration", "./test/no_database_configuration.json", "Database was not defined."},
		{"No password in configuration", "./test/no_password_configuration.json", "Password was not defined."},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			teardown := configurationTestSetup()
			defer teardown()
			defer assertPanic(t, tc.errorMessage)
			LoadConfiguration(tc.fileName)
		})
	}
}

func TestGetConfiguration_ConfigurationNoDatabase(t *testing.T) {
}

func assertPanic(t *testing.T, errorMessage string) {
	r := recover()
	if r == nil {
		t.Error("The code did not panic")
	}

	actualErrorMessage, _ := r.(string)
	err, ok := r.(error)

	if ok {
		actualErrorMessage = err.Error()
	}

	AssertStringContains(actualErrorMessage, errorMessage, t)
}
