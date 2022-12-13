package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Node struct {
	Children []*Node
	Value    *int
	Parent   *Node
}

func main() {
	packets, err := readPackets("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(packets))
	fmt.Println(partTwo(packets))
}

func partOne(packets []*Node) int {
	total := 0
	for i := 0; i < len(packets); i += 2 {
		if correctOrder, _ := rightOrder(packets[i], packets[i+1]); correctOrder {
			total += i/2 + 1
		}
	}
	return total
}

func rightOrder(packetOne, packetTwo *Node) (bool, bool) {
	for childIndex, childOne := range packetOne.Children {
		if childIndex >= len(packetTwo.Children) {
			return false, true
		}
		childTwo := packetTwo.Children[childIndex]
		if childOne.Value != nil && childTwo.Value != nil {
			if *childOne.Value < *childTwo.Value {
				return true, true
			} else if *childOne.Value > *childTwo.Value {
				return false, true
			}
		} else if childOne.Value == nil && childTwo.Value != nil {
			tempNode := &Node{}
			tempNode.Children = append(tempNode.Children, &Node{Value: childTwo.Value})
			correctOrder, complete := rightOrder(childOne, tempNode)
			if complete {
				return correctOrder, complete
			}
		} else if childOne.Value != nil && childTwo.Value == nil {
			tempNode := &Node{}
			tempNode.Children = append(tempNode.Children, &Node{Value: childOne.Value})
			correctOrder, complete := rightOrder(tempNode, childTwo)
			if complete {
				return correctOrder, complete
			}
		} else {
			correctOrder, complete := rightOrder(childOne, childTwo)
			if complete {
				return correctOrder, complete
			}
		}
	}
	if len(packetOne.Children) < len(packetTwo.Children) {
		return true, true
	}
	return true, false
}

func partTwo(packets []*Node) int {
	twoPacket, sixPacket := createPacket("[[2]]"), createPacket("[[6]]")
	packets = append(packets, []*Node{twoPacket, sixPacket}...)
	sort.Slice(packets, func(i, j int) bool {
		correctOrder, _ := rightOrder(packets[i], packets[j])
		return correctOrder
	})
	twoIndex, sixIndex := -1, -1
	for packetIndex, packet := range packets {
		if packet == twoPacket {
			twoIndex = packetIndex + 1
		} else if packet == sixPacket {
			sixIndex = packetIndex + 1
		}
	}
	return twoIndex * sixIndex
}

func createPacket(input string) *Node {
	root := &Node{}
	currentNode := root
	newPacket := root
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			newChild := &Node{Parent: currentNode}
			currentNode.Children = append(currentNode.Children, newChild)
			currentNode = newChild
		} else if input[i] == ']' {
			currentNode = currentNode.Parent
		} else if input[i] == ',' {
			continue
		} else {
			for j := i; j < len(input); j++ {
				if input[j] < '0' || input[j] > '9' {
					value, err := strconv.Atoi(input[i:j])
					if err != nil {
						log.Fatal(err)
					}

					currentNode.Children = append(currentNode.Children, &Node{Parent: currentNode, Value: &value})
					i = j - 1
					break
				}
			}
		}
	}
	return newPacket
}

func readPackets(fileName string) ([]*Node, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	packets := []*Node{}
	for scanner.Scan() {
		if scanner.Text() != "" {
			packets = append(packets, createPacket(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return packets, nil
}
