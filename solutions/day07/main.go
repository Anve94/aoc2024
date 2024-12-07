package main

import (
	"fmt"
	"helper/parser"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fmt.Println(part1(demoInput))
}

func part1(input []string) int {
	sum := 0

	//testValues, values := extractValuesFromInput(input)
	extractValuesFromInput(input)

	return sum
}

func extractValuesFromInput(input []string) ([]int, [][]int) {
	for _, line := range input {
		splitLine := strings.Split(line, ":")
		l, r := splitLine[0], splitLine[1]
		fmt.Println(l, r)
	}

	return []int{1}, [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
}
