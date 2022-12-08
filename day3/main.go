package main

import (
	"fmt"
	"os"
	"strings"
)

func findSharedItem(compartmentOne, compartmentTwo string) rune {
	for _, item := range compartmentOne {
		if strings.Contains(compartmentTwo, string(item)) {
			return item
		}
	}
	return ' '
}

func itemTypeToPriority(itemType rune) int {
	if itemType >= 'a' {
		return int(itemType-'a') + 1
	}

	return int(itemType-'A') + 27
}

func findSharedBadge(ruckSackOne, ruckSackTwo, ruckSackThree string) rune {
	for _, item := range ruckSackOne {
		if strings.Contains(ruckSackTwo, string(item)) && strings.Contains(ruckSackThree, string(item)) {
			return item
		}
	}
	return ' '
}

func main() {
	contents, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	rucksacks := strings.Split(string(contents), "\n")

	prioritySum := 0
	for _, rucksack := range rucksacks {
		sackLength := len(rucksack)
		compartmentOne := rucksack[:sackLength/2]
		compartmentTwo := rucksack[sackLength/2:]
		sharedItemType := findSharedItem(compartmentOne, compartmentTwo)
		prioritySum += itemTypeToPriority(sharedItemType)
	}
	fmt.Println("Part 1 Answer:", prioritySum)

	badgeSum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		groupRucksacks := rucksacks[i : i+3]
		badge := findSharedBadge(groupRucksacks[0], groupRucksacks[1], groupRucksacks[2])
		badgeSum += itemTypeToPriority(badge)
	}
	fmt.Println("Part 2 Answer:", badgeSum)
}
