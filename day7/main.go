package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	label string
	size  int32
}

func newFile(label string, size int32) *file {
	return &file{label, size}
}

type directory struct {
	parentDirectory  *directory
	label            string
	files            []file
	childDirectories []*directory
}

func newDirectory(parentDirectory *directory, label string) *directory {
	return &directory{parentDirectory, label, make([]file, 0), make([]*directory, 0)}
}

func (dir *directory) size() int64 {
	var size int64

	if len(dir.files) != 0 {
		for _, file := range dir.files {
			size += int64(file.size)
		}
	}

	for _, dir := range dir.childDirectories {
		size += dir.size()
	}

	return size
}

func createFilesystem(termOutput []string) *directory {
	rootDir := newDirectory(nil, "/")
	currentDir := rootDir

	findDir := func(fileSystem []*directory, label string) *directory {
		for _, directory := range fileSystem {
			if directory.label == label {
				return directory
			}
		}

		return nil
	}

	for _, line := range termOutput[1:] {
		if line[0] == '$' {
			command := strings.Split(line[2:], " ")
			operation := command[0]
			switch operation {
			case "cd":
				if command[1] == ".." {
					currentDir = currentDir.parentDirectory
				} else {
					currentDir = findDir(currentDir.childDirectories, command[1])
				}
			case "ls":
				continue
			}
		} else {
			data := strings.Split(line, " ")
			if data[0] == "dir" {
				directory := newDirectory(currentDir, data[1])
				currentDir.childDirectories = append(currentDir.childDirectories, directory)
			} else {
				fileSize, err := strconv.ParseInt(data[0], 10, 32)
				if err != nil {
					panic(err)
				}
				newFile := newFile(data[1], int32(fileSize))
				currentDir.files = append(currentDir.files, *newFile)
			}
		}
	}

	return rootDir
}

func getTotalSizeOfDirsLessThan10000(dir *directory) int64 {
	var total int64

	dirSize := dir.size()
	if dirSize <= 100000 {
		total += dirSize
	}

	for _, dir := range dir.childDirectories {
		total += getTotalSizeOfDirsLessThan10000(dir)
	}

	return total
}

func findPotentialDirs(dir *directory, results *[]int64, neededDirSize int64) {
	dirSize := dir.size()
	if dirSize > neededDirSize {
		*results = append(*results, dirSize)
	}

	for _, childDir := range dir.childDirectories {
		findPotentialDirs(childDir, results, neededDirSize)
	}
}

func main() {
	FilePath := "puzzleInput.txt"

	contents, err := os.ReadFile(FilePath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\r\n")

	fileSystem := createFilesystem(lines)

	partOneAnswer := getTotalSizeOfDirsLessThan10000(fileSystem)

	fmt.Println("Part 1:", partOneAnswer)

	neededDirSize := 30000000 - (70000000 - fileSystem.size())

	partTwoAnswer := make([]int64, 0)
	findPotentialDirs(fileSystem, &partTwoAnswer, neededDirSize)
	sort.Slice(partTwoAnswer, func(i, j int) bool { return partTwoAnswer[i] < partTwoAnswer[j] })

	fmt.Println("Part 2:", partTwoAnswer[0])
}
