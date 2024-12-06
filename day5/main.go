package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines, err := readInput("example.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	pageOrdering := []string{}
	updates := []string{}

	split := findEmptyLine(lines)
	// append everything before the empty line to pageOrdering
	pageOrdering = lines[:split]
	// append everything after the empty line to updates
	updates = lines[split+1:]

	fmt.Println("Page ordering:", pageOrdering)
	fmt.Println("Updates:", updates)
}

func readInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func findEmptyLine(lines []string) int {
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			return i
		}
	}
	return -1
}
