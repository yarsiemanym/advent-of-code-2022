package day05

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	text := common.ReadFile(puzzle.InputFile)
	crane1 := parsePart1Crane(text)
	crane2 := parsePart2Crane(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(crane1),
		Part2: solvePart2(crane2),
	}
}

func solvePart1(crane *part1Crane) string {
	log.Debug("Solving part 1.")

	crane.ExecuteInstructions()
	topCrates := crane.PeekTopCrates()

	log.Debug("Part 1 solved.")
	return string(topCrates)
}

func solvePart2(crane *part2Crane) string {
	log.Debug("Solving part 2.")

	crane.ExecuteInstructions()
	topCrates := crane.PeekTopCrates()

	log.Debug("Part 2 solved.")
	return string(topCrates)
}
