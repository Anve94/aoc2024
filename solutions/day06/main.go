package main

import (
	"fmt"
	"helper/parser"
)

const ValidDirections = ">v<^"

var DirectionMap = map[string]Offset{
	"<": {col: -1, row: 0},
	">": {col: 1, row: 0},
	"v": {col: 0, row: 1},
	"^": {col: 0, row: -1},
}

type Position struct {
	col int // Col is the x coordinates
	row int // Row is the y coordinates - I hope
}

type Offset struct {
	col int
	row int
}

type Historian struct {
	walkedPositions  map[Position]int
	currentPosition  Position
	currentDirection string
}

func main() {
	tfp := parser.TextFileParser{}

	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")

	fmt.Println("Demo result:", part1(demoInput))
	fmt.Println("Full result:", part1(fullInput))
}

func part1(input []string) int {
	rowStart, colStart, startChar := findStartingPosition(input)

	historian := Historian{
		walkedPositions:  make(map[Position]int),
		currentPosition:  Position{col: colStart, row: rowStart},
		currentDirection: startChar,
	}

	historian.walkedPositions[historian.currentPosition] = 1
	for {
		if canMoveInCurrentDirection(input, historian) {
			move(&historian)
		} else {
			turn(&historian)
		}

		if isExitingMap(input, historian) {
			break
		}
	}

	return len(historian.walkedPositions)
}

func move(historian *Historian) {
	row, col := historian.currentPosition.row, historian.currentPosition.col
	offset := DirectionMap[historian.currentDirection]
	nextRow, nextCol := row+offset.row, col+offset.col

	historian.currentPosition.col = nextCol
	historian.currentPosition.row = nextRow
	historian.walkedPositions[historian.currentPosition] = 1
}

func turn(historian *Historian) {
	historian.currentDirection = getNextDirection(historian.currentDirection)
}

func isExitingMap(grid []string, historian Historian) bool {
	row, col := historian.currentPosition.row, historian.currentPosition.col
	offset := DirectionMap[historian.currentDirection]
	nextCol, nextRow := col+offset.col, row+offset.row
	if nextCol < 0 || nextCol >= len(grid[0]) || nextRow < 0 || nextRow >= len(grid) {
		return true
	}

	return false
}

func canMoveInCurrentDirection(grid []string, historian Historian) bool {
	row, col := historian.currentPosition.row, historian.currentPosition.col

	offset := DirectionMap[historian.currentDirection]
	nextCol, nextRow := col+offset.col, row+offset.row
	if nextCol < 0 || nextCol >= len(grid[0]) || nextRow < 0 || nextRow >= len(grid) || string(grid[nextRow][nextCol]) == "#" {
		return false
	}

	return true
}

func getNextDirection(currentDirection string) string {
	currentIndex := 0
	for idx, direction := range ValidDirections {
		if string(direction) == currentDirection {
			currentIndex = idx + 1
		}
	}
	if currentIndex >= len(ValidDirections) {
		currentIndex = 0
	}

	return string(ValidDirections[currentIndex])
}

func findStartingPosition(input []string) (int, int, string) {
	for i, row := range input {
		for j, char := range row {
			if char == '<' || char == '>' || char == '^' || char == 'v' {
				return i, j, string(char)
			}
		}
	}

	return -1, -1, ""
}
