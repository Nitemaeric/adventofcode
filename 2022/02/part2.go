package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Hand struct {
	name  string
	wins  string
	loses string
}

var Rock = Hand{name: "rock", wins: "scissors", loses: "paper"}
var Paper = Hand{name: "paper", wins: "rock", loses: "scissors"}
var Scissors = Hand{name: "scissors", wins: "paper", loses: "rock"}

var inputMap = map[string]Hand{
	"A": Rock, "rock": Rock,
	"B": Paper, "paper": Paper,
	"C": Scissors, "scissors": Scissors,
}

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
	firstPlayerHand := inputMap[firstPlayerInput]
	secondPlayerHand, _ := findHand(firstPlayerHand, secondPlayerInput)

	firstPlayerScore, secondPlayerScore := calculateWinningsScore(firstPlayerHand, secondPlayerHand)

	firstPlayerScore += calculateInputScore(firstPlayerHand)
	secondPlayerScore += calculateInputScore(secondPlayerHand)

	return firstPlayerScore, secondPlayerScore
}

func calculateInputScore(input Hand) int64 {
	scoreMap := map[Hand]int64{
		Rock:     1,
		Paper:    2,
		Scissors: 3,
	}

	return scoreMap[input]
}

func calculateWinningsScore(firstPlayerHand Hand, secondPlayerHand Hand) (int64, int64) {
	scoreMap := map[Hand]map[Hand][]int64{
		Rock: {
			Rock:     {3, 3},
			Paper:    {0, 6},
			Scissors: {6, 0},
		},
		Paper: {
			Rock:     {6, 0},
			Paper:    {3, 3},
			Scissors: {0, 6},
		},
		Scissors: {
			Rock:     {0, 6},
			Paper:    {6, 0},
			Scissors: {3, 3},
		},
	}

	winningScores := scoreMap[firstPlayerHand][secondPlayerHand]

	return winningScores[0], winningScores[1]
}

func findHand(firstPlayerHand Hand, expectedResultInput string) (Hand, error) {
	switch expectedResultInput {
	case "X":
		return inputMap[firstPlayerHand.wins], nil
	case "Y":
		return inputMap[firstPlayerHand.name], nil
	case "Z":
		return inputMap[firstPlayerHand.loses], nil
	default:
		return Hand{}, errors.New("No Hand Found")
	}
}

func convertInputToKind(input string) string {
	return inputMap[input].name
}
