package day05

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	text := common.ReadFile(puzzle.InputFile)
	crane1 := parseCrane(text)
	crane2 := parseCrane(text)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(crane1),
		Part2: solvePart2(crane2),
	}
}

func solvePart1(crane *crane) string {
	log.Debug("Solving part 1.")

	crane.ExecuteInstructions(crateMover9000)
	topCrates := crane.PeekTopCrates()

	log.Debug("Part 1 solved.")
	return string(topCrates)
}

func solvePart2(crane *crane) string {
	log.Debug("Solving part 2.")

	crane.ExecuteInstructions(crateMover9001)
	topCrates := crane.PeekTopCrates()

	log.Debug("Part 2 solved.")
	return string(topCrates)
}
