package main

import (
	"fmt"
	"net/http/httptest"
	"runtime/debug"
	"strings"
	"testing"
)

func AssertEqual(actual, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Errorf("AssertEqual failed.\nExpected: %q\n  Actual: %q\n%v", expected, actual, getDebugData())
	}
}

func AssertEqualWithErrorMessage(actual, expected interface{}, errorMessage string, t *testing.T) {
	if actual != expected {
		t.Errorf(errorMessage, expected, actual, getDebugData())
	}
}

func AssertTrue(b bool, t *testing.T) {
	if !b {
		t.Errorf("Statement was not true\n%v", getDebugData())
	}
}

func AssertFalse(b bool, t *testing.T) {
	if b {
		t.Errorf("Statement was not false\n%v", getDebugData())
	}
}

func AssertError(err error, t *testing.T) {
	if err == nil {
		t.Errorf("Error was not returned\n%v", getDebugData())
	}
}

func AssertNotError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error should not be returned\n%v", getDebugData())
	}
}

func AssertStringContains(actual, expected string, t *testing.T) {
	if !strings.Contains(actual, expected) {
		t.Errorf("The string does not contain the expected substring\nExpected: %v\n  Actual: %v\n%v", expected, actual, getDebugData())
	}
}

func getDebugData() string {
	return fmt.Sprintf("Stack: %v", string(debug.Stack()))
}

func CloseServer(s *httptest.Server) {
	s.Close()
}
