package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Valve struct {
	Rate   int
	Valves []string
}

func main() {
	valves, err := readValves("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(valves))
	fmt.Println(partTwo(valves))
}

func partOne(valves map[string]Valve) int {
	valveDistances := findShortestDistancesBetweenValves(valves)
	validValves := []string{}
	for name, valve := range valves {
		if valve.Rate > 0 {
			validValves = append(validValves, name)
		}
	}
	allSequences := findSequences([]string{}, validValves, valveDistances, 30)
	maxFlow := 0
	for _, sequence := range allSequences {
		flow := calculateFlow(30, "AA", sequence, valves, valveDistances)
		if flow > maxFlow {
			maxFlow = flow
		}
	}
	return maxFlow
}

func findShortestDistancesBetweenValves(allValves map[string]Valve) map[string]map[string]int {
	minimumDistances := map[string]map[string]int{}
	for firstValve := range allValves {
		if allValves[firstValve].Rate == 0 && firstValve != "AA" {
			continue
		}
		minimumDistances[firstValve] = map[string]int{}
		for secondValve := range allValves {
			if firstValve == secondValve {
				continue
			}
			if allValves[secondValve].Rate == 0 {
				continue
			}
			minimumDistances[firstValve][secondValve] = findDistanceBetweenTwoValves(firstValve, secondValve, allValves)
		}
	}
	return minimumDistances
}

func findDistanceBetweenTwoValves(valveOne, valveTwo string, valves map[string]Valve) int {
	if valveOne == valveTwo {
		return 0
	}
	valvesFound := map[string]bool{}
	for _, valve := range valves[valveOne].Valves {
		if valve == valveTwo {
			return 2
		}
		valvesFound[valve] = true
	}
	for k := 3; k > 0; k++ {
		newValvesFound := map[string]bool{}
		for valve := range valvesFound {
			for _, otherValve := range valves[valve].Valves {
				if otherValve == valveTwo {
					return k
				}
				newValvesFound[otherValve] = true
			}
		}
		valvesFound = newValvesFound
	}
	return -1
}

func findSequences(existingSequence []string, values []string, valveDistances map[string]map[string]int, distance int) [][]string {
	if len(values) == 0 {
		return [][]string{existingSequence}
	}
	sequences := [][]string{}
	for index, value := range values {
		if len(existingSequence) == 0 {
			newDistance := distance - valveDistances["AA"][value]
			newSequence := []string{}
			newSequence = append(newSequence, existingSequence...)
			newSequence = append(newSequence, value)
			newValues := []string{}
			newValues = append(newValues, values...)
			newValues = append(newValues[:index], newValues[index+1:]...)
			sequences = append(sequences, findSequences(newSequence, newValues, valveDistances, newDistance)...)
		} else if valveDistances[existingSequence[len(existingSequence)-1]][value] < distance {
			newDistance := distance - valveDistances[existingSequence[len(existingSequence)-1]][value]
			newSequence := []string{}
			newSequence = append(newSequence, existingSequence...)
			newSequence = append(newSequence, value)
			newValues := []string{}
			newValues = append(newValues, values...)
			newValues = append(newValues[:index], newValues[index+1:]...)
			sequences = append(sequences, findSequences(newSequence, newValues, valveDistances, newDistance)...)
		}
	}
	if len(sequences) == 0 {
		return [][]string{existingSequence}
	}
	return sequences
}

func calculateFlow(time int, current string, sequence []string, valves map[string]Valve, valveDistances map[string]map[string]int) int {
	flow := 0
	for _, name1 := range sequence {
		distance := valveDistances[current][name1]
		time = time - distance
		if time < 0 {
			break
		}
		flow += time * valves[name1].Rate
		current = name1
	}
	return flow
}

func partTwo(valves map[string]Valve) int {
	valveDistances := findShortestDistancesBetweenValves(valves)
	maxFlow := 0
	for name1, distance1 := range valveDistances["AA"] {
		for name2, distance2 := range valveDistances["AA"] {
			if name1 == name2 {
				continue
			}
			flowsTurnedOn := map[string]bool{}
			flow := findMaxFlow(flowsTurnedOn, name1, name2, distance1, distance2, 0, 26, maxFlow, valveDistances, valves)
			if flow > maxFlow {
				maxFlow = flow
			}
		}
	}
	return maxFlow
}

func findMaxFlow(flowsTurnedOn map[string]bool, dest1, dest2 string, time1, time2, totalFlow, time, target int, valveDistances map[string]map[string]int, valves map[string]Valve) int {
	minimum := min([]int{time1, time2})
	time -= minimum
	if time <= 0 {
		return totalFlow
	}
	time1 -= minimum
	time2 -= minimum
	newTotalFlow := totalFlow
	newFlowsTurnedOn := map[string]bool{}
	for key, value := range flowsTurnedOn {
		newFlowsTurnedOn[key] = value
	}
	if time1 == 0 {
		if !newFlowsTurnedOn[dest1] {
			newTotalFlow += time * valves[dest1].Rate
			newFlowsTurnedOn[dest1] = true
		}
	}
	if time2 == 0 {
		if !newFlowsTurnedOn[dest2] {
			newTotalFlow += time * valves[dest2].Rate
			newFlowsTurnedOn[dest2] = true
		}
	}
	maxRemaining := 0
	for name := range valveDistances[dest1] {
		if !newFlowsTurnedOn[name] {
			maxRemaining += time*valves[name].Rate
		}
	}
	if !newFlowsTurnedOn[dest1] {
		maxRemaining += time*valves[dest1].Rate
	}
	if newTotalFlow + maxRemaining < target {
		return newTotalFlow
	}
	maxFlow := newTotalFlow
	if time1 == 0 && time2 == 0 {
		for name1, distance1 := range valveDistances[dest1] {
			if !newFlowsTurnedOn[name1] {
				for name2, distance2 := range valveDistances[dest2] {
					if !newFlowsTurnedOn[name2] {
						flow := findMaxFlow(newFlowsTurnedOn, name1, name2, distance1, distance2, newTotalFlow, time, target, valveDistances, valves)
						if flow > maxFlow {
							maxFlow = flow
						}
					}
				}
			}
		}
	} else if time1 == 0 && time2 > 0 {
		for name, distance := range valveDistances[dest1] {
			if !newFlowsTurnedOn[name] {
				flow := findMaxFlow(newFlowsTurnedOn, name, dest2, distance, time2, newTotalFlow, time, target, valveDistances, valves)
				if flow > maxFlow {
					maxFlow = flow
				}
			}
		}
	} else if time1 > 0 && time2 == 0 {
		for name, distance := range valveDistances[dest2] {
			if !newFlowsTurnedOn[name] {
				flow := findMaxFlow(newFlowsTurnedOn, dest1, name, time1, distance, newTotalFlow, time, target, valveDistances, valves)
				if flow > maxFlow {
					maxFlow = flow
				}
			}
		}
	} else {
		flow := findMaxFlow(newFlowsTurnedOn, dest1, dest2, time1, time2, newTotalFlow, time, target, valveDistances, valves)
		if flow > maxFlow {
			maxFlow = flow
		}
	}
	return maxFlow
}

func min(ints []int) int {
	minimum := ints[0]
	for _, value := range ints {
		if value < minimum {
			minimum = value
		}
	}
	return minimum
}

func readValves(fileName string) (map[string]Valve, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	allValves := map[string]Valve{}
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), " ")
		name := splitString[1]
		rateString := splitString[4][5 : len(splitString[4])-1]
		rate, err := strconv.Atoi(rateString)
		if err != nil {
			return nil, err
		}
		
		valves := []string{}
		for i := 9; i < len(splitString); i++ {
			valves = append(valves, strings.TrimSuffix(splitString[i], ","))
		}
		allValves[name] = Valve{Rate: rate, Valves: valves}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return allValves, nil
}
