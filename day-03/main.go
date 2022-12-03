package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	rucksacks, err := readRucksacks("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(rucksacks))
	fmt.Println(partTwo(rucksacks))
}

func partOne(rucksacks []string) int {
	incorrectItems := []rune{}
	for _, rucksack := range rucksacks {
		commonItem := findCommonItem(rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:])
		if commonItem != nil {
			incorrectItems = append(incorrectItems, *commonItem)
		}
	}
	return sumPriorities(incorrectItems)
}

func partTwo(rucksacks []string) int {
	commonItems := []rune{}
	for i := 0; i < len(rucksacks)/3; i++ {
		commonItem := findCommonItem(rucksacks[3*i], rucksacks[3*i+1], rucksacks[3*i+2])
		if commonItem != nil {
			commonItems = append(commonItems, *commonItem)
		}
	}
	return sumPriorities(commonItems)
}

func findCommonItem(itemCollections ...string) *rune {
	counts := map[rune]int{}
	for _, items := range itemCollections {
		itemsMap := createItemMap(items)
		for item := range itemsMap {
			counts[item]++
		}
	}
	for item, count := range counts {
		if count == len(itemCollections) {
			return &item
		}
	}
	return nil
}

func createItemMap(items string) map[rune]bool {
	itemMap := map[rune]bool{}
	for _, item := range items {
		itemMap[item] = true
	}
	return itemMap
}

func sumPriorities(items []rune) int {
	total := 0
	for _, item := range items {
		total += priority(item)
	}
	return total
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1)
	} else {
		return int(item - 'A' + 27)
	}
}

func readRucksacks(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	rucksacks := []string{}
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rucksacks, nil
}
