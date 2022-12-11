package day11

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type operation struct {
	operand1 uint64
	operator rune
	operand2 uint64
}

func (thisOperation *operation) Execute(old uint64) uint64 {
	operand1, operand2, new := old, old, uint64(0)

	if thisOperation.operand1 != 0 {
		operand1 = thisOperation.operand1
	}

	if thisOperation.operand2 != 0 {
		operand1 = thisOperation.operand2
	}

	switch thisOperation.operator {
	case '+':
		new = operand1 + operand2
	case '*':
		new = operand1 * operand2
	default:
		log.Panicf("Unsupported operator '%c'.", thisOperation.operator)
	}

	return new
}

func parseOperation(text string) *operation {
	opRegexp := regexp.MustCompile(`new = (\S+) ([\*|+]) (\S+)`)
	matches := opRegexp.FindStringSubmatch(text)
	operand1, operator, operand2 := uint64(0), '+', uint64(0)

	if matches[1] != "old" {
		number, err := strconv.Atoi(matches[1])
		common.Check(err)
		operand1 = uint64(number)
	}

	switch matches[2] {
	case "+":
		operator = '+'
	case "*":
		operator = '*'
	default:
		log.Panicf("Unsupported operator \"%s\".", matches[2])
	}

	if matches[3] != "old" {
		number, err := strconv.Atoi(matches[3])
		common.Check(err)
		operand2 = uint64(number)
	}

	return &operation{
		operand1: operand1,
		operator: operator,
		operand2: operand2,
	}
}
