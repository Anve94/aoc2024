package algorithm

import (
	"testing"
)

func TestCaesarShift(t *testing.T) {
	firstTest := CaesarShift(3, "abc")
	if firstTest != "def" {
		t.Errorf("Expected 'def', got '%s'", firstTest)
	}
}

func TestLoopingCaesarShift(t *testing.T) {
	testCase := CaesarShift(2, "xyz")
	if testCase != "zab" {
		t.Errorf("Expected 'zab', got '%s'", testCase)
	}
}

func TestCaesarShiftWithUpperCaseLetters(t *testing.T) {
	testCase := CaesarShift(3, "ABC")
	if testCase != "DEF" {
		t.Errorf("Expected 'DEF', got '%s'", testCase)
	}
}

func TestCaesarShiftCanHandleNegativeAmounts(t *testing.T) {
	testCase := CaesarShift(-3, "abc")
	if testCase != "xyz" {
		t.Errorf("Expected 'xyz', got '%s'", testCase)
	}
}

func TestCaesarShiftCanHandleMultipleLoops(t *testing.T) {
	testCase := CaesarShift(52, "Hello World!")
	if testCase != "Hello World!" {
		t.Errorf("Expected 'Hello World!', got '%s'", testCase)
	}
}
