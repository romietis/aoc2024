package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, err := readFileToRuneGrid("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// directions := [4]rune{'>', 'v', '<', '^'}

	for isDirection(grid) {
		x, y := findStartingPoint(grid)
		direction := grid[x][y]
		newX, newY := move(x, y, direction)

		// Mark current position before moving
		grid[x][y] = 'X'

		// Check if new position is valid before updating
		if isValid(newX, newY, len(grid), len(grid[0])) {
			// Only update if not at border
			if isHash(newX, newY, grid) {
				direction = directionRules(direction)
				grid[x][y] = direction
			} else if !(newX == 0 || newY == 0 || newX == len(grid)-1 || newY == len(grid[0])-1) {
				grid[newX][newY] = direction
			}
		}
	}

	printGrid(grid)
	fmt.Println("Total occurrences:", countOccurrences(grid, 'X'))
}

func isDirection(grid [][]rune) bool {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '^' || grid[row][col] == 'v' || grid[row][col] == '<' || grid[row][col] == '>' {
				return true
			}
		}
	}
	return false

}

func printGrid(grid [][]rune) {
	for row := range grid {
		for col := range grid[row] {
			fmt.Print(string(grid[row][col]), " ")
		}
		fmt.Println()
	}
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func isHash(x, y int, grid [][]rune) bool {
	return grid[x][y] == '#'
}

func findStartingPoint(grid [][]rune) (int, int) {
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == '^' || grid[row][col] == 'v' || grid[row][col] == '<' || grid[row][col] == '>' {
				return row, col
			}
		}
	}
	return -1, -1
}

func move(x, y int, direction rune) (int, int) {
	if direction == '^' {
		return x - 1, y
	}
	if direction == 'v' {
		return x + 1, y
	}
	if direction == '<' {
		return x, y - 1
	}
	if direction == '>' {
		return x, y + 1
	}
	return x, y
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

func directionRules(currentDirection rune) rune {
	switch currentDirection {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	return currentDirection
}

func countOccurrences(grid [][]rune, x rune) int {
	totalCount := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == x {
				totalCount++
			}
		}
	}
	return totalCount
}
