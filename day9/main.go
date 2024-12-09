package main

import (
	"fmt"
	"strconv"
)

var input string = "2333133121414131402"

func main() {

	// input, _ := os.ReadFile("input.txt")

	lengthFiles, lengthFreeSpaces := getMap(string(input))

	disk := getDisk(lengthFiles, lengthFreeSpaces)

	disk = sortDisk(disk)

	sortedDisk := []int{}
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		sortedDisk = append(sortedDisk, disk[i])
	}
	checksum := 0
	for x := 0; x < len(sortedDisk); x++ {
		checksum += x * sortedDisk[x]
	}
	fmt.Println("Checksum: ", checksum)
}

func getMap(input string) ([]int, []int) {
	lengthFreeSpaces := []int{}
	lengthFiles := []int{}

	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			times, _ := strconv.Atoi(string(input[i]))
			lengthFiles = append(lengthFiles, times)
		}
		if i%2 != 0 {
			times, _ := strconv.Atoi(string(input[i]))
			lengthFreeSpaces = append(lengthFreeSpaces, times)
		}
	}
	return lengthFiles, lengthFreeSpaces
}

func getDisk(lengthFiles []int, lengthFreeSpaces []int) []int {
	disk := []int{}
	for j := 0; j < len(lengthFiles); j++ {
		for i := 0; i < lengthFiles[j]; i++ {
			disk = append(disk, j)
		}
		if j == len(lengthFiles)-1 {
			break
		}
		for i := 0; i < lengthFreeSpaces[j]; i++ {
			disk = append(disk, -1)
		}
	}
	return disk
}

func sortDisk(disk []int) []int {
	isValid := true
	latestDot := 0
	latestNumber := len(disk) - 1

	for i := latestNumber; i > 0; i-- {
		if isValid {
			if disk[i] != -1 {
				// Find first available -1
				for j := latestDot; j < len(disk); j++ {
					if j > i {
						isValid = false
						break
					}
					if disk[j] == -1 {
						// Move digit to -1 position
						disk[j] = disk[i]
						disk[i] = -1
						latestDot = j
						latestNumber = i
						break
					}
				}
				fmt.Println("Disk: ", disk)
			}
		}
	}
	return disk
}

func sortDiskV2(disk []int) []int {
	isValid := true
	latestDot := 0
	latestNumber := len(disk) - 1

	for i := latestNumber; i > 0; i-- {
		if isValid {
			if disk[i] != -1 {
				// Find first available -1
				for j := latestDot; j < len(disk); j++ {
					if j > i {
						isValid = false
						break
					}
					if disk[j] == -1 {
						// Move digit to -1 position
						disk[j] = disk[i]
						disk[i] = -1
						latestDot = j
						latestNumber = i
						break
					}
				}
				fmt.Println("Disk: ", disk)
			}
		}
	}
	return disk
}
