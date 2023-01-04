package day19

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseBlueprint)

	blueprints := map[int]*blueprint{}
	for _, result := range results {
		blueprint := result.(*blueprint)
		blueprints[blueprint.id] = blueprint
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(blueprints),
		Part2: solvePart2(blueprints),
	}
}

func solvePart1(blueprints map[int]*blueprint) string {
	log.Debug("Solving part 1.")

	qualityLevel := 0

	for id := 1; id <= len(blueprints); id++ {
		blueprint := blueprints[id]
		simulation := NewSimulation(blueprint, 24)
		geodeCount := simulation.Run()
		log.Infof("Blueprint %d yielded %d geodes.", id, geodeCount)
		qualityLevel += id * geodeCount

	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(qualityLevel)
}

func solvePart2(blueprints map[int]*blueprint) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}
