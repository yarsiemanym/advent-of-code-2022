package day05

import (
	"fmt"
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type instruction struct {
	Quantity int
	From     int
	To       int
}

func (thisInstruction *instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", thisInstruction.Quantity, thisInstruction.From, thisInstruction.To)
}

func parseInstruction(text string) *instruction {
	if text == "" {
		return nil
	}

	log.Debugf("Parsing instruction \"%s\".", text)

	expression := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	matches := expression.FindStringSubmatch(text)

	quantity, err := strconv.Atoi(matches[1])
	common.Check(err)
	log.Debugf("quantity = %d", quantity)

	from, err := strconv.Atoi(matches[2])
	common.Check(err)
	log.Debugf("from = %d", from)

	to, err := strconv.Atoi(matches[3])
	common.Check(err)
	log.Debugf("to = %d", to)

	return &instruction{
		Quantity: quantity,
		From:     from,
		To:       to,
	}
}
