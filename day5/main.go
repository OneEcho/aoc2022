package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type procedure struct {
	numCrates  int
	startCrate int
	endCrate   int
}

func makeStacks(lines []string) ([]Stack, int) {
	stackData := make([]string, 0)
	var numberLine string
	var movesIdx int

	for idx, line := range lines {
		if strings.Contains(line, "1") {
			numberLine = line
			movesIdx = idx + 2
			break
		} else {
			stackData = append(stackData, line)
		}
	}

	// Using slice to represent stacks so stack 1 is stack 0, stack 2 is stack 1, an so on
	numStacks, err := strconv.Atoi(string(numberLine[len(numberLine)-2]))
	if err != nil {
		panic(err)
	}

	stacks := make([]Stack, numStacks)
	for i := len(stackData) - 1; i >= 0; i-- {
		for idx, char := range stackData[i] {
			if char >= 'A' && char <= 'Z' {
				stackNum, err := strconv.Atoi(string(numberLine[idx]))
				if err != nil {
					panic(err)
				}
				stacks[stackNum-1].Push(string(char))
			}
		}
	}

	return stacks, movesIdx
}

func makeRearrangementProcedure(moves []string) []procedure {
	var result []procedure

	for _, move := range moves {
		moveData := strings.Split(move, " ")
		numCrates, err := strconv.Atoi(moveData[1])
		if err != nil {
			panic(err)
		}
		startCrate, err := strconv.Atoi(moveData[3])
		if err != nil {
			panic(err)
		}
		endCrate, err := strconv.Atoi(moveData[5])
		if err != nil {
			panic(err)
		}

		result = append(result, procedure{numCrates, startCrate - 1, endCrate - 1})
	}

	return result
}

func processMoves(stacks []Stack, rearragementProcedure []procedure) []string {
	for _, procedure := range rearragementProcedure {
		for i := 0; i < procedure.numCrates; i++ {
			crate, _ := stacks[procedure.startCrate].Pop()
			stacks[procedure.endCrate].Push(crate)
		}
	}

	var result []string
	for _, stack := range stacks {
		topCrate, _ := stack.Pop()
		result = append(result, topCrate)
	}

	return result
}

func processMovesPartTwo(stacks []Stack, rearragementProcedure []procedure) []string {
	var crates []string

	for _, procedure := range rearragementProcedure {
		for i := 0; i < procedure.numCrates; i++ {
			crate, _ := stacks[procedure.startCrate].Pop()
			crates = append(crates, crate)
		}

		for j := len(crates) - 1; j >= 0; j-- {
			stacks[procedure.endCrate].Push(crates[j])
		}
		crates = nil
	}

	var result []string
	for _, stack := range stacks {
		topCrate, _ := stack.Pop()
		result = append(result, topCrate)
	}

	return result
}

func main() {
	FilePath := "puzzleInput.txt"

	contents, err := os.ReadFile(FilePath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\r\n")

	stacksOne, movesIdx := makeStacks(lines)
	stacksTwo, movesIdx := makeStacks(lines)

	rearragementProcedure := makeRearrangementProcedure(lines[movesIdx:])

	topCratesP1 := processMoves(stacksOne, rearragementProcedure)
	topCratesP2 := processMovesPartTwo(stacksTwo, rearragementProcedure)

	fmt.Println("Part 1:")
	fmt.Println(strings.Join(topCratesP1, ""))

	fmt.Println("Part 2:")
	fmt.Println(strings.Join(topCratesP2, ""))
}
