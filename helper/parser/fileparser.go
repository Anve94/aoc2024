package parser

import (
	"bufio"
	"os"
)

type FileParser interface {
	ParseLinesFromPathAsString(path string) ([]string, error)
	ParseLinesFromPathAsBytes(path string) ([][]byte, error)
}

type TextFileParser struct{}

func (tfp TextFileParser) ParseLinesFromPathAsString(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func (tfp TextFileParser) ParseLinesFromPathAsBytes(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]byte
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Bytes()
		lines = append(lines, []byte(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
