package main

import (
	"anve/algorithm"
	"fmt"
	"helper/parser"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	tfp := parser.TextFileParser{}

	demoInput, err := tfp.ParseLinesFromPathAsString("demo.txt")
	if err != nil {
		log.Fatal("Could not read from input file")
	}

	demoResult := part1(demoInput)
	fmt.Println("Demo result:", demoResult)

	fullInput, err := tfp.ParseLinesFromPathAsString("input.txt")
	if err != nil {
		log.Fatal("Could not read from input file")
	}

	fullResult := part1(fullInput)
	fmt.Println("Full result:", fullResult)

	demo2Result := part2(demoInput)
	fmt.Println("Demo 2 result:", demo2Result)

	full2Result := part2(fullInput)
	fmt.Println("Full 2 result:", full2Result)
}

func part1(input []string) int {
	left, right := getSortedSlices(input)
	sort.Ints(left)
	sort.Ints(right)
	var total int

	for i := 0; i < len(left); i++ {
		delta := math.Abs(float64(left[i] - right[i]))
		// Technically incorrect since it would give wrong results for > 2^53 and < 2^-53 but eh, I can see the input
		total += int(delta)
	}

	return total
}

func part2(input []string) int {
	left, right := getSortedSlices(input)
	sort.Ints(left)
	sort.Ints(right)
	var total int

	for i := 0; i < len(left); i++ {
		num := left[i]
		count := algorithm.Count(right, func(x int) bool {
			return x == num
		})
		total += count * num
	}

	return total
}

func getSortedSlices(input []string) ([]int, []int) {
	var left []int
	var right []int

	for _, line := range input {
		re := regexp.MustCompile("\\d+")
		matches := re.FindAllString(line, -1)

		// Push to left slice
		num, _ := strconv.Atoi(matches[0])
		left = append(left, num)

		// Push to right slice
		num, _ = strconv.Atoi(matches[1])
		right = append(right, num)
	}

	return left, right
}
