package day08

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	forest := parseForest(common.ReadFile(puzzle.InputFile))

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(forest),
		Part2: solvePart2(forest),
	}
}

func solvePart1(forest *forest) string {
	log.Debug("Solving part 1.")

	visibleTreePoints := forest.GetVisibleTreePoints()
	count := len(visibleTreePoints)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(forest *forest) string {
	log.Debug("Solving part 2.")

	highestScenicScore := forest.GetHighestScenicScore()

	log.Debug("Part 2 solved.")
	return strconv.Itoa(highestScenicScore)
}
