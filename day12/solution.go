package day12

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	heightMap := parseHeightMap(common.ReadFile(puzzle.InputFile))

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(heightMap),
		Part2: solvePart2(heightMap),
	}
}

func solvePart1(heightMap *heightMap) string {
	log.Debug("Solving part 1.")

	path := heightMap.ShortestPath()
	steps := len(path) - 1

	log.Debug("Part 1 solved.")
	return strconv.Itoa(steps)
}

func solvePart2(heightMap *heightMap) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}
