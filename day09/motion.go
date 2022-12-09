package day09

import (
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type motion struct {
	Direction rune
	Steps     int
}

func parseMotion(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, " ")

	direction := rune(tokens[0][0])

	steps, err := strconv.Atoi(tokens[1])
	common.Check(err)

	return &motion{
		Direction: direction,
		Steps:     steps,
	}
}
