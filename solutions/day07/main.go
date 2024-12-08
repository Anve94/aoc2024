package main

import (
	"fmt"
	"helper/parser"
	"strconv"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")
	fmt.Println(part1(demoInput))
	fmt.Println(part1(fullInput))
	fmt.Println(part2(demoInput))
	fmt.Println(part2(fullInput))
}

func part1(input []string) int {
	sum := 0

	testValues, calibrationValues := extractValuesFromInput(input)
	for i := 0; i < len(testValues); i++ {
		testValue := testValues[i]
		if isValid(testValue, calibrationValues[i], false) {
			sum += testValue
		}
	}

	return sum
}

func part2(input []string) int {
	sum := 0

	testValues, calibrationValues := extractValuesFromInput(input)
	for i := 0; i < len(testValues); i++ {
		testValue := testValues[i]
		if isValid(testValue, calibrationValues[i], true) {
			sum += testValue
		}
	}

	return sum
}

func isValid(testValue int, calibrationValues []int, withConcatenation bool) bool {
	if len(calibrationValues) < 2 {
		// No need to check since we just test immediately if they are the same
		return calibrationValues[0] == testValue
	}

	// I guess this is how it works looking at fib: https://gobyexample.com/recursion
	var recurse func(idx, currentValue int) bool
	recurse = func(idx, currentValue int) bool {
		if idx == len(calibrationValues) {
			// Exit condition once we've iterated over all calibration values
			return currentValue == testValue
		}

		if withConcatenation {
			currentAsString := strconv.Itoa(currentValue)
			nextAsString := strconv.Itoa(calibrationValues[idx])
			newValue, _ := strconv.Atoi(currentAsString + nextAsString)

			if recurse(idx+1, newValue) {
				return true
			}
		}

		if recurse(idx+1, currentValue+calibrationValues[idx]) {
			return true
		}

		if recurse(idx+1, currentValue*calibrationValues[idx]) {
			return true
		}

		return false
	}

	return recurse(1, calibrationValues[0])
}

func extractValuesFromInput(input []string) ([]int, [][]int) {
	testValues, calibrationValues := make([]int, 0), make([][]int, 0)
	for _, line := range input {
		exploded := strings.Split(line, ": ")
		mapKey, _ := strconv.Atoi(exploded[0])
		testValues = append(testValues, mapKey)

		valuesToAppend := make([]int, 0)
		for _, v := range strings.Split(exploded[1], " ") {
			calibrationValue, _ := strconv.Atoi(v)
			valuesToAppend = append(valuesToAppend, calibrationValue)
		}
		calibrationValues = append(calibrationValues, valuesToAppend)
	}
	return testValues, calibrationValues
}
