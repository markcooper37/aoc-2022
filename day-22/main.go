package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	tiles, instructions, err := readTiles("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(tiles, instructions))
	fmt.Println(partTwo(tiles, instructions))
}

func partOne(tiles [][]string, instructions string) int {
	currentDirection := 0
	currentRow := 0
	currentColumn := 0
	for columnIndex, tile := range tiles[0] {
		if tile == "." {
			currentColumn = columnIndex
			break
		}
	}
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'R' {
			currentDirection = (currentDirection + 1) % 4
		} else if instructions[i] == 'L' {
			currentDirection = (currentDirection + 3) % 4
		} else {
			endIndex := len(instructions)
			for j := i + 1; j < len(instructions); j++ {
				if instructions[j] < '0' || instructions[j] > '9' {
					endIndex = j
					break
				}
			}
			moves, err := strconv.Atoi(instructions[i:endIndex])
			if err != nil {
				return -1
			}

			currentRow, currentColumn = move(currentRow, currentColumn, currentDirection, moves, tiles)
			i = endIndex - 1
		}
	}
	return 1000*(currentRow+1) + 4*(currentColumn+1) + currentDirection
}

func move(currentRow, currentColumn, currentDirection, moves int, tiles [][]string) (int, int) {
	for i := 0; i < moves; i++ {
		newRow, newColumn := findNewPosition(currentRow, currentColumn, currentDirection, tiles)
		if tiles[newRow][newColumn] == "#" {
			return currentRow, currentColumn
		} else {
			currentRow, currentColumn = newRow, newColumn
		}
	}
	return currentRow, currentColumn
}

func findNewPosition(currentRow, currentColumn, currentDirection int, tiles [][]string) (int, int) {
	if currentDirection == 0 {
		for i := 0; i < len(tiles[0]); i++ {
			currentColumn = (currentColumn + 1) % len(tiles[0])
			if tiles[currentRow][currentColumn] != " " {
				return currentRow, currentColumn
			}
		}
		return -1, -1
	} else if currentDirection == 1 {
		for i := 0; i < len(tiles); i++ {
			currentRow = (currentRow + 1) % len(tiles)
			if tiles[currentRow][currentColumn] != " " {
				return currentRow, currentColumn
			}
		}
		return -1, -1
	} else if currentDirection == 2 {
		for i := 0; i < len(tiles[0]); i++ {
			currentColumn = (currentColumn + len(tiles[0]) - 1) % len(tiles[0])
			if tiles[currentRow][currentColumn] != " " {
				return currentRow, currentColumn
			}
		}
		return -1, -1
	} else {
		for i := 0; i < len(tiles); i++ {
			currentRow = (currentRow + len(tiles) - 1) % len(tiles)
			if tiles[currentRow][currentColumn] != " " {
				return currentRow, currentColumn
			}
		}
		return -1, -1
	}
}

func partTwo(tiles [][]string, instructions string) int {
	newPositions := constructCube(tiles)
	currentDirection := 0
	currentRow := 0
	currentColumn := 0
	for columnIndex, tile := range tiles[0] {
		if tile == "." {
			currentColumn = columnIndex
			break
		}
	}
	for i := 0; i < len(instructions); i++ {
		if instructions[i] == 'R' {
			currentDirection = (currentDirection + 1) % 4
		} else if instructions[i] == 'L' {
			currentDirection = (currentDirection + 3) % 4
		} else {
			endIndex := len(instructions)
			for j := i + 1; j < len(instructions); j++ {
				if instructions[j] < '0' || instructions[j] > '9' {
					endIndex = j
					break
				}
			}
			moves, err := strconv.Atoi(instructions[i:endIndex])
			if err != nil {
				return -1
			}

			currentRow, currentColumn, currentDirection = move2(currentRow, currentColumn, currentDirection, moves, tiles, newPositions)
			i = endIndex - 1
		}
	}
	return 1000*(currentRow+1) + 4*(currentColumn+1) + currentDirection
}

func constructCube(tiles [][]string) map[[3]int][3]int {
	newPositions := map[[3]int][3]int{}
	for i := 0; i < 50; i++ {
		newPositions[[3]int{50 + i, 50, 2}] = [3]int{100, i, 1}
		newPositions[[3]int{100, i, 3}] = [3]int{50 + i, 50, 0}

		newPositions[[3]int{49, 100 + i, 1}] = [3]int{50 + i, 99, 2}
		newPositions[[3]int{50 + i, 99, 0}] = [3]int{49, 100 + i, 3}

		newPositions[[3]int{49 - i, 149, 0}] = [3]int{100 + i, 99, 2}
		newPositions[[3]int{100 + i, 99, 0}] = [3]int{49 - i, 149, 2}

		newPositions[[3]int{i, 50, 2}] = [3]int{149 - i, 0, 0}
		newPositions[[3]int{149 - i, 0, 2}] = [3]int{i, 50, 0}

		newPositions[[3]int{149, 50 + i, 1}] = [3]int{150 + i, 49, 2}
		newPositions[[3]int{150 + i, 49, 0}] = [3]int{149, 50 + i, 3}

		newPositions[[3]int{0, 99 - i, 3}] = [3]int{199 - i, 0, 0}
		newPositions[[3]int{199 - i, 0, 2}] = [3]int{0, 99 - i, 1}

		newPositions[[3]int{0, 149 - i, 3}] = [3]int{199, 49 - i, 3}
		newPositions[[3]int{199, 49 - i, 1}] = [3]int{0, 149 - i, 1}
	}
	return newPositions
}

func move2(currentRow, currentColumn, currentDirection, moves int, tiles [][]string, newPositions map[[3]int][3]int) (int, int, int) {
	for i := 0; i < moves; i++ {
		newRow, newColumn, newDirection := findNewPosition2(currentRow, currentColumn, currentDirection, tiles, newPositions)
		if tiles[newRow][newColumn] == "#" {
			return currentRow, currentColumn, currentDirection
		} else {
			currentRow, currentColumn, currentDirection = newRow, newColumn, newDirection
		}
	}
	return currentRow, currentColumn, currentDirection
}

func findNewPosition2(currentRow, currentColumn, currentDirection int, tiles [][]string, newPositions map[[3]int][3]int) (int, int, int) {
	if newPosition, ok := newPositions[[3]int{currentRow, currentColumn, currentDirection}]; ok {
		return newPosition[0], newPosition[1], newPosition[2]
	} else if currentDirection == 0 {
		return currentRow, currentColumn + 1, currentDirection
	} else if currentDirection == 1 {
		return currentRow + 1, currentColumn, currentDirection
	} else if currentDirection == 2 {
		return currentRow, currentColumn - 1, currentDirection
	} else {
		return currentRow - 1, currentColumn, currentDirection
	}
}

func readTiles(fileName string) ([][]string, string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, "", err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	tiles := [][]string{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		tiles = append(tiles, strings.Split(scanner.Text(), ""))
	}
	scanner.Scan()
	instructions := scanner.Text()
	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	maxRowLength := 0
	for _, row := range tiles {
		if len(row) > maxRowLength {
			maxRowLength = len(row)
		}
	}
	for index, row := range tiles {
		for i := len(row); i < maxRowLength; i++ {
			tiles[index] = append(tiles[index], " ")
		}
	}
	return tiles, instructions, nil
}
