package main

import (
	"testing"
)

func AssertEqual(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("AssertEqual failed.\nExpected: %q\n  Actual: %q", expected, actual)
	}
}
