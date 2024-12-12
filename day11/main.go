package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var input string = "8793800 1629 65 5 960 0 138983 85629"
	// var input string = "125 17"
	inputList := strings.Split(input, " ")

	// Part 1
	blinksNeeded := 75
	blinks := 0
	// for blinks < blinksNeeded {
	// 	newStones := []string{}
	// 	for _, stone := range inputList {
	// 		switch {
	// 		case stone == "0":
	// 			newStones = append(newStones, "1")
	// 		case len(stone)%2 == 0:
	// 			firstHalf := stone[:len(stone)/2]
	// 			firstHalf = strings.TrimLeft(firstHalf, "0")
	// 			secondHalf := stone[len(stone)/2:]
	// 			match, _ := regexp.MatchString("^0+$", secondHalf)
	// 			if match {
	// 				secondHalf = "0"
	// 			} else {
	// 				secondHalf = strings.TrimLeft(secondHalf, "0")
	// 			}
	// 			newStones = append(newStones, firstHalf)
	// 			newStones = append(newStones, secondHalf)
	// 		default:
	// 			stoneInt, _ := strconv.Atoi(stone)
	// 			stoneInt = stoneInt * 2024
	// 			stone = strconv.Itoa(stoneInt)
	// 			newStones = append(newStones, stone)
	// 		}
	// 	}
	// 	inputList = newStones
	// 	blinks++
	// 	// fmt.Println("After Blink", blinks, newStones)
	// 	if blinks == blinksNeeded {
	// 		fmt.Println("Total stones", len(newStones))
	// 	}
	// }

	// Part 2
	inputList = strings.Split(input, " ")
	cache := make(map[string]int)
	for _, stone := range inputList {
		cache[stone]++
	}

	blinks = 0
	for blinks < blinksNeeded {
		newCache := make(map[string]int)
		for key, value := range cache {
			switch {
			case key == "0":
				newCache["1"] += value
			case len(key)%2 == 0:
				firstHalf := key[:len(key)/2]
				firstHalf = strings.TrimLeft(firstHalf, "0")
				secondHalf := key[len(key)/2:]
				match, _ := regexp.MatchString("^0+$", secondHalf)
				if match {
					secondHalf = "0"
				} else {
					secondHalf = strings.TrimLeft(secondHalf, "0")
				}
				newCache[firstHalf] += value
				newCache[secondHalf] += value
			default:
				keyInt, _ := strconv.Atoi(key)
				keyInt = keyInt * 2024
				newCache[strconv.Itoa(keyInt)] += value
			}
		}
		cache = newCache
		blinks++
		if blinks == blinksNeeded {
			count := 0
			for _, value := range cache {
				count += value
			}
			fmt.Println("Total stones", count)
		}
	}
}
