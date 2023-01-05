package day20

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

const dencryptionKey = 811589153

func Solve(puzzle *common.Puzzle) common.Answer {

	lines := common.Split(common.ReadFile(puzzle.InputFile), "\n")
	numbers1 := []int{}
	for _, line := range lines {
		if line != "" {
			value, err := strconv.Atoi(line)
			common.Check(err)
			numbers1 = append(numbers1, value)
		}
	}

	numbers2 := []int{}
	for _, line := range lines {
		if line != "" {
			value, err := strconv.Atoi(line)
			common.Check(err)
			value *= dencryptionKey
			numbers2 = append(numbers2, value)
		}
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(numbers1),
		Part2: solvePart2(numbers2),
	}
}

func solvePart1(numbers []int) string {
	log.Debug("Solving part 1.")

	numbers = mix(numbers)
	coordinates := getGroveCoordinates(numbers)
	sum := coordinates.X() + coordinates.Y() + coordinates.Z()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(sum)
}

func solvePart2(numbers []int) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}

func getGroveCoordinates(numbers []int) *common.Point {
	homePosition := 0

	for index, number := range numbers {
		if number == 0 {
			homePosition = index
			break
		}
	}

	xPosition := (homePosition + 1000) % len(numbers)
	x := numbers[xPosition]

	yPosition := (homePosition + 2000) % len(numbers)
	y := numbers[yPosition]

	zPosition := (homePosition + 3000) % len(numbers)
	z := numbers[zPosition]

	return common.New3DPoint(x, y, z)
}

func mix(numbers []int) []int {

	mixableNumbers := make([]mixableNumber, len(numbers))

	for index, thisNumber := range numbers {
		mixableNumbers[index] = mixableNumber{
			value:   thisNumber,
			isMixed: false,
		}
	}
	for here := 0; here < len(mixableNumbers); {
		if mixableNumbers[here].isMixed {
			here++
			continue
		}

		thisNumber := mixableNumbers[here]
		thisNumber.isMixed = true
		mixableNumbers = append(mixableNumbers[0:here], mixableNumbers[here+1:]...)
		newPosition := (here + thisNumber.value) % len(mixableNumbers)

		if newPosition < 0 {
			newPosition = len(mixableNumbers) + newPosition
		}

		head := make([]mixableNumber, newPosition)
		tail := make([]mixableNumber, len(mixableNumbers)-newPosition)
		copy(head, mixableNumbers[:newPosition])
		copy(tail, mixableNumbers[newPosition:])
		mixableNumbers = append(head, thisNumber)
		mixableNumbers = append(mixableNumbers, tail...)

		if newPosition < here {
			here++
		}
	}

	for index, thisNumber := range mixableNumbers {
		numbers[index] = thisNumber.value
	}

	return numbers
}
