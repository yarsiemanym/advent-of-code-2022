package day02

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part2Round struct {
	OpponentsChoice choice
	DesiredOutcome  outcome
}

// Determine the choice you must make given the opponent's choice and the desired outcome.
func (thisRound *part2Round) YourChoice() choice {
	log.Debug("Determinging your choice.")
	log.Debugf("opponentsChoice = %d", thisRound.OpponentsChoice)
	log.Debugf("desiredOutcome = %d", thisRound.DesiredOutcome)

	if (thisRound.OpponentsChoice == rock && thisRound.DesiredOutcome == draw) ||
		(thisRound.OpponentsChoice == paper && thisRound.DesiredOutcome == lose) ||
		(thisRound.OpponentsChoice == scissors && thisRound.DesiredOutcome == win) {
		log.Debugf("yourChoice = %d", rock)
		return rock
	} else if (thisRound.OpponentsChoice == rock && thisRound.DesiredOutcome == win) ||
		(thisRound.OpponentsChoice == paper && thisRound.DesiredOutcome == draw) ||
		(thisRound.OpponentsChoice == scissors && thisRound.DesiredOutcome == lose) {
		log.Debugf("yourChoice = %d", paper)
		return paper
	} else {
		log.Debugf("yourChoice = %d", scissors)
		return scissors
	}
}

// Calculate the score for this round given your choice and the outcome of the round.
func (thisRound *part2Round) YourScore() int {
	return int(thisRound.YourChoice()) + int(thisRound.DesiredOutcome)
}

// Parse a struct from the given line of text.
func parsePart2Round(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	var opponentsChoice choice
	var desiredOutcome outcome

	switch tokens[0] {
	case "A":
		opponentsChoice = rock
	case "B":
		opponentsChoice = paper
	case "C":
		opponentsChoice = scissors
	default:
		log.Panicf("invalid choice: %s", tokens[0])
	}

	switch tokens[1] {
	case "X":
		desiredOutcome = lose
	case "Y":
		desiredOutcome = draw
	case "Z":
		desiredOutcome = win
	default:
		log.Panicf("invalid choice: %s", tokens[1])
	}

	return &part2Round{
		OpponentsChoice: opponentsChoice,
		DesiredOutcome:  desiredOutcome,
	}
}
