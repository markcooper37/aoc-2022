package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var rocks = [][][]byte{
	{
		{'#', '#', '#', '#'},
	},
	{
		{'.', '#', '.'},
		{'#', '#', '#'},
		{'.', '#', '.'},
	},
	{
		{'.', '.', '#'},
		{'.', '.', '#'},
		{'#', '#', '#'},
	},
	{
		{'#'},
		{'#'},
		{'#'},
		{'#'},
	},
	{
		{'#', '#'},
		{'#', '#'},
	},
}

type Phase struct {
	Positions     map[[2]int]bool
	CurrentRock   int
	Top           int
	RockBottom    int
	LeftMostPoint int
	RocksDropped  int
}

func main() {
	jetPattern, err := readJetPattern("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(jetPattern))
	fmt.Println(partTwo(jetPattern))
}

func partOne(jetPattern string) int {
	return findHeight(jetPattern, 2022)
}

func partTwo(jetPattern string) int {
	return findHeight(jetPattern, 1000000000000)
}

func findHeight(jetPattern string, iterations int) int {
	jetPosition, top, shift := 0, 0, 0
	phases := []Phase{}
	foundEqualPhases := false
	existingRocks := map[[2]int]bool{{0, 0}: true, {1, 0}: true, {2, 0}: true, {3, 0}: true, {4, 0}: true, {5, 0}: true, {6, 0}: true}
	phaseRocks := map[[2]int]bool{}
	for i := 0; i < iterations; i++ {
		rock := rocks[i%5]
		rockBottom := top + 4
		leftMostPoint := 2
		for {
			if !foundEqualPhases {
				if jetPosition == 0 {
					phases = append(phases, findPhase(leftMostPoint, rockBottom, top, i, i%5, phaseRocks))
					phaseRocks = map[[2]int]bool{}
				CheckPhases:
					for j := len(phases) - 1; j > 0; j-- {
						for k := j - 1; k >= 0; k-- {
							if equalPhases(phases[j], phases[k]) {
								remaining := (iterations - i) / (phases[j].RocksDropped - phases[k].RocksDropped)
								shift = remaining * (phases[j].Top - phases[k].Top)
								i += remaining * (phases[j].RocksDropped - phases[k].RocksDropped)
								foundEqualPhases = true
								break CheckPhases
							}
						}
					}
				}
			}
			if jetPattern[jetPosition] == '<' && canMoveLeft(leftMostPoint, rockBottom, rock, existingRocks) {
				leftMostPoint--
			} else if jetPattern[jetPosition] == '>' && canMoveRight(leftMostPoint, rockBottom, rock, existingRocks) {
				leftMostPoint++
			}
			jetPosition = (jetPosition + 1) % len(jetPattern)
			if canMoveDown(leftMostPoint, rockBottom, rock, existingRocks) {
				rockBottom--
			} else {
				existingRocks, top = addToRockMap(leftMostPoint, rockBottom, top, rock, existingRocks)
				phaseRocks, _ = addToRockMap(leftMostPoint, rockBottom, top, rock, phaseRocks)
				break
			}
		}
	}
	return shift + top
}

func canMoveLeft(leftMostPoint, rockBottom int, rock [][]byte, existingRocks map[[2]int]bool) bool {
	if leftMostPoint == 0 {
		return false
	}
	for rowIndex, row := range rock {
		for columnIndex, value := range row {
			if value == '#' {
				xPos := columnIndex + leftMostPoint
				yPos := rockBottom + (len(rock) - rowIndex - 1)
				if existingRocks[[2]int{xPos - 1, yPos}] {
					return false
				}
			}
		}
	}
	return true
}

func canMoveRight(leftMostPoint, rockBottom int, rock [][]byte, existingRocks map[[2]int]bool) bool {
	if leftMostPoint+len(rock[0])-1 == 6 {
		return false
	}
	for rowIndex, row := range rock {
		for columnIndex, value := range row {
			if value == '#' {
				xPos := columnIndex + leftMostPoint
				yPos := rockBottom + (len(rock) - rowIndex - 1)
				if existingRocks[[2]int{xPos + 1, yPos}] {
					return false
				}
			}
		}
	}
	return true
}

func canMoveDown(leftMostPoint, rockBottom int, rock [][]byte, existingRocks map[[2]int]bool) bool {
	for rowIndex, row := range rock {
		for columnIndex, value := range row {
			if value == '#' {
				xPos := columnIndex + leftMostPoint
				yPos := rockBottom + (len(rock) - rowIndex - 1)
				if existingRocks[[2]int{xPos, yPos - 1}] {
					return false
				}
			}
		}
	}
	return true
}

func addToRockMap(leftMostPoint, rockBottom, top int, rock [][]byte, existingRocks map[[2]int]bool) (map[[2]int]bool, int) {
	for rowIndex, row := range rock {
		for columnIndex, value := range row {
			if value == '#' {
				xPos := columnIndex + leftMostPoint
				yPos := rockBottom + (len(rock) - rowIndex - 1)
				existingRocks[[2]int{xPos, yPos}] = true
				if yPos > top {
					top = yPos
				}
			}
		}
	}
	return existingRocks, top
}

func equalPhases(phase1, phase2 Phase) bool {
	if phase1.LeftMostPoint != phase2.LeftMostPoint || phase1.RockBottom != phase2.RockBottom ||
		phase1.CurrentRock != phase2.CurrentRock || len(phase1.Positions) != len(phase2.Positions) {
		return false
	}
	for key := range phase1.Positions {
		if !phase2.Positions[key] {
			return false
		}
	}
	return true
}

func findPhase(leftMostPoint, rockBottom, top, rocksDropped, rock int, existingRocks map[[2]int]bool) Phase {
	phase := Phase{CurrentRock: rock, Top: top, LeftMostPoint: leftMostPoint, RocksDropped: rocksDropped, Positions: map[[2]int]bool{}}
	lowestPoint := -1
	for position := range existingRocks {
		if lowestPoint == -1 || position[1] < lowestPoint {
			lowestPoint = position[1]
		}
	}
	for position := range existingRocks {
		phase.Positions[[2]int{position[0], position[1] - lowestPoint}] = true
	}
	phase.RockBottom = rockBottom - lowestPoint
	return phase
}

func readJetPattern(fileName string) (string, error) {
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
