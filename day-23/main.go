package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	positions, err := readElves("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(positions))
	positions, err = readElves("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partTwo(positions))
}

func partOne(positions map[[2]int]bool) int {
	for i := 0; i < 10; i++ {
		proposedPositions := map[[2]int][2]int{}
		proposedPositionCount := map[[2]int]int{}
		for position := range positions {
			if notSurrounded(position, positions) {
				proposedPositions[position] = position
				proposedPositionCount[position]++
				continue
			}
			newPositionFound := false
			for j := 0; j < 4; j++ {
				if canMove(position, positions, (i+j)%4) {
					newPosition := findNewPosition(position, (i+j)%4)
					proposedPositions[position] = newPosition
					proposedPositionCount[newPosition]++
					newPositionFound = true
					break
				}
			}
			if !newPositionFound {
				proposedPositions[position] = position
				proposedPositionCount[position]++
			}
		}
		newPositions := map[[2]int]bool{}
		for position := range positions {
			proposedPosition := proposedPositions[position]
			if proposedPositionCount[proposedPosition] > 1 {
				newPositions[position] = true
			} else {
				newPositions[proposedPosition] = true
			}
		}
		positions = newPositions
	}
	var minRow, maxRow, minCol, maxCol int
	for position := range positions {
		minRow, maxRow, minCol, maxCol = position[0], position[0], position[1], position[1]
		break
	}
	for position := range positions {
		if position[0] < minRow {
			minRow = position[0]
		} else if position[0] > maxRow {
			maxRow = position[0]
		}
		if position[1] < minCol {
			minCol = position[1]
		} else if position[1] > maxCol {
			maxCol = position[1]
		}
	}
	return (maxRow - minRow + 1) * (maxCol - minCol + 1) - len(positions)
}

func findNewPosition(position [2]int, direction int) [2]int {
	if direction == 0 {
		return [2]int{position[0] + 1, position[1]}
	} else if direction == 1 {
		return [2]int{position[0] - 1, position[1]}
	} else if direction == 2 {
		return [2]int{position[0], position[1] - 1}
	} else {
		return [2]int{position[0], position[1] + 1}
	}
}

func canMove(position [2]int, positions map[[2]int]bool, direction int) bool {
	row, col := position[0], position[1]
	if direction == 0 {
		return !positions[[2]int{row + 1, col - 1}] && !positions[[2]int{row + 1, col}] && !positions[[2]int{row + 1, col + 1}]
	} else if direction == 1 {
		return !positions[[2]int{row - 1, col - 1}] && !positions[[2]int{row - 1, col}] && !positions[[2]int{row - 1, col + 1}]
	} else if direction == 2 {
		return !positions[[2]int{row - 1, col - 1}] && !positions[[2]int{row, col - 1}] && !positions[[2]int{row + 1, col - 1}]
	} else {
		return !positions[[2]int{row - 1, col + 1}] && !positions[[2]int{row, col + 1}] && !positions[[2]int{row + 1, col + 1}]
	}
}

func notSurrounded(position [2]int, positions map[[2]int]bool) bool {
	row, col := position[0], position[1]
	return !positions[[2]int{row - 1, col - 1}] && !positions[[2]int{row - 1, col}] &&
		!positions[[2]int{row - 1, col + 1}] && !positions[[2]int{row, col - 1}] &&
		!positions[[2]int{row, col + 1}] && !positions[[2]int{row + 1, col - 1}] &&
		!positions[[2]int{row + 1, col}] && !positions[[2]int{row + 1, col + 1}]
}

func partTwo(positions map[[2]int]bool) int {
	for i := 0; i >= 0; i++ {
		proposedPositions := map[[2]int][2]int{}
		proposedPositionCount := map[[2]int]int{}
		for position := range positions {
			if notSurrounded(position, positions) {
				proposedPositions[position] = position
				proposedPositionCount[position]++
				continue
			}
			newPositionFound := false
			for j := 0; j < 4; j++ {
				if canMove(position, positions, (i+j)%4) {
					newPosition := findNewPosition(position, (i+j)%4)
					proposedPositions[position] = newPosition
					proposedPositionCount[newPosition]++
					newPositionFound = true
					break
				}
			}
			if !newPositionFound {
				proposedPositions[position] = position
				proposedPositionCount[position]++
			}
		}
		newPositions := map[[2]int]bool{}
		moved := false
		for position := range positions {
			proposedPosition := proposedPositions[position]
			if proposedPositionCount[proposedPosition] > 1 {
				newPositions[position] = true
			} else {
				newPositions[proposedPosition] = true
				if proposedPosition != position {
					moved = true
				}
			}
		}
		if !moved {
			return i + 1
		}
		positions = newPositions
	}
	return -1
}

func readElves(fileName string) (map[[2]int]bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	elves := [][]string{}
	for scanner.Scan() {
		elves = append(elves, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	positions := map[[2]int]bool{}
	for rowIndex, row := range elves {
		for columnIndex, position := range row {
			if position == "#" {
				positions[[2]int{-rowIndex, columnIndex}] = true
			}
		}
	}
	return positions, nil
}
