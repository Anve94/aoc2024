package main

import (
	"reflect"
	"testing"
)

func TestDiagonalExtraction(t *testing.T) {
	toTest := []string{
		"abc",
		"def",
		"ghi",
	}
	expected := []string{
		"aei",
		"bf",
		"c",
		"dh",
		"g",
	}

	result := getDiagonals(toTest)
	if reflect.DeepEqual(result, expected) {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestOppositeDiagonalExtraction(t *testing.T) {
	toTest := []string{
		"abc",
		"def",
		"ghi",
	}
	expected := []string{
		"ceg",
		"bd",
		"a",
		"fh",
		"i",
	}

	result := getOppositeDiagonals(toTest)
	if reflect.DeepEqual(result, expected) {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestVerticalsExtraction(t *testing.T) {
	toTest := []string{
		"abc",
		"def",
		"ghi",
	}
	expected := []string{
		"adg",
		"bef",
		"cfi",
	}

	result := getOppositeDiagonals(toTest)
	if reflect.DeepEqual(result, expected) {
		t.Error("Expected ", expected, " but got ", result)
	}
}
