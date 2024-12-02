package main

import (
	"helper/parser"
	"reflect"
	"testing"
)

func TestDemoInputCanBeSplit(t *testing.T) {
	ftp := parser.TextFileParser{}
	input, _ := ftp.ParseLinesFromPathAsString("demo.txt")

	slices := getInputLinesAs2DSlices(input)

	if len(slices) != 6 {
		t.Error("Expected 6 slices, got ", len(slices))
	}

	expected := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}, {9, 7, 6, 2, 1}, {1, 3, 2, 4, 5}, {8, 6, 4, 4, 1}, {1, 3, 6, 7, 9}}
	if !reflect.DeepEqual(slices, expected) {
		t.Error("Expected slices to be ", expected, " but got ", slices)
	}
}

func TestDetermineReportIsAscending(t *testing.T) {
	var ascendingReport = []int{1, 2, 3, 4, 5}
	var descendingReport = []int{5, 4, 3, 2, 1}
	var ascendingReportWithSameValues = []int{1, 1, 1, 1, 2}
	var descendingReportWithSameValues = []int{2, 2, 2, 2, 1}

	if !isReportAscending(ascendingReport) {
		t.Error("Expected report to be ascending, but got descending, tested ", ascendingReport)
	}

	if isReportAscending(descendingReport) {
		t.Error("Expected report to be descending, but got ascending, tested ", descendingReport)
	}

	if !isReportAscending(ascendingReportWithSameValues) {
		t.Error("Expected report to be ascending, but got descending, tested ", ascendingReportWithSameValues)
	}

	if isReportAscending(descendingReportWithSameValues) {
		t.Error("Expected report to be descending, but got ascending, tested ", descendingReportWithSameValues)
	}
}

func TestDetermineReportIsSafe(t *testing.T) {
	var safeAscendingReport = []int{1, 2, 3, 4, 5}
	var safeDescendingReport = []int{5, 4, 3, 2, 1}
	var unsafeReportOutOfLimits = []int{1, 2, 7, 8, 10}
	var unsafeReportChangesAscension = []int{12, 8, 7, 6, 3}

	if !isReportSafe(safeAscendingReport, true, 1, 3) {
		t.Error("Expected report to be safe, but got unsafe, tested ", safeAscendingReport)
	}

	if !isReportSafe(safeDescendingReport, false, 1, 3) {
		t.Error("Expected report to be safe, but got unsafe, tested ", safeDescendingReport)
	}

	if isReportSafe(unsafeReportOutOfLimits, true, 1, 3) {
		t.Error("Expected report to be unsafe, but got safe, tested ", unsafeReportOutOfLimits)
	}

	if isReportSafe(unsafeReportChangesAscension, false, 1, 3) {
		t.Error("Expected report to be unsafe, but got safe, tested ", unsafeReportChangesAscension)
	}
}

func TestPart1AgainstDemoInput(t *testing.T) {
	ftp := parser.TextFileParser{}
	input, _ := ftp.ParseLinesFromPathAsString("demo.txt")

	result := part1(input)
	if result != 2 {
		t.Error("Expected 2 as result, got ", result)
	}
}
