package day01

import (
	"sort"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n\n", parseElf)

	var elves []elf
	for _, result := range results {
		elves = append(elves, result.(elf))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(elves),
		Part2: solvePart2(elves),
	}
}

func solvePart1(elves []elf) string {
	mostCalories := 0

	for _, elf := range elves {
		caloriesCarried := elf.CountTotalCaloriesCarried()
		if caloriesCarried > mostCalories {
			mostCalories = caloriesCarried
		}
	}

	return strconv.Itoa(mostCalories)
}

func solvePart2(elves []elf) string {
	caloriesCarriedByElf := make([]int, 0)

	for _, elf := range elves {
		caloriesCarriedByElf = append(caloriesCarriedByElf, elf.CountTotalCaloriesCarried())
	}

	sort.Ints(caloriesCarriedByElf)
	length := len(caloriesCarriedByElf)
	topThree := caloriesCarriedByElf[length-3:]
	sumOfTopThreeCalories := 0

	for _, calories := range topThree {
		sumOfTopThreeCalories += calories
	}

	return strconv.Itoa(sumOfTopThreeCalories)
}

func parseElf(text string) interface{} {
	lines := common.Split(text, "\n")

	elf := elf{}

	for _, line := range lines {
		calories, err := strconv.Atoi(line)
		common.Check(err)

		elf.Foods = append(elf.Foods, calories)
	}

	return elf
}
