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
	stacks, instructions, err := readCrates("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(stacks, instructions))
	fmt.Println(partTwo(stacks, instructions))
}

func partOne(stacks [][]byte, instructions [][]int) string {
	stacks = copyStacks(stacks)
	for _, instruction := range instructions {
		for i := 0; i < instruction[0]; i++ {
			stacks = moveStack(stacks, instruction[1]-1, instruction[2]-1, 1)
		}
	}
	return getDesiredCrates(stacks)
}

func partTwo(stacks [][]byte, instructions [][]int) string {
	stacks = copyStacks(stacks)
	for _, instruction := range instructions {
		stacks = moveStack(stacks, instruction[1]-1, instruction[2]-1, instruction[0])
	}
	return getDesiredCrates(stacks)
}

func copyStacks(stacks [][]byte) [][]byte {
	newStacks := make([][]byte, len(stacks))
	for index, stack := range stacks {
		newStacks[index] = append(newStacks[index], stack...)
	}
	return newStacks
}

func moveStack(stacks [][]byte, startStack, endStack, number int) [][]byte {
	values := stacks[startStack][len(stacks[startStack])-number:]
	stacks[startStack] = stacks[startStack][:len(stacks[startStack])-number]
	stacks[endStack] = append(stacks[endStack], values...)
	return stacks
}

func getDesiredCrates(stacks [][]byte) string {
	output := ""
	for _, stack := range stacks {
		output += string(stack[len(stack)-1])
	}
	return output
}

func readCrates(fileName string) ([][]byte, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	stacks := make([][]byte, 9)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		} else if scanner.Text()[1] == '1' {
			continue
		}
		for i := 1; i < len(scanner.Text()); i += 4 {
			if scanner.Text()[i] != ' ' {
				stacks[(i-1)/4] = append([]byte{scanner.Text()[i]}, stacks[(i-1)/4]...)
			}
		}
	}
	for i := 8; i >= 0; i-- {
		if len(stacks[i]) == 0 {
			stacks = stacks[:i]
		} else {
			break
		}
	}
	instructions := [][]int{}
	for scanner.Scan() {
		newInstructions := []int{}
		splitInstructions := strings.Split(scanner.Text(), " ")
		for i := 1; i <= 5; i += 2 {
			intInstruction, err := strconv.Atoi(splitInstructions[i])
			if err != nil {
				return nil, nil, err
			}

			newInstructions = append(newInstructions, intInstruction)
		}
		instructions = append(instructions, newInstructions)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return stacks, instructions, nil
}
