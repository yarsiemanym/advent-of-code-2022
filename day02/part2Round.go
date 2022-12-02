package day02

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part2Round struct {
	OpponentsChoice choice
	DesiredOutcome  outcome
}

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

func (thisRound *part2Round) YourScore() int {
	return int(thisRound.YourChoice()) + int(thisRound.DesiredOutcome)
}

func parsePart2Round(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	var opponentsChoice choice
	var desiredOutcome outcome
	var err error

	switch tokens[0] {
	case "A":
		opponentsChoice = rock
	case "B":
		opponentsChoice = paper
	case "C":
		opponentsChoice = scissors
	default:
		err = fmt.Errorf("invalid choice: %s", tokens[0])
	}
	common.Check(err)

	switch tokens[1] {
	case "X":
		desiredOutcome = lose
	case "Y":
		desiredOutcome = draw
	case "Z":
		desiredOutcome = win
	default:
		err = fmt.Errorf("invalid choice: %s", tokens[1])
	}
	common.Check(err)

	return &part2Round{
		OpponentsChoice: opponentsChoice,
		DesiredOutcome:  desiredOutcome,
	}
}
