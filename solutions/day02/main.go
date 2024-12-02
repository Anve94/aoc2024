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
	fmt.Println("Part 2 result:", part2(input))
}

func part1(input []string) int {
	reports := getInputLinesAs2DSlices(input)

	var safeReports int
	for _, report := range reports {
		isAscending := isReportAscending(report)
		if isReportSafe(report, isAscending, 1, 3, false) {
			safeReports++
		}
	}
	return safeReports
}

func part2(input []string) int {
	reports := getInputLinesAs2DSlices(input)

	var safeReports int
	for _, report := range reports {
		isAscending := isReportAscending(report)
		if isReportSafe(report, isAscending, 1, 3, true) {
			safeReports++
		}
	}
	return safeReports
}

func isReportSafe(report []int, isAscending bool, min int, max int, allowUnsafe bool) bool {
	var acceptableReports [][]int
	acceptableReports = append(acceptableReports, report)

	if allowUnsafe {
		// YOLO
		for i := 0; i < len(report); i++ {
			// Slice magic as if I'm writing python again
			acceptableReport := make([]int, 0, len(report)-1)
			acceptableReport = append(acceptableReport, report[:i]...)
			acceptableReport = append(acceptableReport, report[i+1:]...)
			acceptableReports = append(acceptableReports, acceptableReport)
		}
	}

	unsafeReportCount := 0
	for _, acceptableReport := range acceptableReports {
		if allowUnsafe {
			// We need to redetermine if it's ascending or descending since it can have changed
			isAscending = isReportAscending(acceptableReport)
		}
		for i := 1; i < len(acceptableReport); i++ {
			loc1, loc2 := acceptableReport[i-1], acceptableReport[i]
			// If location check doesn't match ascending/descending state anymore, location is invalid
			if isAscending && loc2 < loc1 {
				unsafeReportCount++
				break
			}

			if !isAscending && loc2 > loc1 {
				unsafeReportCount++
				break
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
				unsafeReportCount++
				break
			}
		}

		// If the amount of unsafe reports is the same as our acceptable reports, no reports are valid
		if unsafeReportCount == len(acceptableReports) {
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
