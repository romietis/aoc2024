package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	matches := findMul(string(file))
	multiplications := []int{}
	sum := 0

	//Part 1
	for _, i := range matches {
		multiplications = append(multiplications, mul(i))
	}

	for _, i := range multiplications {
		sum += i
	}
	fmt.Println("Sum of multiplications:", sum)

	//Part 2
	sum = 0
	multiplications = []int{}
	for i := 0; i < len(matches); i++ {

		if matches[i] == "don't()" {
			for matches[i] != "do()" {
				i++
			}
		}
		multiplications = append(multiplications, mul(matches[i]))
	}

	for _, i := range multiplications {
		sum += i
	}
	fmt.Println("Sum of multiplications with do() and don't():", sum)

}

func mul(input string) int {
	if input == "do()" || input == "don't()" {
		return 0
	}
	re := regexp.MustCompile(`(\d+),(\d+)`)
	matches := re.FindAllString(input, -1)
	matches = strings.Split(matches[0], ",")
	matchOne, _ := strconv.Atoi(matches[0])
	matchTwo, _ := strconv.Atoi(matches[1])
	multiplication := matchOne * matchTwo
	return multiplication
}

func findMul(input string) []string {
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(string(input), -1)
	return matches
}
