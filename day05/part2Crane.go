package day05

import (
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part2Crane struct {
	Stacks       []*common.Stack
	Instructions []*instruction
}

func (thisCrane *part2Crane) ExecuteInstructions() {
	log.Debug("Executing instructions.")

	for _, instruction := range thisCrane.Instructions {
		log.Debugf("Executing instruction: \"%s\"", instruction)

		from := thisCrane.Stacks[instruction.From-1]
		to := thisCrane.Stacks[instruction.To-1]

		crates := common.NewStack()

		for quantity := 0; quantity < instruction.Quantity; quantity++ {
			crates.Push(from.Pop())
		}

		for quantity := 0; quantity < instruction.Quantity; quantity++ {
			to.Push(crates.Pop())
		}
	}
}

func (thisCrane *part2Crane) PeekTopCrates() []rune {
	topCrates := []rune{}

	for _, stack := range thisCrane.Stacks {
		topCrates = append(topCrates, stack.Peek().(rune))
	}

	return topCrates
}

func parsePart2Crane(text string) *part2Crane {
	if text == "" {
		return nil
	}

	log.Debug("Parsing crane.")

	tokens := common.Split(text, "\n\n")

	return &part2Crane{
		Stacks:       parseStacks(tokens[0]),
		Instructions: parseInstrictions(tokens[1]),
	}

}
