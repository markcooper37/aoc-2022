package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	blizzards, start, end, err := readBlizzards("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(blizzards, start, end))
	fmt.Println(partTwo(blizzards, start, end))
}

func partOne(blizzards [][]string, start, end [2]int) int {
	blizzardMap := constructBlizzardMap(blizzards)
	_, moves := findFewestMoves(blizzardMap, blizzards, start, end)
	return moves
}

func constructBlizzardMap(blizzards [][]string) map[[3]int]bool {
	blizzardMap := map[[3]int]bool{}
	for rowIndex, row := range blizzards {
		for columnIndex, position := range row {
			if position == ">" {
				blizzardMap[[3]int{rowIndex, columnIndex, 0}] = true
			} else if position == "v" {
				blizzardMap[[3]int{rowIndex, columnIndex, 1}] = true
			} else if position == "<" {
				blizzardMap[[3]int{rowIndex, columnIndex, 2}] = true
			} else if position == "^" {
				blizzardMap[[3]int{rowIndex, columnIndex, 3}] = true
			}
		}
	}
	return blizzardMap
}

func findFewestMoves(blizzardMap map[[3]int]bool, blizzards [][]string, start [2]int, end [2]int) (map[[3]int]bool, int) {
	positions := map[[2]int]bool{start: true}
	for i := 1; i > 0; i++ {
		newBlizzardMap, blizzardPositions := newBlizzards(blizzardMap, blizzards)
		newPositions := map[[2]int]bool{}
		found := false
		for position := range positions {
			moves := findMoves(position, blizzardPositions, blizzards)
			for _, move := range moves {
				if move == end {
					found = true
				}
				newPositions[move] = true
			}
		}
		if found {
			return newBlizzardMap, i
		}
		blizzardMap = newBlizzardMap
		positions = newPositions
	}
	return nil, -1
}

func findMoves(position [2]int, blizzardPositions map[[2]int]bool, blizzards [][]string) [][2]int {
	moves := [][2]int{}
	if !blizzardPositions[position] {
		moves = append(moves, position)
	}
	possibleMoves := [][2]int{{position[0] + 1, position[1]}, {position[0] - 1, position[1]}, {position[0], position[1] + 1}, {position[0], position[1] - 1}}
	for _, possibleMove := range possibleMoves {
		if possibleMove[0] >= 0 && possibleMove[0] <= len(blizzards)-1 && possibleMove[1] >= 0 && possibleMove[1] <= len(blizzards[0])-1 {
			if blizzards[possibleMove[0]][possibleMove[1]] != "#" && !blizzardPositions[possibleMove] {
				moves = append(moves, possibleMove)
			}
		}
	}
	return moves
}

func newBlizzards(blizzardMap map[[3]int]bool, blizzards [][]string) (map[[3]int]bool, map[[2]int]bool) {
	newBlizzardMap := map[[3]int]bool{}
	blizzardPositions := map[[2]int]bool{}
	for position := range blizzardMap {
		newPosition := findNewPosition(position, blizzards)
		newBlizzardMap[newPosition] = true
		blizzardPositions[[2]int{newPosition[0], newPosition[1]}] = true
	}
	return newBlizzardMap, blizzardPositions
}

func findNewPosition(position [3]int, blizzards [][]string) [3]int {
	if position[2] == 0 {
		if position[1]+1 == len(blizzards[0])-1 {
			return [3]int{position[0], 1, position[2]}
		} else {
			return [3]int{position[0], position[1] + 1, position[2]}
		}
	} else if position[2] == 1 {
		if position[0]+1 == len(blizzards)-1 {
			return [3]int{1, position[1], position[2]}
		} else {
			return [3]int{position[0] + 1, position[1], position[2]}
		}
	} else if position[2] == 2 {
		if position[1]-1 == 0 {
			return [3]int{position[0], len(blizzards[0]) - 2, position[2]}
		} else {
			return [3]int{position[0], position[1] - 1, position[2]}
		}
	} else {
		if position[0]-1 == 0 {
			return [3]int{len(blizzards) - 2, position[1], position[2]}
		} else {
			return [3]int{position[0] - 1, position[1], position[2]}
		}
	}
}

func partTwo(blizzards [][]string, start, end [2]int) int {
	blizzardMap := constructBlizzardMap(blizzards)
	moves := 0
	newBlizzardMap, newMoves := findFewestMoves(blizzardMap, blizzards, start, end)
	moves += newMoves
	newBlizzardMap, newMoves = findFewestMoves(newBlizzardMap, blizzards, end, start)
	moves += newMoves
	_, newMoves = findFewestMoves(newBlizzardMap, blizzards, start, end)
	return moves + newMoves
}

func readBlizzards(fileName string) ([][]string, [2]int, [2]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, [2]int{}, [2]int{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	blizzards := [][]string{}
	for scanner.Scan() {
		blizzards = append(blizzards, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, [2]int{}, [2]int{}, err
	}

	start, end := [2]int{0}, [2]int{len(blizzards) - 1}
	for index, position := range blizzards[0] {
		if position == "." {
			start[1] = index
		}
	}
	for index, position := range blizzards[len(blizzards)-1] {
		if position == "." {
			end[1] = index
		}
	}
	return blizzards, start, end, nil
}
