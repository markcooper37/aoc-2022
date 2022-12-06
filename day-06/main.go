package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	signal, err := readSignal("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(signal))
	fmt.Println(partTwo(signal))
}

func partOne(signal string) int {
	return findMarker(signal, 4)
}

func partTwo(signal string) int {
	return findMarker(signal, 14)
}

func findMarker(signal string, unduplicatedLength int) int {
	for i := unduplicatedLength; i < len(signal); i++ {
		if !containsDuplicate(signal[i-unduplicatedLength : i]) {
			return i
		}
	}
	return -1
}

func containsDuplicate(sequence string) bool {
	characters := map[rune]bool{}
	for _, character := range sequence {
		characters[character] = true
	}
	return len(characters) != len(sequence)
}

func readSignal(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	return scanner.Text(), nil
}
