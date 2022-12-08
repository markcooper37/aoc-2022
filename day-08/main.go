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
	trees, err := readTrees("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(trees))
	fmt.Println(partTwo(trees))
}

func partOne(trees [][]int) int {
	visibleTrees := map[[2]int]bool{}
	for rowIndex, row := range trees {
		tallestTreeRight, tallestTreeLeft := -1, -1
		for columnIndex, tree := range row {
			if tree > tallestTreeRight {
				visibleTrees[[2]int{rowIndex, columnIndex}] = true
				tallestTreeRight = tree
			}
			if row[len(row)-columnIndex-1] > tallestTreeLeft {
				visibleTrees[[2]int{rowIndex, len(row) - columnIndex - 1}] = true
				tallestTreeLeft = row[len(row)-columnIndex-1]
			}
		}
	}
	for i := 0; i < len(trees[0]); i++ {
		tallestTreeDown, tallestTreeUp := -1, -1
		for j := 0; j < len(trees); j++ {
			if trees[j][i] > tallestTreeDown {
				visibleTrees[[2]int{j, i}] = true
				tallestTreeDown = trees[j][i]
			}
			if trees[len(trees)-j-1][i] > tallestTreeUp {
				visibleTrees[[2]int{len(trees) - j - 1, i}] = true
				tallestTreeUp = trees[len(trees)-j-1][i]
			}
		}
	}
	return len(visibleTrees)
}

func partTwo(trees [][]int) int {
	maxScore := 0
	for rowIndex, row := range trees {
		for columnIndex := range row {
			score := scenicScore(rowIndex, columnIndex, trees)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func scenicScore(rowIndex, columnIndex int, trees [][]int) int {
	treeHeight := trees[rowIndex][columnIndex]
	rowLeft, rowRight, columnUp, columnDown := 0, 0, 0, 0
	for i := rowIndex - 1; i >= 0; i-- {
		rowLeft++
		if trees[i][columnIndex] >= treeHeight {
			break
		}
	}
	for i := rowIndex + 1; i < len(trees[rowIndex]); i++ {
		rowRight++
		if trees[i][columnIndex] >= treeHeight {
			break
		}
	}
	for i := columnIndex - 1; i >= 0; i-- {
		columnUp++
		if trees[rowIndex][i] >= treeHeight {
			break
		}
	}
	for i := columnIndex + 1; i < len(trees); i++ {
		columnDown++
		if trees[rowIndex][i] >= treeHeight {
			break
		}
	}
	return rowLeft * rowRight * columnUp * columnDown
}

func readTrees(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	treeGrid := [][]int{}
	for scanner.Scan() {
		row := []int{}
		treeStrings := strings.Split(scanner.Text(), "")
		for _, tree := range treeStrings {
			height, err := strconv.Atoi(tree)
			if err != nil {
				return nil, err
			}

			row = append(row, height)
		}
		treeGrid = append(treeGrid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return treeGrid, nil
}
