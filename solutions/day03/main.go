package main

import (
	"fmt"
	"helper/parser"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")
	part2DemoInput, _ := tfp.ParseLinesFromPathAsString("demo2.txt")

	fmt.Println("Demo result:", part1(demoInput))
	fmt.Println("Full result:", part1(fullInput))
	fmt.Println("Demo 2 result:", part2(part2DemoInput))
	fmt.Println("Full 2 result:", part2(fullInput))

}

func part1(input []string) int {
	total := 0

	for _, line := range input {
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			total += a * b
		}
	}

	return total
}

func part2(input []string) int {
	total := 0

	for _, line := range input {
		doStarts := findAllIndices(line, `do()`, true)
		dontStarts := findAllIndices(line, `don't()`, false)
		// Get match results
		re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
		matches := re.FindAllStringSubmatch(line, -1)
		// Get match indexes
		matchIndexes := re.FindAllStringSubmatchIndex(line, -1)

		for idx, match := range matches {
			startIndex := matchIndexes[idx][0]

			if canDoOperation(doStarts, dontStarts, startIndex) {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				total += a * b
			}

		}
	}

	return total
}

func canDoOperation(doStarts []int, dontStarts []int, n int) bool {
	size := 0
	if len(doStarts) > len(dontStarts) {
		size = len(doStarts)
	} else {
		size = len(dontStarts)
	}

	closestDoStart := 0
	closestDontStart := 0
	for i := 0; i < size; i++ {
		if i < len(doStarts) && doStarts[i] < n {
			closestDoStart = doStarts[i]
		}
		if i < len(dontStarts) && dontStarts[i] < n {
			closestDontStart = dontStarts[i]
		}
	}

	return closestDoStart >= closestDontStart
}

func findAllIndices(slice, substr string, appendStartAsValid bool) []int {
	var indices []int

	if appendStartAsValid {
		indices = append(indices, 0)
	}

	index := strings.Index(slice, substr)

	for index != -1 {
		indices = append(indices, index)
		index = strings.Index(slice[index+1:], substr)
		if index != -1 {
			index += indices[len(indices)-1] + 1
		}
	}

	return indices
}
