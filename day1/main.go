package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	var totalCalorieCounts []int
	calorieCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			totalCalorieCounts = append(totalCalorieCounts, calorieCount)
			calorieCount = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}

			calorieCount += calories
		}
	}

	sort.Ints(totalCalorieCounts)
	fmt.Println(totalCalorieCounts[len(totalCalorieCounts)-1], totalCalorieCounts[len(totalCalorieCounts)-2], totalCalorieCounts[len(totalCalorieCounts)-3])
	fmt.Println(totalCalorieCounts[len(totalCalorieCounts)-1] + totalCalorieCounts[len(totalCalorieCounts)-2] + totalCalorieCounts[len(totalCalorieCounts)-3])
}
