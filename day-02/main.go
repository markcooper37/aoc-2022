package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	moves, err := readMoves("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(moves))
	fmt.Println(partTwo(moves))
}

var (
	playerOneScoreMap = map[string]int{"A": 1, "B": 2, "C": 3}
	playerTwoScoreMap = map[string]int{"X": 1, "Y": 2, "Z": 3}
	playerTwoChoices  = []string{"X", "Y", "Z"}
)

func partOne(moves [][]string) int {
	score := 0
	for _, move := range moves {
		score += scoreMove(move[0], move[1])
	}
	return score
}

func partTwo(moves [][]string) int {
	score := 0
	for _, move := range moves {
		score += scoreMove(move[0], playerTwoChoices[(playerOneScoreMap[move[0]]+playerTwoScoreMap[move[1]])%3])
	}
	return score
}

func scoreMove(firstPlayerMove, secondPlayerMove string) int {
	return playerTwoScoreMap[secondPlayerMove] + ((playerTwoScoreMap[secondPlayerMove]-playerOneScoreMap[firstPlayerMove]+4)%3)*3
}

func readMoves(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	moves := [][]string{}
	for scanner.Scan() {
		moves = append(moves, strings.Split(scanner.Text(), " "))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}
