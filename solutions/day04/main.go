package main

import (
	"fmt"
	"helper/parser"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")

	fmt.Println("Demo result:", part1(demoInput))
	fmt.Println("Full result:", part1(fullInput))
	//fmt.Println("Demo 2 result:", part2(part2DemoInput))
	//fmt.Println("Full 2 result:", part2(fullInput))

}

func part1(input []string) int {
	var lines []string
	lines = append(lines, input...)
	lines = append(lines, getDiagonals(input)...)
	lines = append(lines, getOppositeDiagonals(input)...)
	lines = append(lines, getVertical(input)...)

	count := 0
	for _, line := range lines {
		count += strings.Count(line, "XMAS")
		count += strings.Count(line, "SAMX")
	}

	return count
}

func getDiagonals(lines []string) []string {
	var diagonals []string
	n := len(lines)

	for startRow := 0; startRow < n; startRow++ {
		var diag []byte
		for i, j := startRow, 0; i < n && j < n; i, j = i+1, j+1 {
			diag = append(diag, lines[i][j])
		}
		diagonals = append(diagonals, string(diag))
	}

	for startCol := 1; startCol < n; startCol++ {
		var diag []byte
		for i, j := 0, startCol; i < n && j < n; i, j = i+1, j+1 {
			diag = append(diag, lines[i][j])
		}
		diagonals = append(diagonals, string(diag))
	}

	return diagonals
}

func getOppositeDiagonals(lines []string) []string {
	// Entire lines to make it easier to work with
	var diagonals []string
	// For boundary check
	n := len(lines)

	for startRow := 0; startRow < n; startRow++ {
		var diag []byte
		for i, j := startRow, n-1; i < n && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, lines[i][j])
		}
		diagonals = append(diagonals, string(diag))
	}

	for startCol := n - 2; startCol >= 0; startCol-- {
		var diag []byte
		for i, j := 0, startCol; i < n && j >= 0; i, j = i+1, j-1 {
			diag = append(diag, lines[i][j])
		}
		diagonals = append(diagonals, string(diag))
	}

	return diagonals
}

func getVertical(lines []string) []string {
	numColumns := len(lines[0])
	numRows := len(lines)

	verticals := make([]string, numColumns)

	for col := 0; col < numColumns; col++ {
		var verticalLine []byte
		for row := 0; row < numRows; row++ {
			if col < len(lines[row]) {
				verticalLine = append(verticalLine, lines[row][col])
			}
		}
		verticals[col] = string(verticalLine)
	}

	return verticals
}
