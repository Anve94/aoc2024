package main

import (
	"github.com/stretchr/testify/assert"
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
	if !assert.ElementsMatch(t, result, expected) {
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

	if !assert.ElementsMatch(t, result, expected) {
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
		"beh",
		"cfi",
	}

	result := getVertical(toTest)
	if !assert.ElementsMatch(t, result, expected) {
		t.Error("Expected ", expected, " but got ", result)
	}
}

func TestValidStartingIndexes(t *testing.T) {
	toTest := []string{
		"AvAA",
		"AAfA",
		"AgAA",
		"AAAA",
	}
	expected := [][]int{
		{1, 1},
		{2, 2},
	}

	result := getStartingIndexes(toTest)
	if !assert.ElementsMatch(t, result, expected) {
		t.Error("Expected ", expected, " but got ", result)
	}
}
