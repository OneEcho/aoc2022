package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	startSection int
	endSection   int
}

func newSectionData(pair string) *section {
	sections := strings.Split(pair, "-")
	startSection, err := strconv.Atoi(sections[0])

	if err != nil {
		panic(err)
	}

	endSection, err := strconv.Atoi(sections[1])

	if err != nil {
		panic(err)
	}

	return &section{startSection: startSection, endSection: endSection}
}

func findSmallerSection(sectionOne, sectionTwo section) (section, section) {
	var smallerSection section
	var biggerSection section

	if sectionOne.endSection-sectionOne.startSection <= sectionTwo.endSection-sectionTwo.startSection {
		smallerSection = sectionOne
		biggerSection = sectionTwo
	} else {
		smallerSection = sectionTwo
		biggerSection = sectionOne
	}

	return smallerSection, biggerSection
}

func contains(smallerSection, biggerSection section) bool {
	if smallerSection.startSection >= biggerSection.startSection && smallerSection.endSection <= biggerSection.endSection {
		return true
	} else {
		return false
	}
}

func overlap(smallerSection, biggerSection section) bool {
	if smallerSection.endSection >= biggerSection.startSection && smallerSection.startSection <= biggerSection.endSection {
		return true
	} else {
		return false
	}
}

func main() {
	filepath := "puzzleInput.txt"

	contents, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\r\n")

	sectionInSectionCount := 0
	sectionsOverlap := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		sectionOne := newSectionData(pairs[0])
		sectionTwo := newSectionData(pairs[1])

		smallerSection, biggerSection := findSmallerSection(*sectionOne, *sectionTwo)

		if contains(smallerSection, biggerSection) {
			sectionInSectionCount++
		}

		if overlap(smallerSection, biggerSection) {
			sectionsOverlap++
		}
	}

	fmt.Println("Part 1 answer: ", sectionInSectionCount)
	fmt.Println("Part 2 answer: ", sectionsOverlap)
}
