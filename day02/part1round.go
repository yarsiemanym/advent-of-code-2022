package day02

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part1Round struct {
	OpponentsChoice choice
	YourChoice      choice
}

// Determine the outcome of the round given your choice and the opponent's choice.
func (thisRound *part1Round) Outcome() outcome {
	log.Debug("Determinging outcome of round.")
	log.Debugf("opponentsChoice = %d", thisRound.OpponentsChoice)
	log.Debugf("yourChoice = %d", thisRound.YourChoice)

	if (thisRound.OpponentsChoice == rock && thisRound.YourChoice == paper) ||
		(thisRound.OpponentsChoice == paper && thisRound.YourChoice == scissors) ||
		(thisRound.OpponentsChoice == scissors && thisRound.YourChoice == rock) {
		log.Debugf("outcome = %d", win)
		return win
	} else if thisRound.OpponentsChoice == thisRound.YourChoice {
		log.Debugf("outcome = %d", draw)
		return draw
	} else {
		log.Debugf("outcome = %d", lose)
		return lose
	}
}

// Calculate the score for this round given your choice and the outcome of the round.
func (thisRound *part1Round) YourScore() int {
	return int(thisRound.YourChoice) + int(thisRound.Outcome())
}

// Parse a struct from the given line of text.
func parsePart1Round(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	var opponentsChoice choice
	var yourChoice choice

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
		yourChoice = rock
	case "Y":
		yourChoice = paper
	case "Z":
		yourChoice = scissors
	default:
		log.Panicf("invalid choice: %s", tokens[1])
	}

	return &part1Round{
		OpponentsChoice: opponentsChoice,
		YourChoice:      yourChoice,
	}
}
