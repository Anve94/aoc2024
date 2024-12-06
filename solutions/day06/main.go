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

var initialRoute map[Position]int

type Position struct {
	col int // Col is the x coordinates
	row int // Row is the y coordinates - I hope
}

type Offset struct {
	col int
	row int
}

type Guard struct {
	walkedPositions  map[Position]int
	currentPosition  Position
	currentDirection string
}

func main() {
	tfp := parser.TextFileParser{}

	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")

	fmt.Println("Demo result:", part1(demoInput))
	fmt.Println("Demo result 2:", part2FiveHead(demoInput))
	fmt.Println("Full result:", part1(fullInput))
	fmt.Println("Full result:", part2FiveHead(fullInput))
}

func part1(input []string) int {
	rowStart, colStart, startChar := findStartingPosition(input)

	guard := Guard{
		walkedPositions:  make(map[Position]int),
		currentPosition:  Position{col: colStart, row: rowStart},
		currentDirection: startChar,
	}

	guard.walkedPositions[guard.currentPosition] = 1
	for {
		if canMoveInCurrentDirection(input, guard) {
			move(&guard)
		} else {
			turn(&guard)
		}

		if isExitingMap(input, guard) {
			break
		}
	}

	initialRoute = deepCopyMap(guard.walkedPositions)

	return len(guard.walkedPositions)
}

func part2FiveHead(input []string) int {
	// We start sane, start position is the same for every map permutation.
	rowStart, colStart, startChar := findStartingPosition(input)

	// We continue less sane by generating a shit ton of starting grid permutations.
	startingMaps := getStartGridPermutations(initialRoute, input)

	// We go off the chart and go mad.
	// Assuming there are NO obstacles and every position is visited once for EVERY direction option, there is
	// a theoretical MAX of cycles we need to test against. If the threshold is exceeded we are in an infinite loop. The sane
	// approach would be to test whether any new position is already known with the same direction, but this is day 6.
	// This 5 head move algorithm will just keep on spinning until the max is reached by just running for a theoretical
	// maximum amount of cycles. Assuming visiting every place once and in 4 direction, it would be
	THEORETICAL_MAX_CYCLES := len(input) * len(input[0]) * 4 // This variable is screaming purposefully

	// It would be possible to reduce THEORETICAL_MAX_CYCLES with the known amount of obstacles, but I am not going to
	// since I don't want to.

	loopsFound := 0

	for _, startingMap := range startingMaps {
		hasTheGuardEscaped := false

		// YOLO
		guard := Guard{
			walkedPositions:  make(map[Position]int),
			currentPosition:  Position{col: colStart, row: rowStart},
			currentDirection: startChar,
		}

		for i := 0; i <= THEORETICAL_MAX_CYCLES; i++ {
			if canMoveInCurrentDirection(startingMap, guard) {
				move(&guard)
			} else {
				turn(&guard)
			}

			if isExitingMap(startingMap, guard) {
				hasTheGuardEscaped = true
				break
			}
		}

		if !hasTheGuardEscaped {
			loopsFound++
		}
	}

	return loopsFound
}

func getStartGridPermutations(initialRoute map[Position]int, originalMap []string) [][]string {
	// We only put obstacles on the known path from part 1, since placing it in any other place the guard would never
	// bump into the obstacle.
	var permutations [][]string
	for position, _ := range initialRoute {
		rowIndex, colIndex := position.row, position.col

		if originalMap[rowIndex][colIndex] == '#' {
			continue
		}

		newMap := make([]string, len(originalMap))
		// I've played enough Runescape to know how tedious runes are but at least I prepared myself for the future
		// Strings are immutable, apparently, so we build all the rows for all the maps from scratch too in O(n*m)
		// for every map instead of injecting it in O(1)
		for curRowIndex, _ := range originalMap {
			var newRow []rune
			for curColIndex, char := range originalMap[curRowIndex] {
				if rowIndex == curRowIndex && colIndex == curColIndex {
					newRow = append(newRow, '#')
				} else {
					newRow = append(newRow, char)
				}
			}
			newMap[curRowIndex] = string(newRow)
		}

		permutations = append(permutations, newMap)
	}

	return permutations
}

func deepCopyMap(original map[Position]int) map[Position]int {
	newMap := make(map[Position]int)
	for key, value := range original {
		newMap[key] = value
	}
	return newMap
}

func move(guard *Guard) {
	row, col := guard.currentPosition.row, guard.currentPosition.col
	offset := DirectionMap[guard.currentDirection]
	nextRow, nextCol := row+offset.row, col+offset.col

	guard.currentPosition.col = nextCol
	guard.currentPosition.row = nextRow
	guard.walkedPositions[guard.currentPosition] = 1
}

func turn(guard *Guard) {
	guard.currentDirection = getNextDirection(guard.currentDirection)
}

func isExitingMap(grid []string, guard Guard) bool {
	row, col := guard.currentPosition.row, guard.currentPosition.col
	offset := DirectionMap[guard.currentDirection]
	nextCol, nextRow := col+offset.col, row+offset.row
	if nextCol < 0 || nextCol >= len(grid[0]) || nextRow < 0 || nextRow >= len(grid) {
		return true
	}

	return false
}

func canMoveInCurrentDirection(grid []string, guard Guard) bool {
	row, col := guard.currentPosition.row, guard.currentPosition.col

	offset := DirectionMap[guard.currentDirection]
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
