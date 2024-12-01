package main

import (
	"helper/parser"
	"reflect"
	"testing"
)

func TestDemoInputCanBeLoaded(t *testing.T) {
	ftp := parser.TextFileParser{}
	_, err := ftp.ParseLinesFromPathAsString("demo.txt")
	if err != nil {
		t.Error("Expected no error when opening file")
	}
}

func TestInputCanBeSorted(t *testing.T) {
	ftp := parser.TextFileParser{}
	content, _ := ftp.ParseLinesFromPathAsString("demo.txt")
	left, right := getSortedSlices(content)

	correctLeft := []int{1, 2, 3, 3, 3, 4}
	correctRight := []int{3, 3, 3, 4, 5, 9}

	if !reflect.DeepEqual(left, correctLeft) {
		t.Error("Expected left slice to be ", correctLeft, " but got ", left)
	}

	if !reflect.DeepEqual(right, correctRight) {
		t.Error("Expected right slice to be ", correctRight, " but got ", right)
	}

}

func TestParts(t *testing.T) {
	ftp := parser.TextFileParser{}
	content, err := ftp.ParseLinesFromPathAsString("demo.txt")
	if err != nil {
		t.Fatal("Expected no error when opening file")
	}

	part1Result := part1(content)

	if part1Result != 11 {
		t.Fatal("[Part 1 demo] Expected 11 as result, got ", part1Result)
	}

	part2Result := part2(content)

	if part2Result != 31 {
		t.Fatal("[Part 2 demo] Expected 33 as result, got ", part2Result)
	}
}
