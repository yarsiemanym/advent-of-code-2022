package day11

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n\n", parseMonkey)

	var part1Monkeys []*monkey
	for _, result := range results {
		part1Monkeys = append(part1Monkeys, result.(*monkey))
	}

	results = common.ParseFile(puzzle.InputFile, "\n\n", parseMonkey)

	var part2Monkeys []*monkey
	for _, result := range results {
		part2Monkeys = append(part2Monkeys, result.(*monkey))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(part1Monkeys),
		Part2: solvePart2(part2Monkeys),
	}
}

func solvePart1(monkeys []*monkey) string {
	log.Debug("Solving part 1.")

	monkeyBusiness := runSimulation(monkeys, 3, 20)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(monkeyBusiness)
}

func solvePart2(monkeys []*monkey) string {
	log.Debug("Solving part 2.")

	monkeyBusiness := runSimulation(monkeys, 1, 10000)

	log.Debug("Part 2 solved.")
	return strconv.Itoa(monkeyBusiness)
}

func runSimulation(monkeys []*monkey, reliefFactor int, maxRounds int) int {
	leastCommonMultiple := uint64(1)
	for _, monkey := range monkeys {
		leastCommonMultiple *= uint64(monkey.divisibleBy)
	}
	log.Debugf("Least common multiple of all monkey tests is %d.", leastCommonMultiple)

	for round := 1; round <= maxRounds; round++ {
		log.Debugf("Beginning round %d.", round)
		for _, monkey := range monkeys {
			monkey.TakeTurn(monkeys, reliefFactor, leastCommonMultiple)
		}
	}

	first, second := findTwoMostActiveMonkeys(monkeys)
	monkeyBusiness := first.activity * second.activity

	return monkeyBusiness
}

func findTwoMostActiveMonkeys(monkeys []*monkey) (first *monkey, second *monkey) {
	for _, monkey := range monkeys {
		log.Debugf("Monkey %d activity was %d.", monkey.id, monkey.activity)

		if first == nil || monkey.activity > first.activity {
			second = first
			first = monkey
		} else if second == nil || monkey.activity > second.activity {
			second = monkey
		}
	}

	log.Debugf("Monkeys %d and %d were the most active with %d and %d inspections.", first.id, second.id, first.activity, second.activity)
	return
}
