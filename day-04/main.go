package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	LowerID int
	UpperID int
}

func main() {
	assignments, err := readAssignments("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(assignments))
	fmt.Println(partTwo(assignments))
}

func partOne(assignments [][2]Assignment) int {
	count := 0
	for _, assignmentPair := range assignments {
		if assignmentContains(assignmentPair[0], assignmentPair[1]) || assignmentContains(assignmentPair[1], assignmentPair[0]) {
			count++
		}
	}
	return count
}

func assignmentContains(assignmentOne, assignmentTwo Assignment) bool {
	return assignmentOne.LowerID <= assignmentTwo.LowerID && assignmentOne.UpperID >= assignmentTwo.UpperID
}

func partTwo(assignments [][2]Assignment) int {
	count := 0
	for _, assignmentPair := range assignments {
		if assignmentsOverlap(assignmentPair[0], assignmentPair[1]) {
			count++
		}
	}
	return count
}

func assignmentsOverlap(assignmentOne, assignmentTwo Assignment) bool {
	return assignmentOne.LowerID <= assignmentTwo.UpperID && assignmentOne.UpperID >= assignmentTwo.LowerID
}

func readAssignments(fileName string) ([][2]Assignment, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	allAssignments := [][2]Assignment{}
	for scanner.Scan() {
		newAssignments := [2]Assignment{}
		assignments := strings.Split(scanner.Text(), ",")
		for i := 0; i <= 1; i++ {
			bounds := strings.Split(assignments[i], "-")
			lowerID, err := strconv.Atoi(bounds[0])
			if err != nil {
				return nil, err
			}

			upperID, err := strconv.Atoi(bounds[1])
			if err != nil {
				return nil, err
			}

			newAssignments[i] = Assignment{LowerID: lowerID, UpperID: upperID}
		}
		allAssignments = append(allAssignments, newAssignments)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allAssignments, nil
}
