package main

import (
	"testing"
)

func TestGetConfiguration_BasicRead(t *testing.T) {
	config, err := GetConfiguration("./test_configuration.json")

	if err != nil {
		t.Errorf("Error thrown: %v", err)
		return
	}

	if config.Database != "testDatabase" {
		t.Errorf("GetCost returned %q, should be \"testDatabase\"", config.Database)
	}

	if config.Password != "testPassword" {
		t.Errorf("GetCost returned %q, should be \"testPassword\"", config.Password)
	}
}
