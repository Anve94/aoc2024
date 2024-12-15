package main

import (
	"fmt"
	"helper/parser"
)

type Pos struct {
	y int
	x int
}

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")
	fmt.Println(part1(demoInput))
	fmt.Println(part1(fullInput))
	//fmt.Println(part2(demoInput))
	//fmt.Println(part2(fullInput))
}

func part1(input []string) int {
	antennas := getAntennasFromInput(input)
	antinodes := getAntinodesFromAntennas(antennas)
	var uniqueAntinodes []Pos

	// Remove duplicates
	for _, pos := range antinodes {
		isDuplicate := false
		for _, uniquePos := range uniqueAntinodes {
			if uniquePos.x == pos.x && uniquePos.y == pos.y {
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			uniqueAntinodes = append(uniqueAntinodes, pos)
		}
	}

	maxX, maxY := len(input[0])-1, len(input)-1

	count := 0
	for _, dir := range uniqueAntinodes {
		if dir.x >= 0 && dir.x <= maxX && dir.y >= 0 && dir.y <= maxY {
			count++
		}
	}

	return count
}

func getAntinodesFromAntennas(antennas map[string][]Pos) []Pos {
	var antinodes []Pos
	for _, positions := range antennas {
		for leftIdx, leftPos := range positions {
			for rightIdx, rightPos := range positions {
				if leftIdx != rightIdx {
					antinode := getAntinodeLocation(leftPos, rightPos)
					antinodes = append(antinodes, antinode)
				}
			}
		}
	}
	return antinodes
}

func getAntinodeLocation(primaryNode Pos, otherNode Pos) Pos {
	newX := primaryNode.x + (primaryNode.x - otherNode.x)
	newY := primaryNode.y + (primaryNode.y - otherNode.y)

	return Pos{newY, newX}
}

func getAntennasFromInput(areaMap []string) map[string][]Pos {
	antennaMap := make(map[string][]Pos)
	for y, row := range areaMap {
		for x, c := range row {
			char := string(c)
			if char != "." {
				positions := antennaMap[char]
				positions = append(positions, Pos{y, x})
				antennaMap[char] = positions
			}
		}
	}

	return antennaMap
}
