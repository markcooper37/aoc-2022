package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	nodes, requiredNode, err := readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(nodes, requiredNode))
	nodes, requiredNode, err = readFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partTwo(nodes, requiredNode))
}

func partOne(nodes []*Node, requiredNode *Node) int {
	mixNodes(nodes)
	return findSum(requiredNode)
}

func partTwo(nodes []*Node, requiredNode *Node) int {
	for _, node := range nodes {
		node.Value *= 811589153
	}
	for iters := 0; iters < 10; iters++ {
		mixNodes(nodes)
	}
	return findSum(requiredNode)
}

func mixNodes(nodes []*Node) {
	for _, node := range nodes {
		moveDistance := node.Value % (len(nodes) - 1)
		if moveDistance > 0 {
			newLeft := node.Right
			for i := 1; i < moveDistance; i++ {
				newLeft = newLeft.Right
			}
			moveNode(node, newLeft)
		} else if moveDistance < 0 {
			newLeft := node.Left
			for i := 0; i > moveDistance; i-- {
				newLeft = newLeft.Left
			}
			moveNode(node, newLeft)
		}
	}
}

func moveNode(node *Node, newLeft *Node) {
	node.Left.Right = node.Right
	node.Right.Left = node.Left
	node.Right = newLeft.Right
	node.Left = newLeft
	newLeft.Right = node
	node.Right.Left = node
}

func findSum(requiredNode *Node) int {
	sum := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			requiredNode = requiredNode.Right
		}
		sum += requiredNode.Value
	}
	return sum
}

func readFile(fileName string) ([]*Node, *Node, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	nodes := []*Node{}
	var requiredNode *Node
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, nil, err
		}

		newNode := &Node{Value: value}
		if newNode.Value == 0 {
			requiredNode = newNode
		}
		if len(nodes) > 0 {
			nodes[len(nodes)-1].Right = newNode
			newNode.Left = nodes[len(nodes)-1]
		}
		nodes = append(nodes, newNode)
	}
	nodes[0].Left = nodes[len(nodes)-1]
	nodes[len(nodes)-1].Right = nodes[0]
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return nodes, requiredNode, nil
}
