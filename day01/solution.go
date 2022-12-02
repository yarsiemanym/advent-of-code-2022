package day01

import (
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n\n", parseElf)

	var elves []*elf
	for _, result := range results {
		elves = append(elves, result.(*elf))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(elves),
		Part2: solvePart2(elves),
	}
}

func solvePart1(elves []*elf) string {
	log.Debug("Solving part 1.")

	mostCalories := 0

	for index, elf := range elves {
		log.Debugf("Inspecting elf %d", index)
		caloriesCarried := elf.CountTotalCaloriesCarried()
		log.Debugf("Elf %d is carrying %d calories.", index, caloriesCarried)
		if caloriesCarried > mostCalories {
			log.Debug("New maximum found.")
			mostCalories = caloriesCarried
		}
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(mostCalories)
}

func solvePart2(elves []*elf) string {
	log.Debug("Solving part 2.")

	caloriesCarriedByElves := make([]int, 0)

	for index, elf := range elves {
		log.Debugf("Inspecting elf %d", index)
		caloriesCarriedByThisElf := elf.CountTotalCaloriesCarried()
		log.Debugf("Elf %d is carrying %d calories.", index, caloriesCarriedByThisElf)
		caloriesCarriedByElves = append(caloriesCarriedByElves, caloriesCarriedByThisElf)
	}

	log.Debugf("Sorting elves by total calories carried in descending order.")
	sort.Slice(caloriesCarriedByElves, func(i int, j int) bool { return caloriesCarriedByElves[i] > caloriesCarriedByElves[j] })

	log.Debugf("Selecting the top 3 elves.")
	topThreeCalories := caloriesCarriedByElves[0:3]
	log.Debugf("Top 3 elves are carrying %d, %d, and %d calories.", topThreeCalories[0], topThreeCalories[1], topThreeCalories[2])

	sumOfTopThreeCalories := 0

	for _, calories := range topThreeCalories {
		sumOfTopThreeCalories += calories
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(sumOfTopThreeCalories)
}
