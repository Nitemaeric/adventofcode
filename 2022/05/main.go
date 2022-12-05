package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var stacks3000 [][]rune
var stacks3001 [][]rune

func main() {
	file, fileError := os.Open("data.txt")

	if fileError != nil {
		fmt.Println(fileError)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if len(stacks3000) == 0 {
			stacks3000 = SetupStacks(line)
			stacks3001 = SetupStacks(line)
		}

		if !strings.HasPrefix(line, "move") {
			newRunes := ParseSetupLine(line)

			stacks3000 = Prepend(stacks3000, newRunes)
			stacks3001 = Prepend(stacks3001, newRunes)
		} else {
			from, to, times := ParseMoveLine(line)

			stacks3000 = Move3000(stacks3000, from, to, times)
			stacks3001 = Move3001(stacks3001, from, to, times)
		}
	}

	PrintStacks(stacks3000)
	PrintStacks(stacks3001)

	fmt.Printf("Step 1: %s\n", TopOfStacks(stacks3000))
	fmt.Printf("Step 2: %s\n", TopOfStacks(stacks3001))
}

func SetupStacks(line string) [][]rune {
	stackCount, _ := ParseStackCount(line)

	return make([][]rune, stackCount)
}

func Prepend(stacks [][]rune, newRunes []rune) [][]rune {
	for index, char := range newRunes {
		if unicode.IsLetter(char) {
			stacks[index] = append([]rune{char}, stacks[index]...)
		}
	}

	return stacks
}

func ParseSetupLine(line string) []rune {
	stackCount, lineLength := ParseStackCount(line)

	runes := make([]rune, stackCount)

	for index := 0; index < lineLength; index += 4 {
		maxLength := index + 4

		if maxLength >= lineLength {
			maxLength = lineLength
		}

		char := rune(line[index:maxLength][1])

		if unicode.IsLetter(char) {
			runes[index/4] = char
		}
	}

	return runes
}

func Move3000(stacks [][]rune, from int, to int, times int) [][]rune {
	fromStack := stacks[from-1]
	toStack := stacks[to-1]

	for time := 0; time < times; time += 1 {
		var toMove rune

		toMove, fromStack = fromStack[len(fromStack)-1], fromStack[:len(fromStack)-1]

		toStack = append(toStack, toMove)

		stacks[from-1] = fromStack
		stacks[to-1] = toStack
	}

	return stacks
}

func Move3001(stacks [][]rune, from int, to int, times int) [][]rune {
	fromStack := stacks[from-1]
	toStack := stacks[to-1]

	var toMove []rune

	toMove, fromStack = fromStack[len(fromStack)-times:], fromStack[:len(fromStack)-times]

	toStack = append(toStack, toMove...)

	stacks[from-1] = fromStack
	stacks[to-1] = toStack

	return stacks
}

func ParseMoveLine(line string) (int, int, int) {
	pattern := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")

	matches := pattern.FindStringSubmatch(line)

	from, _ := strconv.ParseInt(matches[2], 10, 36)
	to, _ := strconv.ParseInt(matches[3], 10, 36)
	times, _ := strconv.ParseInt(matches[1], 10, 36)

	return int(from), int(to), int(times)
}

func ParseStackCount(line string) (int, int) {
	lineLength := len(line)
	stackCount := int(math.Ceil(float64(lineLength) / 4))

	return stackCount, lineLength
}

func PrintStacks(stacks [][]rune) {
	tallestStackSize := GetTallestStackSize(stacks)

	for height := tallestStackSize; height >= 0; height -= 1 {
		line := ""

		for _, stack := range stacks {
			if height < len(stack) {
				line += "[" + string(stack[height]) + "] "
			} else {
				line += "    "
			}
		}

		fmt.Printf("%s\n", line)
	}
}

func GetTallestStackSize(stacks [][]rune) int {
	tallestStackSize := 0

	for _, stack := range stacks {
		if len(stack) > tallestStackSize {
			tallestStackSize = len(stack)
		}
	}

	return tallestStackSize
}

func TopOfStacks(stacks [][]rune) string {
	lastChars := []rune{}

	for _, stack := range stacks {
		lastCharIndex := len(stack) - 1

		lastChars = append(lastChars, stack[lastCharIndex])
	}

	return string(lastChars)
}
