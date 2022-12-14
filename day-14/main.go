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
	rocks, err := readRocks("test_data.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(rocks))
	fmt.Println(partTwo(rocks))
}

func partOne(rocks [][][2]int) int {
	rockMap := constructRockMap(rocks)
	lowestVerticalPosition := 0
	for key := range rockMap {
		if key[1] > lowestVerticalPosition {
			lowestVerticalPosition = key[1]
		}
	}
	sandCount, currentPosition := 0, [2]int{500, 0}
	for {
		if currentPosition[1] > lowestVerticalPosition {
			break
		}
		if !rockMap[[2]int{currentPosition[0], currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0], currentPosition[1] + 1}
		} else if !rockMap[[2]int{currentPosition[0] - 1, currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0] - 1, currentPosition[1] + 1}
		} else if !rockMap[[2]int{currentPosition[0] + 1, currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0] + 1, currentPosition[1] + 1}
		} else {
			rockMap[currentPosition] = true
			sandCount++
			currentPosition = [2]int{500, 0}
		}
	}
	return sandCount
}

func partTwo(rocks [][][2]int) int {
	rockMap := constructRockMap(rocks)
	floor := 0
	for key := range rockMap {
		if key[1]+2 > floor {
			floor = key[1] + 2
		}
	}
	sandCount, currentPosition := 0, [2]int{500, 0}
	for {
		if rockMap[currentPosition] {
			break
		}
		if currentPosition[1]+1 == floor {
			rockMap[currentPosition] = true
			sandCount++
			currentPosition = [2]int{500, 0}
		} else if !rockMap[[2]int{currentPosition[0], currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0], currentPosition[1] + 1}
		} else if !rockMap[[2]int{currentPosition[0] - 1, currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0] - 1, currentPosition[1] + 1}
		} else if !rockMap[[2]int{currentPosition[0] + 1, currentPosition[1] + 1}] {
			currentPosition = [2]int{currentPosition[0] + 1, currentPosition[1] + 1}
		} else {
			rockMap[currentPosition] = true
			sandCount++
			currentPosition = [2]int{500, 0}
		}
	}
	return sandCount
}

func constructRockMap(rocks [][][2]int) map[[2]int]bool {
	rockMap := map[[2]int]bool{}
	for _, rockSequence := range rocks {
		for i := 0; i < len(rockSequence)-1; i++ {
			rockMap = addRocks(rockMap, rockSequence[i], rockSequence[i+1])
		}
	}
	return rockMap
}

func addRocks(rockMap map[[2]int]bool, positionOne, positionTwo [2]int) map[[2]int]bool {
	if positionOne[0] > positionTwo[0] {
		for i := positionTwo[0]; i <= positionOne[0]; i++ {
			rockMap[[2]int{i, positionTwo[1]}] = true
		}
	} else if positionTwo[0] > positionOne[0] {
		for i := positionOne[0]; i <= positionTwo[0]; i++ {
			rockMap[[2]int{i, positionOne[1]}] = true
		}
	} else if positionOne[1] > positionTwo[1] {
		for i := positionTwo[1]; i <= positionOne[1]; i++ {
			rockMap[[2]int{positionTwo[0], i}] = true
		}
	} else if positionTwo[1] > positionOne[1] {
		for i := positionOne[1]; i <= positionTwo[1]; i++ {
			rockMap[[2]int{positionOne[0], i}] = true
		}
	}
	return rockMap
}

func readRocks(fileName string) ([][][2]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	rocks := [][][2]int{}
	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), " -> ")
		newRocks := [][2]int{}
		for _, pair := range pairs {
			values := strings.Split(pair, ",")
			rockPair := [2]int{}
			for i := 0; i <= 1; i++ {
				value, err := strconv.Atoi(values[i])
				if err != nil {
					return nil, err
				}

				rockPair[i] = value
			}
			newRocks = append(newRocks, rockPair)
		}
		rocks = append(rocks, newRocks)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rocks, nil
}
