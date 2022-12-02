package day02

import (
	"fmt"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part1Round struct {
	OpponentsChoice choice
	YourChoice      choice
}

func (thisRound *part1Round) Outcome() outcome {

	if (thisRound.OpponentsChoice == rock && thisRound.YourChoice == paper) ||
		(thisRound.OpponentsChoice == paper && thisRound.YourChoice == scissors) ||
		(thisRound.OpponentsChoice == scissors && thisRound.YourChoice == rock) {
		return win
	} else if thisRound.OpponentsChoice == thisRound.YourChoice {
		return draw
	} else {
		return lose
	}
}

func (thisRound *part1Round) YourScore() int {
	return int(thisRound.YourChoice) + int(thisRound.Outcome())
}

func parsePart1Round(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	var opponentsChoice choice
	var yourChoice choice
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
		yourChoice = rock
	case "Y":
		yourChoice = paper
	case "Z":
		yourChoice = scissors
	default:
		err = fmt.Errorf("invalid choice: %s", tokens[1])
	}
	common.Check(err)

	return &part1Round{
		OpponentsChoice: opponentsChoice,
		YourChoice:      yourChoice,
	}
}
