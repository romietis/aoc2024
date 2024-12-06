package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type XY struct {
	before int
	after  int
}

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Parse input
	ordering, updates := getOrderingAndUpdates(lines)
	xyList := getXYList(ordering)
	listOfIntUpdates := getListOfIntUpdates(updates)
	correctOrders, incorrectOrders := getCorrectAndIncorrectOrders(listOfIntUpdates, xyList)

	// Part 1
	sumOfMiddleElements := getSumOfMiddleElements(correctOrders)
	fmt.Println("Sum of middle elements:", sumOfMiddleElements)

	// Part 2
	// Stupid way to correct the order, but it works
	finalCorrectedOrders := [][]int{}
	for i := 0; i < len(incorrectOrders); i++ {
		incorrectOrders = correctTheOrder(incorrectOrders, xyList)
		if i == len(incorrectOrders)-1 {
			finalCorrectedOrders = incorrectOrders
		}
	}
	sumOfMiddleElements = getSumOfMiddleElements(finalCorrectedOrders)
	fmt.Println("Sum of middle elements after correction:", sumOfMiddleElements)

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

func getOrderingAndUpdates(lines []string) ([]string, []string) {
	split := findEmptyLine(lines)
	return lines[:split], lines[split+1:]
}

func getXYList(ordering []string) []XY {
	xyList := []XY{}
	for _, xy := range ordering {
		xyTuple := strings.Split(xy, "|")
		x, _ := strconv.Atoi(xyTuple[0])
		y, _ := strconv.Atoi(xyTuple[1])
		xyList = append(xyList, XY{x, y})
	}
	return xyList
}

func getListOfIntUpdates(updates []string) [][]int {
	listOfIntUpdates := [][]int{}
	for _, update := range updates {
		updateList := []int{}
		for _, page := range strings.Split(update, ",") {
			pageInt, _ := strconv.Atoi(page)
			updateList = append(updateList, pageInt)
		}
		listOfIntUpdates = append(listOfIntUpdates, updateList)
	}
	return listOfIntUpdates
}

func rightOrder(x, y int, xys []XY) bool {
	for index, xy := range xys {
		if x != xy.after && y != xy.before {
			index++
		}
		if x == xy.before && y == xy.after {
			return true
		}
	}
	return false
}

func getCorrectAndIncorrectOrders(listOfIntUpdates [][]int, xyList []XY) ([][]int, [][]int) {
	correctOrders := [][]int{}
	incorrectOrder := [][]int{}
	for _, update := range listOfIntUpdates {
		isValid := true
		for i := 0; i < len(update)-1; i++ {
			if !rightOrder(update[i], update[i+1], xyList) {
				isValid = false
				break
			}
		}
		if isValid {
			correctOrders = append(correctOrders, update)
		}
		if !isValid {
			incorrectOrder = append(incorrectOrder, update)
		}
	}
	return correctOrders, incorrectOrder
}

func getSumOfMiddleElements(correctOrders [][]int) int {
	sumOfMiddleElements := 0
	for _, order := range correctOrders {
		index := len(order) / 2
		sumOfMiddleElements += order[index]
	}
	return sumOfMiddleElements
}

func correctTheOrder(incorrectOrders [][]int, xyList []XY) [][]int {
	correctedOrders := [][]int{}
	for _, order := range incorrectOrders {
		for i := 0; i < len(order)-1; i++ {
			if !rightOrder(order[i], order[i+1], xyList) {
				tempValue := order[i]
				order[i] = order[i+1]
				order[i+1] = tempValue
			}
		}
		correctedOrders = append(correctedOrders, order)
	}
	return correctedOrders
}
