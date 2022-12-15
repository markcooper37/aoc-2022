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
	positions, err := readPositions("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(partOne(positions, 2000000))
	fmt.Println(partTwo(positions, 4000000))
}

func partOne(positions [][]int, row int) int {
	sensors := map[[2]int]int{}
	for _, positionGroup := range positions {
		sensors[[2]int{positionGroup[0], positionGroup[1]}] = manhattanDistance(positionGroup[0], positionGroup[1], positionGroup[2], positionGroup[3])
	}
	invalidPositions := map[int]bool{}
	for sensor, distance := range sensors {
		maxXDifference := distance - abs(row-sensor[1])
		for i := -maxXDifference; i <= maxXDifference; i++ {
			invalidPositions[sensor[0]+i] = true

		}
	}
	for _, positionGroup := range positions {
		if positionGroup[3] == row {
			delete(invalidPositions, positionGroup[2])
		}
	}
	return len(invalidPositions)
}

func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func partTwo(positions [][]int, searchArea int) int {
	sensors := map[[2]int]int{}
	for _, positionGroup := range positions {
		sensors[[2]int{positionGroup[0], positionGroup[1]}] = manhattanDistance(positionGroup[0], positionGroup[1], positionGroup[2], positionGroup[3])
	}
	for sensor, distance := range sensors {
		for i := -distance - 1; i <= distance+1; i++ {
			x1, y1 := sensor[0]+i, sensor[1]+distance+1-abs(i)
			if canContainBeacon(x1, y1, searchArea, sensors) {
				return 4000000*x1 + y1
			}
			x2, y2 := sensor[0]+i, sensor[1]-distance-1+abs(i)
			if canContainBeacon(x2, y2, searchArea, sensors) {
				return 4000000*x2 + y2
			}
		}
	}
	return -1
}

func canContainBeacon(x, y, searchArea int, sensors map[[2]int]int) bool {
	if x < 0 || x > searchArea || y < 0 || y > searchArea {
		return false
	}
	for secondSensor, distance := range sensors {
		if manhattanDistance(secondSensor[0], secondSensor[1], x, y) <= distance {
			return false
		}
	}
	return true
}

func readPositions(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	positions := [][]int{}
	for scanner.Scan() {
		splitString := strings.Split(strings.ReplaceAll(strings.ReplaceAll(scanner.Text(), ":", ""), ",", ""), " ")
		newPositions := []int{}
		requiredValues := []string{splitString[2], splitString[3], splitString[8], splitString[9]}
		for _, value := range requiredValues {
			intString := value[2:]
			position, err := strconv.Atoi(intString)
			if err != nil {
				return nil, err
			}

			newPositions = append(newPositions, position)
		}
		positions = append(positions, newPositions)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
