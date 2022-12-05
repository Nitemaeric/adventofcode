package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Section struct {
	start int
	end   int
}

func main() {
	file, fileError := os.Open("data.txt")

	if fileError != nil {
		fmt.Println(fileError)
	}

	fileScanner := bufio.NewScanner(file)

	countCovering := 0
	countOverlapping := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()

		firstSection, secondSection := ParseSections(line)

		isCovering := CheckCovering(firstSection, secondSection)
		isOverlapping := CheckOverlapping(firstSection, secondSection)

		if isCovering {
			countCovering += 1
		}

		if isOverlapping {
			countOverlapping += 1
		}
	}

	coveringString := fmt.Sprintf("There are %d ranges fully containing the other.", countCovering)
	fmt.Println(coveringString)

	overlappingString := fmt.Sprintf("There are %d ranges overlapping the other.", countOverlapping)
	fmt.Println(overlappingString)
}

func ParseSections(line string) (Section, Section) {
	sectionStrings := strings.Split(line, ",")

	firstSectionValues := strings.Split(sectionStrings[0], "-")
	secondSectionValues := strings.Split(sectionStrings[1], "-")

	firstSectionStart, _ := strconv.ParseInt(firstSectionValues[0], 10, 36)
	firstSectionEnd, _ := strconv.ParseInt(firstSectionValues[1], 10, 36)

	secondSectionStart, _ := strconv.ParseInt(secondSectionValues[0], 10, 36)
	secondSectionEnd, _ := strconv.ParseInt(secondSectionValues[1], 10, 36)

	firstSection := Section{start: int(firstSectionStart), end: int(firstSectionEnd)}
	secondSection := Section{start: int(secondSectionStart), end: int(secondSectionEnd)}

	return firstSection, secondSection
}

func CheckCovering(firstSection Section, secondSection Section) bool {
	firstSectionCovering := firstSection.start <= secondSection.start && firstSection.end >= secondSection.end
	secondSectionCovering := secondSection.start <= firstSection.start && secondSection.end >= firstSection.end

	return firstSectionCovering || secondSectionCovering
}

func CheckOverlapping(firstSection Section, secondSection Section) bool {
	firstSectionOverlapping := (firstSection.start >= secondSection.start && firstSection.start <= secondSection.end) || (firstSection.end >= secondSection.start && firstSection.end <= secondSection.end)
	secondSectionOverlapping := (secondSection.start >= firstSection.start && secondSection.start <= firstSection.end) || (secondSection.end >= firstSection.start && secondSection.end <= firstSection.end)

	return firstSectionOverlapping || secondSectionOverlapping
}
