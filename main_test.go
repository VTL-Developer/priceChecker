package main

import (
	"io/ioutil"
	"os"
	"runtime"
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
	logInfo("Amount of gochannels start: %v", runtime.NumGoroutine())
}

func teardownFunction() {
	logInfo("Amount of gochannels end: %v", runtime.NumGoroutine())
}
