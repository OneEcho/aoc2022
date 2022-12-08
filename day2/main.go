package main

import (
	"bufio"
	"fmt"
	"os"
)

type round struct {
	opponentMove string
	myMove       string
}

func newRound(opponentMove string, myMove string) *round {
	return &round{opponentMove: opponentMove, myMove: myMove}
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
		switch round.myMove {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
	case "B":
		switch round.myMove {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
	case "C":
		switch round.myMove {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
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
	myMoveScoreMap["X"] = 1
	myMoveScoreMap["Y"] = 2
	myMoveScoreMap["Z"] = 3

	totalScore := 0
	for _, round := range tourney {
		totalScore += myMoveScoreMap[round.myMove] + outcome(round)
	}

	fmt.Println(totalScore)
}
