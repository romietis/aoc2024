package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputs, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	intLists := [][]int{}
	for _, input := range inputs {
		intList := listToInts(input)
		intLists = append(intLists, intList)
	}

	safe := 0
	for _, list := range intLists {
		if isValidSequence(list) || canBeMadeValid(list) {
			safe++
			fmt.Printf("List %v is safe\n", list)
		} else {
			fmt.Printf("List %v is unsafe\n", list)
		}
	}
	fmt.Println("Total safe lists:", safe)
}

func canBeMadeValid(list []int) bool {
	// Try removing each number once
	for i := 0; i < len(list); i++ {
		newList := make([]int, 0, len(list)-1)
		// Append all before i
		newList = append(newList, list[:i]...)
		// Append all after i
		newList = append(newList, list[i+1:]...)

		if isValidSequence(newList) {
			return true
		}
	}
	return false
}

func isValidSequence(list []int) bool {
	if !increasingOrDecreasing(list) {
		return false
	}

	for i := 0; i < len(list)-1; i++ {
		abs := math.Abs(float64(list[i]) - float64(list[i+1]))
		if !absValueInRange(int(abs), 1, 3) {
			return false
		}
	}
	return true
}

func readInput(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.Split(line, " "))
	}

	return lines, scanner.Err()
}

func increasingOrDecreasing(intList []int) bool {
	if slices.IsSorted(intList) {
		return true
	}
	reversedList := slices.Clone(intList)
	slices.Reverse(reversedList)
	if slices.IsSorted(reversedList) {
		return true
	}
	return false
}

func listToInts(stringList []string) (intList []int) {
	for _, i := range stringList {
		i, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return []int{}
		}
		intList = append(intList, i)
	}
	return intList
}

func absValueInRange(value int, min int, max int) bool {
	if value >= min && value <= max {
		return true
	}
	return false
}
