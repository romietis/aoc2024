package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [8][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // down-right
	{1, -1},  // down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}

var directions2 = [4][2]int{
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}

// Done with LLM assistance
func main() {
	grid, err := readFileToRuneGrid("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grid2 := [][]rune{
		{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
		{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
		{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
		{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
		{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
		{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
		{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
		{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
		{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
		{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
	}
	word := "XMAS"

	fmt.Println("Total occurrences of", word, ":", countOccurrences(grid, word))
	fmt.Println("Total occurrences of", word, ":", countOccurrences(grid2, word))

	fmt.Println("Total X-MAS patterns:", findXMAS(grid))
	fmt.Println("Total X-MAS patterns:", findXMAS(grid2))
}

func readFileToRuneGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func searchWord(grid [][]rune, word string, x, y int) int {
	count := 0
	wordLen := len(word)
	rows := len(grid)
	cols := len(grid[0])

	for _, dir := range directions {
		k, newX, newY := 0, x, y
		for k < wordLen {
			if !isValid(newX, newY, rows, cols) || grid[newX][newY] != rune(word[k]) {
				break
			}
			newX += dir[0]
			newY += dir[1]
			k++
		}
		if k == wordLen {
			count++
		}
	}
	return count
}

func countOccurrences(grid [][]rune, word string) int {
	totalCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == rune(word[0]) {
				checkXPattern(grid, i, j)
				totalCount += searchWord(grid, word, i, j)
			}
		}
	}
	return totalCount
}

// Part 2
var (
	// Diagonal directions for X pattern
	diagonal1 = [][]int{{-1, -1}, {1, 1}} // top-left to bottom-right
	diagonal2 = [][]int{{-1, 1}, {1, -1}} // top-right to bottom-left
)

func findXMAS(grid [][]rune) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			if grid[i][j] == 'A' {
				if checkXPattern(grid, i, j) {
					count++
				}
			}
		}
	}
	return count
}

func checkXPattern(grid [][]rune, x, y int) bool {
	rows, cols := len(grid), len(grid[0])

	// Check first diagonal (top-left to bottom-right)
	validDiag1 := false
	if isValid(x-1, y-1, rows, cols) && isValid(x+1, y+1, rows, cols) {
		if (grid[x-1][y-1] == 'M' && grid[x+1][y+1] == 'S') ||
			(grid[x-1][y-1] == 'S' && grid[x+1][y+1] == 'M') {
			validDiag1 = true
		}
	}

	// Check second diagonal (top-right to bottom-left)
	validDiag2 := false
	if isValid(x-1, y+1, rows, cols) && isValid(x+1, y-1, rows, cols) {
		if (grid[x-1][y+1] == 'M' && grid[x+1][y-1] == 'S') ||
			(grid[x-1][y+1] == 'S' && grid[x+1][y-1] == 'M') {
			validDiag2 = true
		}
	}

	return validDiag1 && validDiag2
}
