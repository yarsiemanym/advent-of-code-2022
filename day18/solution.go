package day18

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	droplet := parseDroplet(common.ReadFile(puzzle.InputFile))

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(droplet),
		Part2: solvePart2(droplet),
	}
}

func solvePart1(droplet *droplet) string {
	log.Debug("Solving part 1.")

	surfaceArea := droplet.SurfaceArea()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(surfaceArea)
}

func solvePart2(droplet *droplet) string {
	log.Debug("Solving part 2.")

	droplet.DetectPockets()
	surfaceArea := droplet.SurfaceArea()

	log.Debug("Part 2 solved.")
	return strconv.Itoa(surfaceArea)
}
