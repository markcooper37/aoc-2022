package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type KeyLocations struct {
	Start [2]int
	End   [2]int
}

func main() {
	heightMap, keyLocations, err := readHeightmap("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(heightMap, keyLocations))
	fmt.Println(partTwo(heightMap, keyLocations))
}

func partOne(heightMap [][]byte, keyLocations KeyLocations) int {
	minDistances := findMinDistances(heightMap, keyLocations.Start, canMovePartOne)
	return minDistances[keyLocations.End]
}

func partTwo(heightMap [][]byte, keyLocations KeyLocations) int {
	minDistances := findMinDistances(heightMap, keyLocations.End, canMovePartTwo)
	minA := -1
	for key, value := range minDistances {
		if heightMap[key[0]][key[1]] == 'a' && (minA == -1 || value < minA) {
			minA = value
		}
	}
	return minA
}

func findMinDistances(heightMap [][]byte, start [2]int, canMove func(byte, byte) bool) map[[2]int]int {
	minDistances, locationsToCheck := map[[2]int]int{start: 0}, [][2]int{start}
	for iter := 1; iter <= len(heightMap)*len(heightMap[0]); iter++ {
		if len(locationsToCheck) == 0 {
			break
		}
		newLocations := [][2]int{}
		for _, location := range locationsToCheck {
			i, j := 1, 0
			for k := 0; k <= 3; k++ {
				if canMoveToNew(heightMap, location[0]+i, location[1]+j, location[0], location[1], canMove, minDistances) {
					minDistances[[2]int{location[0]+i, location[1]+j}] = iter
					newLocations = append(newLocations, [2]int{location[0]+i, location[1]+j})
				}
				i, j = -j, i
			}
		}
		locationsToCheck = newLocations
	}
	return minDistances
}

func canMoveToNew(heightMap [][]byte, newRow, newColumn, oldRow, oldColumn int, canMove func(byte, byte) bool, minDistances map[[2]int]int) bool {
	if newRow >= 0 && newRow < len(heightMap) && newColumn >= 0 && newColumn < len(heightMap[0]) {
		if canMove(heightMap[oldRow][oldColumn], heightMap[newRow][newColumn]) {
			if _, ok := minDistances[[2]int{newRow, newColumn}]; !ok {
				return true
			}
		}
	}
	return false
}

func canMovePartOne(current, new byte) bool {
	return current+1 >= new
}

func canMovePartTwo(current, new byte) bool {
	return current <= new+1
}

func readHeightmap(fileName string) ([][]byte, KeyLocations, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, KeyLocations{}, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	heightMapStrings := [][]string{}
	for scanner.Scan() {
		heightMapStrings = append(heightMapStrings, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, KeyLocations{}, err
	}

	keyLocations := KeyLocations{}
	heightMapBytes := [][]byte{}
	for rowIndex, row := range heightMapStrings {
		bytesRow := []byte{}
		for columnIndex, value := range row {
			if value == "S" {
				keyLocations.Start = [2]int{rowIndex, columnIndex}
				bytesRow = append(bytesRow, 'a')
			} else if value == "E" {
				keyLocations.End = [2]int{rowIndex, columnIndex}
				bytesRow = append(bytesRow, 'z')
			} else {
				bytesRow = append(bytesRow, value[0])
			}
		}
		heightMapBytes = append(heightMapBytes, bytesRow)
	}

	return heightMapBytes, keyLocations, nil
}
