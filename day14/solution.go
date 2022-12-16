package day14

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	cave1 := parseCave(common.ReadFile(puzzle.InputFile))
	cave2 := parseCave(common.ReadFile(puzzle.InputFile))

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(cave1),
		Part2: solvePart2(cave2),
	}
}

func solvePart1(cave *cave) string {
	log.Debug("Solving part 1.")

	count := cave.FillWithSand()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(cave *cave) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}
