package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Motion struct {
	Direction string
	Size      int
}

func main() {
	motions, err := readMotions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(motions))
	fmt.Println(partTwo(motions))
}

func partOne(motions []Motion) int {
	return tailVisits(motions, 2)
}

func partTwo(motions []Motion) int {
	return tailVisits(motions, 10)
}

func tailVisits(motions []Motion, knots int) int {
	positions := make([][2]int, 10)
	for i := 0; i < knots; i++ {
		positions[i] = [2]int{0, 0}
	}
	visited := map[[2]int]bool{{0, 0}: true}
	for _, motion := range motions {
		for i := 0; i < motion.Size; i++ {
			positions = moveHead(positions, motion.Direction)
			for j := 1; j < knots; j++ {
				positions = moveFollowingKnot(positions, j)
				if j == knots-1 {
					visited[[2]int{positions[j][0], positions[j][1]}] = true
				}
			}
		}
	}
	return len(visited)
}

func moveHead(positions [][2]int, direction string) [][2]int {
	if direction == "R" {
		positions[0][0]++
	} else if direction == "L" {
		positions[0][0]--
	} else if direction == "U" {
		positions[0][1]++
	} else if direction == "D" {
		positions[0][1]--
	}
	return positions
}

func moveFollowingKnot(positions [][2]int, knotNumber int) [][2]int {
	xDifference := positions[knotNumber][0] - positions[knotNumber-1][0]
	yDifference := positions[knotNumber][1] - positions[knotNumber-1][1]
	if abs(xDifference) > 1 {
		positions[knotNumber][0] -= xDifference / abs(xDifference)
		if abs(yDifference) != 0 {
			positions[knotNumber][1] -= yDifference / abs(yDifference)
		}
	} else if abs(yDifference) > 1 {
		positions[knotNumber][1] -= yDifference / abs(yDifference)
		positions[knotNumber][0] = positions[knotNumber-1][0]
	}
	return positions
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func readMotions(fileName string) ([]Motion, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	motions := []Motion{}
	for scanner.Scan() {
		splitMotion := strings.Split(scanner.Text(), " ")
		size, err := strconv.Atoi(splitMotion[1])
		if err != nil {
			return nil, err
		}

		motions = append(motions, Motion{Direction: splitMotion[0], Size: size})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return motions, nil
}
