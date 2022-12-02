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

	points := []int64{0, 0}

	for fileScanner.Scan() {
		inputs := strings.Split(fileScanner.Text(), " ")

		firstPlayerScore, secondPlayerScore := calculateResult(inputs[0], inputs[1])

		points[0] += firstPlayerScore
		points[1] += secondPlayerScore
	}

	resultsString := fmt.Sprintf("Results: %d vs %d", points[0], points[1])

	fmt.Println(resultsString)

	file.Close()
}

func calculateResult(firstPlayerInput string, secondPlayerInput string) (int64, int64) {
	firstPlayerScore, secondPlayerScore := calculateWinningsScore(firstPlayerInput, secondPlayerInput)

	firstPlayerScore += calculateInputScore(firstPlayerInput)
	secondPlayerScore += calculateInputScore(secondPlayerInput)

	return firstPlayerScore, secondPlayerScore
}

func calculateInputScore(input string) int64 {
	scoreMap := map[string]int64{
		"rock":     1,
		"paper":    2,
		"scissors": 3,
	}

	return scoreMap[convertInputToKind(input)]
}

func calculateWinningsScore(firstPlayerInput string, secondPlayerInput string) (int64, int64) {
	scoreMap := map[string][]int64{
		"rockrock":         {3, 3},
		"rockpaper":        {0, 6},
		"rockscissors":     {6, 0},
		"paperrock":        {6, 0},
		"paperpaper":       {3, 3},
		"paperscissors":    {0, 6},
		"scissorsrock":     {0, 6},
		"scissorspaper":    {6, 0},
		"scissorsscissors": {3, 3},
	}

	winningScores := scoreMap[convertInputToKind(firstPlayerInput)+convertInputToKind(secondPlayerInput)]

	return winningScores[0], winningScores[1]
}

func convertInputToKind(input string) string {
	if isRock(input) {
		return "rock"
	} else if isPaper(input) {
		return "paper"
	} else if isScissors(input) {
		return "scissors"
	} else {
		return ""
	}
}

func isRock(input string) bool {
	return input == "A" || input == "X"
}

func isPaper(input string) bool {
	return input == "B" || input == "Y"
}

func isScissors(input string) bool {
	return input == "C" || input == "Z"
}
