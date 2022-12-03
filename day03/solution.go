package day03

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parseRucksack)

	var rucksacks []*rucksack
	for _, result := range results {
		rucksacks = append(rucksacks, result.(*rucksack))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(rucksacks),
		Part2: solvePart2(rucksacks),
	}
}

func solvePart1(rucksacks []*rucksack) string {
	log.Debug("Solving part 1.")

	priority := 0

	for _, thisRucksack := range rucksacks {
		packingErrors := thisRucksack.FindPackingErrors()
		priority += determinePriority(packingErrors)
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(priority)
}

func solvePart2(rucksacks []*rucksack) string {
	log.Debug("Solving part 2.")

	priority := 0

	for index := 2; index < len(rucksacks); index += 3 {
		group := rucksacks[index-2 : index+1]
		badge := determineBadge(group)
		priority += determinePriority([]rune{badge})
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(priority)
}
