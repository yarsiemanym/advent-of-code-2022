package day04

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseCleaningAssignmentPair)

	var pairs []*cleaningAssignmentPair
	for _, result := range results {
		pairs = append(pairs, result.(*cleaningAssignmentPair))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(pairs),
		Part2: solvePart2(pairs),
	}
}

func solvePart1(pairs []*cleaningAssignmentPair) string {
	log.Debug("Solving part 1.")

	count := 0

	for _, pair := range pairs {
		if pair.OneContainsTheOther() {
			count++
		}
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(count)
}

func solvePart2(pairs []*cleaningAssignmentPair) string {
	log.Debug("Solving part 2.")

	count := 0

	for _, pair := range pairs {
		if pair.HasOverlap() {
			count++
		}
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(count)
}
