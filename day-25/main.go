package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	requirements, err := readRequirements("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(requirements))
}

func partOne(requirements [][]string) string {
	sum := 0
	for _, requirement := range requirements {
		number := convertSNAFUToNumber(requirement)
		sum += number
	}
	return convertNumberToSNAFU(sum)
}

func convertSNAFUToNumber(requirement []string) int {
	number := 0
	powerOfFive := 1
	for i := len(requirement) - 1; i >= 0; i-- {
		switch requirement[i] {
		case "=":
			number -= 2 * powerOfFive
		case "-":
			number -= powerOfFive
		case "1":
			number += powerOfFive
		case "2":
			number += 2 * powerOfFive
		}
		powerOfFive *= 5
	}
	return number
}

func convertNumberToSNAFU(number int) string {
	powerOfFive := 1
	snafuNumber := ""
	for number != 0 {
		value := (number % (powerOfFive * 5)) / powerOfFive
		if value >= 3 {
			value -= 5
		}
		switch value {
		case -2:
			snafuNumber = "=" + snafuNumber
		case -1:
			snafuNumber = "-" + snafuNumber
		case 0:
			snafuNumber = "0" + snafuNumber
		case 1:
			snafuNumber = "1" + snafuNumber
		case 2:
			snafuNumber = "2" + snafuNumber
		}
		number -= value * powerOfFive
		powerOfFive *= 5
	}
	return snafuNumber
}

func readRequirements(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	requirements := [][]string{}
	for scanner.Scan() {
		requirements = append(requirements, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return requirements, nil
}
