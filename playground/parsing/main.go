package main

import (
	"fmt"
	"helper/parser"
	"strings"
)

func main() {
	tfp := parser.TextFileParser{}
	content, err := tfp.ParseLinesFromPathAsString("sample-input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	sentence := strings.Join(content, " ")
	fmt.Println(sentence)
}
