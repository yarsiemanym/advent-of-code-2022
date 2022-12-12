package day11

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       11,
		InputFile: "test1.txt",
	}

	expectedPart1 := "10605"
	expectedPart2 := "2713310158"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}