package main

import (
	"fmt"
	"helper/parser"
	"strconv"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	demoInput, _ := tfp.ParseLinesFromPathAsString("demo.txt")
	fullInput, _ := tfp.ParseLinesFromPathAsString("input.txt")

	fmt.Println("Demo result:", part1(demoInput))
	fmt.Println("Full result:", part1(fullInput))
	fmt.Println("Demo 2 result:", part2(demoInput))
	fmt.Println("Full 2 result:", part2(fullInput))
}

func part1(input []string) int {
	orders, pages := processInputs(input)
	hashMap := createHashMap(orders)
	total := 0

	for _, page := range pages {
		_, isValid := isValidPage(page, hashMap)
		if isValid {
			found := getMiddleValue(page)
			total += found
		}
	}
	return total
}

func part2(input []string) int {
	orders, pages := processInputs(input)
	hashMap := createHashMap(orders)
	total := 0
	var incorrectPages []string

	for _, page := range pages {
		_, isValid := isValidPage(page, hashMap)
		if !isValid {
			incorrectPages = append(incorrectPages, page)
		}
	}

	for _, incorrectPage := range incorrectPages {
		validPage := correctPage(incorrectPage, hashMap)
		total += getMiddleValue(validPage)
	}

	return total
}

func correctPage(page string, hashMap map[int][]int) string {
	values := strings.Split(page, ",")

	for {
		wrongIndex, ok := isValidPage(strings.Join(values, ","), hashMap)
		if ok {
			return strings.Join(values, ",")
		}
		values[wrongIndex], values[wrongIndex-1] = values[wrongIndex-1], values[wrongIndex]
	}
}

func getMiddleValue(page string) int {
	var values []int
	pageStrings := strings.Split(page, ",")
	for _, pageStr := range pageStrings {
		v, _ := strconv.Atoi(pageStr)
		values = append(values, v)
	}
	return values[len(values)/2]
}

func isValidPage(page string, orderMap map[int][]int) (int, bool) {
	pageValues := strings.Split(page, ",")
	var pages []int
	for _, str := range pageValues {
		i, _ := strconv.Atoi(str)
		pages = append(pages, i)
	}

	for orderIdx, value := range pages {
		// Hashmap data e.g. {pageIdx: [1, 2, 3]}
		mustComeLater, ok := orderMap[value]
		// Check for existence, if there are no pages that need to come later it can be anywhere
		if !ok {
			continue
		}

		for _, mustComeLaterValue := range mustComeLater {
			pageIndexes := FindAllIndexes(pages, mustComeLaterValue)
			for _, i := range pageIndexes {
				if i < orderIdx {
					return orderIdx, false
				}
			}
		}
	}
	return -1, true
}

func processInputs(input []string) ([]string, []string) {
	var orderingRules, pages []string
	isOrderingRule := true

	for _, line := range input {
		if line == "" {
			isOrderingRule = false
			continue
		}

		if isOrderingRule {
			orderingRules = append(orderingRules, line)
		} else {
			pages = append(pages, line)
		}
	}

	return orderingRules, pages
}

func createHashMap(orderingRules []string) map[int][]int {
	var hashMap = make(map[int][]int)
	for _, rule := range orderingRules {
		values := strings.Split(rule, "|")
		i, _ := strconv.Atoi(values[0])
		v, _ := strconv.Atoi(values[1])
		hashMap[i] = append(hashMap[i], v)
	}

	return hashMap
}

func FindAllIndexes[T comparable](slice []T, target T) []int {
	var indexes []int
	for i, v := range slice {
		if v == target {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
