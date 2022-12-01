package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	calories, err := readCalories("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(calories))
	fmt.Println(partTwo(calories))
}

func partOne(calories [][]int) int {
	max := 0
	for _, elfCalories := range calories {
		sum := sum(elfCalories)
		if sum > max {
			max = sum
		}
	}
	return max
}

func partTwo(calories [][]int) int {
	topThree := []int{0, 0, 0}
	for _, elfCalories := range calories {
		sum := sum(elfCalories)
		for i := 2; i >= 0; i-- {
			if sum > topThree[i] {
				topThree = append(topThree[1:i+1], append([]int{sum}, topThree[i+1:]...)...)
				break
			}
		}
	}
	return topThree[0] + topThree[1] + topThree[2]
}

func sum(calories []int) int {
	sum := 0
	for _, calorieValue := range calories {
		sum += calorieValue
	}
	return sum
}

func readCalories(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	allElfCalories := [][]int{}
	elfCalories := []int{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			allElfCalories = append(allElfCalories, elfCalories)
			elfCalories = []int{}
		} else {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}

			elfCalories = append(elfCalories, calories)
		}
	}
	allElfCalories = append(allElfCalories, elfCalories)

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allElfCalories, nil
}
