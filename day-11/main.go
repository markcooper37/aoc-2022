package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items            []int
	Operation        WorryOperation
	DivisibilityTest int
	TrueMonkey       int
	FalseMonkey      int
}

type WorryOperation struct {
	Operation string
	Old       bool
	Value     *int
}

func main() {
	monkeys, err := readMonkeys("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(monkeys))

	monkeys, err = readMonkeys("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(partTwo(monkeys))
}

func partOne(monkeys []Monkey) int {
	counts := performIterations(monkeys, func(worry int) int { return worry / 3 }, 20)
	return calculateMonkeyBusiness(counts)
}

func partTwo(monkeys []Monkey) int {
	modValue := 1
	for _, monkey := range monkeys {
		modValue *= monkey.DivisibilityTest
	}
	counts := performIterations(monkeys, func(worry int) int { return worry % modValue }, 10000)
	return calculateMonkeyBusiness(counts)
}

func performIterations(monkeys []Monkey, manageWorry func(int) int, iterations int) []int {
	counts := make([]int, len(monkeys))
	for i := 0; i < iterations; i++ {
		for monkeyIndex, monkey := range monkeys {
			for _, item := range monkey.Items {
				counts[monkeyIndex]++
				newItem := performOperation(item, monkey.Operation)
				newItem = manageWorry(newItem)
				if newItem%monkey.DivisibilityTest == 0 {
					monkeys[monkey.TrueMonkey].Items = append(monkeys[monkey.TrueMonkey].Items, newItem)
				} else {
					monkeys[monkey.FalseMonkey].Items = append(monkeys[monkey.FalseMonkey].Items, newItem)
				}
			}
			monkeys[monkeyIndex].Items = []int{}
		}
	}
	return counts
}

func performOperation(value int, operation WorryOperation) int {
	secondValue := value
	if !operation.Old {
		secondValue = *operation.Value
	}
	switch operation.Operation {
	case "+":
		return value + secondValue
	case "*":
		return value * secondValue
	default:
		return value
	}
}

func calculateMonkeyBusiness(counts []int) int {
	sort.Ints(counts)
	return counts[len(counts)-1] * counts[len(counts)-2]
}

func readMonkeys(fileName string) ([]Monkey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	monkeys := []Monkey{}
	monkey := Monkey{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		}
		trimmedText := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(trimmedText, "Starting items") {
			itemsStrings := strings.Split(strings.TrimPrefix(trimmedText, "Starting items: "), ", ")
			items := []int{}
			for _, itemString := range itemsStrings {
				item, err := strconv.Atoi(itemString)
				if err != nil {
					return nil, err
				}

				items = append(items, item)
			}
			monkey.Items = items
		} else if strings.HasPrefix(trimmedText, "Operation") {
			splitOperation := strings.Split(strings.TrimPrefix(trimmedText, "Operation: new = old "), " ")
			monkey.Operation = WorryOperation{Operation: splitOperation[0]}
			if splitOperation[1] == "old" {
				monkey.Operation.Old = true
			} else {
				value, err := strconv.Atoi(splitOperation[1])
				if err != nil {
					return nil, err
				}

				monkey.Operation.Value = &value
			}
		} else if strings.HasPrefix(trimmedText, "Test") {
			divisibilityTest, err := strconv.Atoi(strings.TrimPrefix(trimmedText, "Test: divisible by "))
			if err != nil {
				return nil, err
			}

			monkey.DivisibilityTest = divisibilityTest
		} else if strings.HasPrefix(trimmedText, "If true") {
			trueMonkey, err := strconv.Atoi(strings.TrimPrefix(trimmedText, "If true: throw to monkey "))
			if err != nil {
				return nil, err
			}

			monkey.TrueMonkey = trueMonkey
		} else if strings.HasPrefix(trimmedText, "If false") {
			falseMonkey, err := strconv.Atoi(strings.TrimPrefix(trimmedText, "If false: throw to monkey "))
			if err != nil {
				return nil, err
			}

			monkey.FalseMonkey = falseMonkey
		}
	}
	monkeys = append(monkeys, monkey)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return monkeys, nil
}
