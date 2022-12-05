package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, fileError := os.Open("data.txt")

	if fileError != nil {
		fmt.Println(fileError)
	}

	fileScanner := bufio.NewScanner(file)
	allLines := []string{}

	partOneTotalPriority := 0
	partTwoTotalPriority := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		firstHalf, secondHalf := parseLine(line)

		// duplicate := FindDuplicate(firstHalf, secondHalf)

		halves := []string{firstHalf, secondHalf}
		duplicate := FindCommonRune(halves)

		partOneTotalPriority += CalculatePriority(duplicate)

		// Add all lines for Part 2

		allLines = append(allLines, line)
	}

	for index := 0; index < len(allLines); index += 3 {
		firstGroupInput := allLines[index]
		secondGroupInput := allLines[index+1]
		thirdGroupInput := allLines[index+2]

		// duplicate := FindGroupDuplicate(firstGroupInput, secondGroupInput, thirdGroupInput)

		groupInputs := []string{firstGroupInput, secondGroupInput, thirdGroupInput}
		duplicate := FindCommonRune(groupInputs)

		partTwoTotalPriority += CalculatePriority(duplicate)
	}

	partOneString := fmt.Sprintf("Part 1: %d", partOneTotalPriority)
	fmt.Println(partOneString)

	partTwoString := fmt.Sprintf("Part 2: %d", partTwoTotalPriority)
	fmt.Println(partTwoString)

	file.Close()
}

func CalculatePriority(input rune) int {
	ascii := int(input) - 65

	if ascii > 26 {
		return ascii - 31
	} else {
		return ascii + 27
	}
}

// Initial Part 1 - O(n^2) [n: input size] as this tries to find a duplicate in the secondInput [O(n)] n times [firstInput size]
func FindDuplicate(firstInput string, secondInput string) rune {
	for _, char := range firstInput {
		index := strings.IndexRune(secondInput, char)

		if index != -1 {
			return char
		}
	}

	return -1
}

// Initial Part 2 - O(n) [n: input size] as this method only allows for 3 inputs
func FindGroupDuplicate(firstInput string, secondInput string, thirdInput string) rune {
	counts := map[rune]int{}

	for _, char := range firstInput {
		counts[char] = 1
	}

	for _, char := range secondInput {
		if counts[char] == 1 {
			counts[char] = 2
		}
	}

	for _, char := range thirdInput {
		if counts[char] == 2 {
			counts[char] = 3
		}
	}

	for char, count := range counts {
		if count == 3 {
			return char
		}
	}

	return -1
}

// Let's optimise a little - O(kn) [k: inputs length, n: input size]
func FindCommonRune(inputs []string) rune {
	counts := map[rune]int{}
	length := len(inputs)

	for _, input := range inputs {
		groupPresence := map[rune]bool{}

		for _, char := range input {
			if !groupPresence[char] {
				counts[char] += 1
				groupPresence[char] = true
			}

			if counts[char] == length {
				return char
			}
		}
	}

	return -1
}

func parseLine(input string) (string, string) {
	midPoint := len(input) / 2

	return input[:midPoint], input[midPoint:]
}
