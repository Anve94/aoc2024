package parser_test

import (
	"helper/parser"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func createTempFile(t *testing.T, content string) string {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "example")

	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %s", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %s", err)
	}

	return tmpFile.Name()
}

func TestParseLinesFromPathAsString(t *testing.T) {
	content := "line1\nline2\nline3"
	path := createTempFile(t, content)
	defer os.Remove(path)

	parser := parser.TextFileParser{}

	expectedLines := []string{"line1", "line2", "line3"}
	lines, err := parser.ParseLinesFromPathAsString(path)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("expected %v, got %v", expectedLines, lines)
	}
}

func TestParseLinesFromPathAsBytes(t *testing.T) {
	content := "line1\nline2\nline3"
	path := createTempFile(t, content)
	defer os.Remove(path)

	parser := parser.TextFileParser{}

	expectedLines := [][]byte{
		[]byte("line1"),
		[]byte("line2"),
		[]byte("line3"),
	}
	lines, err := parser.ParseLinesFromPathAsBytes(path)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("expected %v, got %v", expectedLines, lines)
	}
}

// Additional test for error handling when the file does not exist
func TestParseLinesFromNonExistentFile(t *testing.T) {
	parser := parser.TextFileParser{}

	_, err := parser.ParseLinesFromPathAsString("non_existent_file.txt")
	if err == nil {
		t.Error("expected an error for non-existent file")
	}

	_, err = parser.ParseLinesFromPathAsBytes("non_existent_file.txt")
	if err == nil {
		t.Error("expected an error for non-existent file")
	}
}
