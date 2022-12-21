package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var operationMap = map[string]string{"+": "-", "-": "+", "*": "/", "/": "*"}

type Monkey struct {
	Value     *int
	Operation []string
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

	partTwo, err := partTwo(monkeys)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partTwo)
}

func partOne(monkeys map[string]Monkey) int {
	for monkeys["root"].Value == nil {
		for name, monkey := range monkeys {
			updateMonkeyValue(name, monkey, monkeys)
		}
	}
	return *monkeys["root"].Value
}

func updateMonkeyValue(name string, monkey Monkey, monkeys map[string]Monkey) {
	if monkey.Value == nil && monkeys[monkey.Operation[0]].Value != nil && monkeys[monkey.Operation[2]].Value != nil {
		value1 := *monkeys[monkey.Operation[0]].Value
		value2 := *monkeys[monkey.Operation[2]].Value
		value := calculateValue(value1, value2, monkey.Operation[1])
		monkey.Value = &value
		monkeys[name] = monkey
	}
}

func calculateValue(value1, value2 int, operation string) int {
	if operation == "+" {
		return value1 + value2
	} else if operation == "-" {
		return value1 - value2
	} else if operation == "*" {
		return value1 * value2
	} else {
		return value1 / value2
	}
}

func partTwo(monkeys map[string]Monkey) (int, error) {
	monkeys["humn"] = Monkey{Value: nil, Operation: []string{"humn"}}
	monkeysNotCompleted := map[string]bool{}
	for name := range monkeys {
		monkeysNotCompleted[name] = true
	}
	delete(monkeysNotCompleted, "humn")
	for monkeysNotCompleted["root"] {
		for name := range monkeysNotCompleted {
			updateMonkeyOperation(name, monkeys, monkeysNotCompleted)
		}
	}
	if value, err := strconv.Atoi(monkeys["root"].Operation[0]); err == nil {
		return reverseEquation(strings.Split(monkeys["root"].Operation[2], " "), value)
	} else {
		value, err := strconv.Atoi(monkeys["root"].Operation[2])
		if err != nil {
			return 0, err
		}

		return reverseEquation(strings.Split(monkeys["root"].Operation[0], " "), value)
	}
}

func updateMonkeyOperation(name string, monkeys map[string]Monkey, monkeysNotCompleted map[string]bool) {
	monkey := monkeys[name]
	if monkey.Value != nil {
		delete(monkeysNotCompleted, name)
		return
	}
	if !monkeysNotCompleted[monkey.Operation[0]] && !monkeysNotCompleted[monkey.Operation[2]] {
		if monkeys[monkey.Operation[0]].Value != nil && monkeys[monkey.Operation[2]].Value != nil {
			value := calculateValue(*monkeys[monkey.Operation[0]].Value, *monkeys[monkey.Operation[2]].Value, monkey.Operation[1])
			monkey.Value = &value
		} else if monkeys[monkey.Operation[0]].Value != nil {
			monkey.Operation[0] = strconv.Itoa(*monkeys[monkey.Operation[0]].Value)
			monkey.Operation[2] = "( " + strings.Join(monkeys[monkey.Operation[2]].Operation, " ") + " )"
		} else if monkeys[monkey.Operation[2]].Value != nil {
			monkey.Operation[0] = "( " + strings.Join(monkeys[monkey.Operation[0]].Operation, " ") + " )"
			monkey.Operation[2] = strconv.Itoa(*monkeys[monkey.Operation[2]].Value)
		} else {
			monkey.Operation[0] = "( " + strings.Join(monkeys[monkey.Operation[0]].Operation, " ") + " )"
			monkey.Operation[2] = "( " + strings.Join(monkeys[monkey.Operation[2]].Operation, " ") + " )"
		}
		monkeys[name] = monkey
		delete(monkeysNotCompleted, name)
	}
}

func reverseEquation(operation []string, value int) (int, error) {
	if operation[0] == "(" && operation[len(operation)-1] == ")" {
		operation = operation[1 : len(operation)-1]
	}
	if len(operation) == 1 {
		return value, nil
	}
	if operation[0] == "(" {
		value2, err := strconv.Atoi(operation[len(operation)-1])
		if err != nil {
			return 0, err
		}

		op := operationMap[operation[len(operation)-2]]
		value = calculateValue(value, value2, op)
		return reverseEquation(operation[:len(operation)-2], value)
	} else {
		value2, err := strconv.Atoi(operation[0])
		if err != nil {
			return 0, err
		}

		if operation[1] == "+" {
			value = calculateValue(value, value2, "-")
			return reverseEquation(operation[2:], value)
		} else if operation[1] == "-" {
			value = calculateValue(value2, value, "-")
			return reverseEquation(operation[2:], value)
		} else if operation[1] == "*" {
			value = calculateValue(value, value2, "/")
			return reverseEquation(operation[2:], value)
		} else {
			value = calculateValue(value2, value, "/")
			return reverseEquation(operation[2:], value)
		}
	}
}

func readMonkeys(fileName string) (map[string]Monkey, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	monkeys := map[string]Monkey{}
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), " ")
		if len(splitString) == 2 {
			value, err := strconv.Atoi(splitString[1])
			if err != nil {
				return nil, err
			}

			monkeys[splitString[0][:len(splitString[0])-1]] = Monkey{Value: &value}
		} else {
			monkeys[splitString[0][:len(splitString[0])-1]] = Monkey{Operation: splitString[1:]}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return monkeys, nil
}
