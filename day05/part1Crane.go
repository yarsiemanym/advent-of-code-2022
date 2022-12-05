package day05

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type part1Crane struct {
	Stacks       []*common.Stack
	Instructions []*instruction
}

func (thisCrane *part1Crane) ExecuteInstructions() {
	log.Debug("Executing instructions.")

	for _, instruction := range thisCrane.Instructions {
		log.Debugf("Executing instruction: \"%s\"", instruction)

		for quantity := 0; quantity < instruction.Quantity; quantity++ {
			from := thisCrane.Stacks[instruction.From-1]
			to := thisCrane.Stacks[instruction.To-1]
			to.Push(from.Pop())
		}
	}
}

func (thisCrane *part1Crane) PeekTopCrates() []rune {
	topCrates := []rune{}

	for _, stack := range thisCrane.Stacks {
		topCrates = append(topCrates, stack.Peek().(rune))
	}

	return topCrates
}

func parsePart1Crane(text string) *part1Crane {
	if text == "" {
		return nil
	}

	log.Debug("Parsing crane.")

	tokens := common.Split(text, "\n\n")

	return &part1Crane{
		Stacks:       parseStacks(tokens[0]),
		Instructions: parseInstrictions(tokens[1]),
	}

}

func parseStacks(text string) []*common.Stack {
	if text == "" {
		return nil
	}

	log.Debug("Parsing stacks.")

	lines := common.Split(text, "\n")
	stackNumbers := lines[len(lines)-1]
	lastStackNumber := string(stackNumbers[len(stackNumbers)-2])
	stackCount, err := strconv.Atoi(lastStackNumber)
	common.Check(err)

	stacks := make([]*common.Stack, stackCount)

	for stackIndex := 0; stackIndex < stackCount; stackIndex++ {
		stacks[stackIndex] = common.NewStack()
	}

	for lineIndex := len(lines) - 2; lineIndex >= 0; lineIndex-- {
		line := lines[lineIndex]

		for charIndex := 1; charIndex < len(line); charIndex += 4 {
			crate := rune(line[charIndex])

			if crate != ' ' {
				stackIndex := int(charIndex / 4)
				log.Debugf("Pushing crate '%c' onto stack %d.", crate, stackIndex+1)
				stacks[stackIndex].Push(crate)
			}
		}
	}

	return stacks
}

func parseInstrictions(text string) []*instruction {
	if text == "" {
		return nil
	}

	log.Debug("Parsing instructions.")

	lines := common.Split(text, "\n")
	instructions := []*instruction{}

	for _, line := range lines {
		instruction := parseInstruction(line)

		if instruction == nil {
			continue
		}

		instructions = append(instructions, parseInstruction(line))
	}

	return instructions
}
