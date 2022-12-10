package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name  string
	Value int
}

func main() {
	instructions, err := readInstructions("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(instructions))
	partTwoResult := partTwo(instructions)
	for _, line := range partTwoResult {
		fmt.Println(line)
	}
}

func partOne(instructions []Instruction) int {
	registerValues := findRegisterValues(instructions)
	total := 0
	for i := 20; i <= 220; i += 40 {
		total += registerValues[i] * i
	}
	return total
}

func partTwo(instructions []Instruction) []string {
	registerValues := findRegisterValues(instructions)
	currentLine, drawing := "", []string{}
	for i := 1; i <= 240; i++ {
		if i%40 >= registerValues[i] && i%40 <= registerValues[i]+2 {
			currentLine += "#"
		} else {
			currentLine += "."
		}
		if i%40 == 0 {
			drawing = append(drawing, currentLine)
			currentLine = ""
		}
	}
	return drawing
}

func findRegisterValues(instructions []Instruction) []int {
	addValue := false
	instructionIndex := 0
	registerValue, registerValues := 1, []int{1}
	for instructionIndex < len(instructions) {
		registerValues = append(registerValues, registerValue)
		if instructions[instructionIndex].Name == "addx" {
			if addValue {
				registerValue += instructions[instructionIndex].Value
				instructionIndex++
			}
			addValue = !addValue
		} else {
			instructionIndex++
		}
	}
	return registerValues
}

func readInstructions(fileName string) ([]Instruction, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	instructions := []Instruction{}
	for scanner.Scan() {
		splitInstruction := strings.Split(scanner.Text(), " ")
		newInstruction := Instruction{Name: splitInstruction[0]}
		if len(splitInstruction) > 1 {
			value, err := strconv.Atoi(splitInstruction[1])
			if err != nil {
				return nil, err
			}

			newInstruction.Value = value
		}
		instructions = append(instructions, newInstruction)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return instructions, nil
}
