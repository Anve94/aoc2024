package main

import (
	"fmt"
	"helper/parser"
)

type Pos struct {
	y int
	x int
}

var maxX, maxY int

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")
	fmt.Println(solve(demoInput, false))
	fmt.Println(solve(fullInput, false))
	fmt.Println(solve(demoInput, true))
	fmt.Println(solve(fullInput, true))
}

func solve(input []string, withResonantHarmonics bool) int {
	maxX, maxY = len(input[0])-1, len(input)-1

	antennas := getAntennasFromInput(input)
	antinodes := getAntinodesFromAntennas(antennas, withResonantHarmonics)
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

	count := 0
	for _, dir := range uniqueAntinodes {
		if dir.x >= 0 && dir.x <= maxX && dir.y >= 0 && dir.y <= maxY {
			count++
		}
	}

	return count
}

func getAntinodesFromAntennas(antennas map[string][]Pos, withResonantHarmonics bool) []Pos {
	var antinodes []Pos
	for _, positions := range antennas {
		for leftIdx, primary := range positions {
			for rightIdx, other := range positions {
				if leftIdx != rightIdx {
					antinode := getAntinodeLocation(primary, other)
					antinodes = append(antinodes, antinode)
					left := antinode
					right := primary

					for {
						if !withResonantHarmonics || antinode.x > maxX || antinode.x < 0 || antinode.y > maxY || antinode.y < 0 {
							break
						}
						// Keep running deltas until exit condition
						antinode = getAntinodeLocation(left, right)
						antinodes = append(antinodes, antinode)
						right = left
						left = antinode
					}

					// Include this antenna, maybe?
					if withResonantHarmonics {
						antinodes = append(antinodes, primary)
					}
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
