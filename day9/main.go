package main

import (
	"fmt"
	"os"
	"strconv"
)

// var input string = "2333133121414131402"

func main() {

	input, _ := os.ReadFile("input.txt")

	lengthFiles, lengthFreeSpaces := getMap(string(input))

	// Part 1
	disk := getDisk(lengthFiles, lengthFreeSpaces)
	disk = sortDisk(disk)
	sortedDisk := []int{}
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			continue
		}
		sortedDisk = append(sortedDisk, disk[i])
	}
	fmt.Println("Disk: ", sortedDisk)
	checksum := 0
	for x := 0; x < len(sortedDisk); x++ {
		checksum += x * sortedDisk[x]
	}
	fmt.Println("Checksum: ", checksum)

	// Part 2
	disk2 := getDisk(lengthFiles, lengthFreeSpaces)
	disk2 = sortDiskV2(disk2)

	fmt.Println("Disk2: ", disk2)
	checksum2 := 0
	for x := 0; x < len(disk2); x++ {
		if disk2[x] == -1 {
			continue
		}
		checksum2 += x * disk2[x]
		fmt.Println(x, "*", disk2[x], "Checksum: ", checksum2)
	}
	fmt.Println("Checksum: ", checksum2)
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
	// Find max file ID
	maxID := -1
	for _, v := range disk {
		if v > maxID {
			maxID = v
		}
	}

	// Process each file ID from highest to lowest
	for fileID := maxID; fileID >= 0; fileID-- {
		// Find file size and position
		fileSize := 0
		fileStart := -1
		for i := 0; i < len(disk); i++ {
			if disk[i] == fileID {
				if fileStart == -1 {
					fileStart = i
				}
				fileSize++
			}
		}

		if fileSize == 0 {
			continue
		}

		// Find leftmost valid free space
		bestFreeStart := -1
		currentFreeStart := -1
		currentFreeSize := 0

		for i := 0; i < fileStart; i++ {
			if disk[i] == -1 {
				if currentFreeStart == -1 {
					currentFreeStart = i
				}
				currentFreeSize++

				if currentFreeSize >= fileSize {
					bestFreeStart = currentFreeStart
					break
				}
			} else {
				currentFreeStart = -1
				currentFreeSize = 0
			}
		}

		// Move file if valid space found
		if bestFreeStart != -1 {
			// Copy file to new position
			for i := 0; i < fileSize; i++ {
				disk[bestFreeStart+i] = fileID
			}
			// Clear old position
			for i := fileStart; i < fileStart+fileSize; i++ {
				disk[i] = -1
			}
			fmt.Println("Disk: ", disk)
		}
	}
	return disk
}
