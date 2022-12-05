package day05

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type model int

const (
	crateMover9000 model = 9000
	crateMover9001 model = 9001
)

type crane struct {
	Stacks       []*common.Stack
	Instructions []*instruction
}

func (thisCrane *crane) ExecuteInstructions(thisModel model) {
	log.Debug("Executing instructions.")

	for _, instruction := range thisCrane.Instructions {
		log.Debugf("Executing instruction: \"%s\"", instruction)

		switch thisModel {
		case crateMover9000:
			for quantity := 0; quantity < instruction.Quantity; quantity++ {
				thisCrane.MoveCrates(1, instruction.From, instruction.To)
			}
		case crateMover9001:
			thisCrane.MoveCrates(instruction.Quantity, instruction.From, instruction.To)
		default:
			log.Panicf("Unsupported model: %d", thisModel)
		}
	}
}

func (thisCrane *crane) MoveCrates(quantity int, from int, to int) {
	log.Debugf("Moving %d crates from stack %d to stack %d.", quantity, from, to)

	fromStack := thisCrane.Stacks[from-1]
	toStack := thisCrane.Stacks[to-1]
	crates := common.NewStack()

	for i := 0; i < quantity; i++ {
		crates.Push(fromStack.Pop())
	}

	for i := 0; i < quantity; i++ {
		toStack.Push(crates.Pop())
	}
}

func (thisCrane *crane) PeekTopCrates() []rune {
	topCrates := []rune{}

	for _, stack := range thisCrane.Stacks {
		topCrates = append(topCrates, stack.Peek().(rune))
	}

	return topCrates
}

func parseCrane(text string) *crane {
	if text == "" {
		return nil
	}

	log.Debug("Parsing crane.")

	tokens := common.Split(text, "\n\n")

	return &crane{
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
