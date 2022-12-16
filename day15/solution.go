package day15

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	sensorMap := parseSensorMap(common.ReadFile(puzzle.InputFile))

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(sensorMap),
		Part2: solvePart2(sensorMap),
	}
}

func solvePart1(sensorMap *sensorMap) string {
	log.Debug("Solving part 1.")

	count := sensorMap.Inspect()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(sensorMap *sensorMap) string {
	log.Debug("Solving part 2.")

	frequency := sensorMap.FindTuningFrequency()

	log.Debug("Part 2 solved.")
	return strconv.FormatUint(frequency, 10)
}
