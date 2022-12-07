package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type Directory struct {
	Parent   *Directory
	Children []*Directory
	Name     string
	Files    []File
}

func main() {
	output, err := readOutput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(output))
	fmt.Println(partTwo(output))
}

func partOne(output [][]string) int {
	_, allDirectories := constructFileTree(output)
	total := 0
	for _, directory := range allDirectories {
		directorySize := calculateDirectorySize(directory)
		if directorySize <= 100000 {
			total += directorySize
		}
	}
	return total
}

func partTwo(output [][]string) int {
	root, allDirectories := constructFileTree(output)
	requiredSpace := calculateDirectorySize(root) - 40000000
	sizeToDelete := 70000000
	for _, directory := range allDirectories {
		directorySize := calculateDirectorySize(directory)
		if directorySize >= requiredSpace && directorySize < sizeToDelete {
			sizeToDelete = directorySize
		}
	}
	return sizeToDelete
}

func constructFileTree(output [][]string) (*Directory, []*Directory) {
	root := &Directory{Name: "/"}
	allDirectories := []*Directory{root}
	currentDirectory := root
	for _, line := range output {
		if line[0] == "$" {
			if line[1] == "cd" {
				if line[2] == ".." {
					currentDirectory = currentDirectory.Parent
				} else if line[2] == "/" {
					currentDirectory = root
				} else {
					for _, child := range currentDirectory.Children {
						if child.Name == line[2] {
							currentDirectory = child
						}
					}

				}
			}
		} else if line[0] == "dir" {
			newChild := &Directory{Name: line[1]}
			allDirectories = append(allDirectories, newChild)
			newChild.Parent = currentDirectory
			currentDirectory.Children = append(currentDirectory.Children, newChild)
		} else {
			size, err := strconv.Atoi(line[0])
			if err != nil {
				continue
			}

			currentDirectory.Files = append(currentDirectory.Files, File{Name: line[1], Size: size})
		}
	}
	return root, allDirectories
}

func calculateDirectorySize(directory *Directory) int {
	size := 0
	for _, file := range directory.Files {
		size += file.Size
	}
	for _, childDirectory := range directory.Children {
		size += calculateDirectorySize(childDirectory)
	}
	return size
}

func readOutput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	output := [][]string{}
	for scanner.Scan() {
		output = append(output, strings.Split(scanner.Text(), " "))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output, nil
}
