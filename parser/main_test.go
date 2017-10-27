package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupFunction()
	retCode := m.Run()
	teardownFunction()
	os.Exit(retCode)
}

func setupFunction() {
	configuration.LogLevel = debugLevel
	logger.SetOutput(ioutil.Discard)
}

func teardownFunction() {
}
