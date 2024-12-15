package main

import (
	"fmt"
	"helper/parser"
	"strconv"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")
	fmt.Println(solve(demoInput[0]))
	fmt.Println(solve(fullInput[0]))
	//fmt.Println(solve(demoInput, true))
	//fmt.Println(solve(fullInput, true))
}

func solve(input string) int64 {
	disk := generateDisk(input)
	newDisk := compressDisk(disk)
	return getChecksum(newDisk)
}

func getChecksum(disk []int) int64 {
	var total int64
	total = 0

	for idx, value := range disk {
		total += int64(idx) * int64(value)
	}

	return total
}

func compressDisk(disk []int) []int {
	// Use 2 pointer approach to move left and right pointers and create compressed disk.
	// In-memory would be better, but I ain't getting paid for this
	var newDisk []int
	left := 0
	right := len(disk) - 1

	for {
		if left > right {
			break
		}
		if disk[right] == -1 {
			right--
			continue
		}

		if disk[left] != -1 {
			newDisk = append(newDisk, disk[left])
			left++
			continue
		}

		if disk[right] != -1 && disk[left] == -1 {
			newDisk = append(newDisk, disk[right])
			right--
			left++ // We essentially write over the empty space but still need to move the pointer
			continue
		}
	}

	return newDisk
}

func generateDisk(input string) []int {
	// We use -1 for '.' I guess.
	// Runes seem painful and bytes only go up to 255, which would work with the demo but not with
	// full input which goes up to 10k integers (20k line length / 2)
	var disk []int
	counter := 0

	for i, char := range input {
		fileNumber, _ := strconv.Atoi(string(char))

		if i%2 == 0 {
			// Append the numeric value as many times as needed
			for v := 0; v < fileNumber; v++ {
				disk = append(disk, counter)
			}
			counter++
		} else {
			// Use -1 to represent '.'
			for v := 0; v < fileNumber; v++ {
				disk = append(disk, -1)
			}
		}
	}

	return disk
}

func printDisk(disk []int) {
	for _, value := range disk {
		if value == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(value)
		}
	}
	fmt.Println()
}
