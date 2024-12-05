package main

import (
	"github.com/stretchr/testify/assert"
	"helper/parser"
	"testing"
)

func TestCanProcessInputs(t *testing.T) {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")

	parsedOrder, parsedPages := processInputs(demoInput)
	if parsedOrder[0] != "47|53" {
		t.Error("Expected orders to start with 47|53 but got", parsedOrder[0])
	}
	if parsedPages[0] != "75,47,61,53,29" {
		t.Error("Expected pages to start with 75,47,61,53,29 but got", parsedPages[0])
	}
}

func TestCanGenerateHashmap(t *testing.T) {
	orders := []string{"11|1", "11|2", "11|3", "12|4"}
	expected := map[int][]int{11: {1, 2, 3}, 12: {4}}

	if !assert.Equal(t, expected, createHashMap(orders)) {
		t.Error("Expected hashmap to be equal to", expected)
	}
}
