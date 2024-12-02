package main

import (
	"fmt"
	"helper/parser"
	"log"
	"strconv"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}

	input, err := tfp.ParseLinesFromPathAsString("input.txt")
	if err != nil {
		log.Fatal("Could not read from input file")
	}

	fmt.Println("Part 1 result:", part1(input))
	//fmt.Println("Part 2 result:", part2(input))
}

func part1(input []string) int {
	reports := getInputLinesAs2DSlices(input)

	var safeReports int
	for _, report := range reports {
		isAscending := isReportAscending(report)
		if isReportSafe(report, isAscending, 1, 3) {
			safeReports++
		}
	}
	return safeReports
}

func isReportSafe(report []int, isAscending bool, min int, max int) bool {
	for i := 1; i < len(report); i++ {
		loc1, loc2 := report[i-1], report[i]
		// If location check doesn't match ascending/descending state anymore, location is invalid
		if isAscending && loc2 < loc1 {
			return false
		}

		if !isAscending && loc2 > loc1 {
			return false
		}

		// Determine if location values exceed limits
		var minFound, maxFound int
		if loc1 <= loc2 {
			minFound = loc1
			maxFound = loc2
		} else {
			minFound = loc2
			maxFound = loc1
		}

		delta := maxFound - minFound
		if delta < min || delta > max {
			return false
		}
	}
	return true
}

func isReportAscending(report []int) bool {
	for i := 1; i < len(report); i++ {
		if report[i] > report[i-1] {
			return true
		}
	}
	return false
}

func getInputLinesAs2DSlices(input []string) [][]int {
	var reports [][]int
	for _, line := range input {
		report := strings.Split(line, " ")
		var reportLine []int
		for _, char := range report {
			value, _ := strconv.Atoi(char)
			reportLine = append(reportLine, value)
		}
		reports = append(reports, reportLine)
	}
	return reports
}
