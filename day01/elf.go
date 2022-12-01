package day01

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
