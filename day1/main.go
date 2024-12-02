package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := readInput("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	firstList := []string{}
	secondList := []string{}
	for _, line := range lines {
		firstList = append(firstList, strings.Split(line, "   ")[0])
		secondList = append(secondList, strings.Split(line, "   ")[1])
	}

	firstIntList := listToFloat(firstList)
	secondIntList := listToFloat(secondList)

	sort.Float64s(firstIntList)
	sort.Float64s(secondIntList)

	// Part 1
	absoluteList := []float64{}
	for index := range firstIntList {
		absolute := math.Abs(firstIntList[index] - secondIntList[index])
		absoluteList = append(absoluteList, absolute)
	}

	absoluteSum := 0.0
	for _, i := range absoluteList {
		absoluteSum += i
	}
	fmt.Println(int(absoluteSum))

	// Part 2
	similarityScoreList := []float64{}
	for index := range firstIntList {
		count := repeatsInList(secondIntList, firstIntList[index])
		fmt.Println("count:", count)
		similarityScore := count * firstIntList[index]
		fmt.Println("similarityScore:", similarityScore)
		similarityScoreList = append(similarityScoreList, similarityScore)
	}

	similarityScoreSum := 0.0
	for _, i := range similarityScoreList {
		similarityScoreSum += i
	}
	fmt.Println(int(similarityScoreSum))

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

func listToFloat(stringList []string) (floatList []float64) {
	for _, i := range stringList {
		i, err := strconv.ParseFloat(i, 64)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return []float64{}
		}
		floatList = append(floatList, i)
	}
	return floatList
}

func repeatsInList(list []float64, repeater float64) (repeats float64) {
	for _, i := range list {
		if i == repeater {
			repeats++
		}
	}
	fmt.Println(repeater, repeats)
	return repeats
}
