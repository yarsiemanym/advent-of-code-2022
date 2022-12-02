package day01

import (
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type elf struct {

	// Each element is a food item and the value is the number of calories that item has.
	Foods []int
}

// Returns the total number of calories carried by this elf.
func (thisElf *elf) CountTotalCaloriesCarried() int {
	totalCaloriesCarried := 0

	for _, calories := range thisElf.Foods {
		totalCaloriesCarried += calories
	}

	return totalCaloriesCarried
}

func parseElf(text string) interface{} {
	if text == "" {
		return nil
	}

	lines := common.Split(text, "\n")

	elf := &elf{}

	for _, line := range lines {
		calories, err := strconv.Atoi(line)
		common.Check(err)

		elf.Foods = append(elf.Foods, calories)
	}

	return elf
}
