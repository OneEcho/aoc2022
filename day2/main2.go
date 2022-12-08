package main

import (
	"bufio"
	"fmt"
	"os"
)

type round struct {
	opponentMove  string
	neededOutcome string
}

func newRound(opponentMove string, neededOutcome string) *round {
	return &round{opponentMove: opponentMove, neededOutcome: neededOutcome}
}

func createTournament(tourneyFile string) []round {
	file, err := os.Open(tourneyFile)
	if err != nil {
		panic(err)
	}

	var rounds []round

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		moves := scanner.Text()

		rounds = append(rounds, *newRound(string(moves[0]), string(moves[2])))
	}

	return rounds
}

func outcome(round round) int {
	switch round.opponentMove {
	case "A":
		switch round.neededOutcome {
		case "X":
			return 3
		case "Y":
			return 1
		case "Z":
			return 2
		}
	case "B":
		switch round.neededOutcome {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		}
	case "C":
		switch round.neededOutcome {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		}
	default:
		return 99
	}
	return 99
}

func main() {
	tourney := createTournament("./input.txt")
	fmt.Println(len(tourney))

	myMoveScoreMap := make(map[string]int)
	myMoveScoreMap["X"] = 0
	myMoveScoreMap["Y"] = 3
	myMoveScoreMap["Z"] = 6

	totalScore := 0
	for _, round := range tourney {
		totalScore += outcome(round) + myMoveScoreMap[round.neededOutcome]
	}

	fmt.Println(totalScore)
}
