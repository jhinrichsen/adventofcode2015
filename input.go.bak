package adventofcode2015

import (
	"bufio"
	"io"
	"os"
)

// ReadLines reads lines from any reader
func ReadLines[R io.Reader](r R) ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// linesFromFilename reads lines from a file
func linesFromFilename(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()
	return ReadLines(f)
}

// linesFromReader reads lines from a reader
func linesFromReader(r io.Reader) ([]string, error) {
	lines, err := ReadLines(r)
	if err != nil {
		return nil, err
	}
	return lines, nil
}
