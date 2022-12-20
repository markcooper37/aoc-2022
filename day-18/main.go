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
	cubes, err := readCubes("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(cubes))
	fmt.Println(partTwo(cubes))
}

func partOne(cubes map[[3]int]bool) int {
	count := 0
	for cube := range cubes {
		adjacentCubes := getAdjacentToCube(cube)
		for _, adjacentCube := range adjacentCubes {
			if !cubes[adjacentCube] {
				count++
			}
		}
	}
	return count
}

func partTwo(cubes map[[3]int]bool) int {
	min, max := findMinCoordinates(cubes), findMaxCoordinates(cubes)
	trappedCubes := map[[3]int]bool{}
	for adjacentCube := range getAllAdjacentCubes(cubes) {
		if !cubes[adjacentCube] && isTrapped(adjacentCube, max, min, cubes) {
			trappedCubes[adjacentCube] = true
		}
	}
	count := 0
	for cube := range cubes {
		adjacentCubes := getAdjacentToCube(cube)
		for _, adjacentCube := range adjacentCubes {
			if !cubes[adjacentCube] && !trappedCubes[adjacentCube] {
				count++
			}
		}
	}
	return count
}

func findMinCoordinates(cubes map[[3]int]bool) [3]int {
	min := [3]int{}
	for cube := range cubes {
		min = cube
		break
	}
	for cube := range cubes {
		if cube[0] < min[0] {
			min[0] = cube[0]
		}
		if cube[1] < min[1] {
			min[1] = cube[1]
		}
		if cube[2] < min[2] {
			min[2] = cube[2]
		}
	}
	return min
}

func findMaxCoordinates(cubes map[[3]int]bool) [3]int {
	max := [3]int{}
	for cube := range cubes {
		max = cube
		break
	}
	for cube := range cubes {
		if cube[0] > max[0] {
			max[0] = cube[0]
		}
		if cube[1] > max[1] {
			max[1] = cube[1]
		}
		if cube[2] > max[2] {
			max[2] = cube[2]
		}
	}
	return max
}

func isTrapped(cube, max, min [3]int, cubes map[[3]int]bool) bool {
	cubesToConsider := map[[3]int]bool{cube: true}
	cubesConsidered := map[[3]int]bool{}
	for {
		if len(cubesToConsider) == 0 {
			break
		}
		newCubesToConsider := map[[3]int]bool{}
		for cubeToConsider := range cubesToConsider {
			if cubeToConsider[0] < min[0] || cubeToConsider[1] < min[1] || cubeToConsider[2] < min[2] ||
				cubeToConsider[0] > max[0] || cubeToConsider[1] > max[1] || cubeToConsider[2] > max[2] {
				return false
			}
			adjacentCubes := getAdjacentToCube(cubeToConsider)
			for _, adjacentCube := range adjacentCubes {
				if !cubes[adjacentCube] && !cubesConsidered[adjacentCube] {
					newCubesToConsider[adjacentCube] = true
				}
			}
			cubesConsidered[cubeToConsider] = true
		}
		cubesToConsider = newCubesToConsider
	}
	return true
}

func getAllAdjacentCubes(cubes map[[3]int]bool) map[[3]int]bool {
	adjacentCubesMap := map[[3]int]bool{}
	for cube := range cubes {
		adjCubes := getAdjacentToCube(cube)
		for _, adjCube := range adjCubes {
			adjacentCubesMap[adjCube] = true
		}
	}
	return adjacentCubesMap
}

func getAdjacentToCube(cube [3]int) [][3]int {
	cubes := [][3]int{}
	cubes = append(cubes, [3]int{cube[0] + 1, cube[1], cube[2]})
	cubes = append(cubes, [3]int{cube[0] - 1, cube[1], cube[2]})
	cubes = append(cubes, [3]int{cube[0], cube[1] + 1, cube[2]})
	cubes = append(cubes, [3]int{cube[0], cube[1] - 1, cube[2]})
	cubes = append(cubes, [3]int{cube[0], cube[1], cube[2] + 1})
	cubes = append(cubes, [3]int{cube[0], cube[1], cube[2] - 1})
	return cubes
}

func readCubes(fileName string) (map[[3]int]bool, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	cubes := map[[3]int]bool{}
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), ",")
		cube := [3]int{}
		for i := 0; i <= 2; i++ {
			position, err := strconv.Atoi(splitString[i])
			if err != nil {
				return nil, err
			}

			cube[i] = position
		}
		cubes[cube] = true
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cubes, nil
}
