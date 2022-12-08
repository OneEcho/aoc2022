package main

import (
	"fmt"
	"os"
	"strings"
)

func findRepeatingStartIndex(dataStream string, distinctChars int) int {
	for idx := range dataStream {
		packet := dataStream[idx : idx+distinctChars]

		letterCount := make(map[string]int)
		for _, char := range packet {
			letterCount[string(char)]++
		}
		if len(letterCount) == distinctChars {
			return idx + distinctChars
		}
	}

	return 0
}

func main() {
	FilePath := "puzzleInput.txt"

	contents, err := os.ReadFile(FilePath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\r\n")

	dataStream := lines[0]

	partOneAns := findRepeatingStartIndex(dataStream, 4)
	fmt.Println("Part 1:", partOneAns)

	partTwoAns := findRepeatingStartIndex(dataStream, 14)
	fmt.Println("Part 2:", partTwoAns)
}
