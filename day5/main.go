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
	lines, err := readInput("example.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	ordering, updates := getOrderingAndUpdates(lines)

	xyList := getXYList(ordering)
	listOfIntUpdates := getListOfIntUpdates(updates)

	correctOrders := [][]int{}
	for _, update := range listOfIntUpdates {
		isValid := true
		for i := 0; i < len(update)-1; i++ {
			fmt.Println("Checking:", update)
			if !rightOrder(update[i], update[i+1], xyList) {
				fmt.Println("Not in right order:", update[i], update[i+1])
				isValid = false
				break
			}
		}
		if isValid {
			correctOrders = append(correctOrders, update)
		}
	}
	fmt.Println("Correct orders:", correctOrders)

	sumOfMiddleElements := 0
	// get middle element
	for _, order := range correctOrders {
		index := len(order) / 2
		fmt.Println("Middle element:", order[index])
		sumOfMiddleElements += order[index]
	}
	fmt.Println("Sum of middle elements:", sumOfMiddleElements)

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
