package day09

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       9,
		InputFile: "test1.txt",
	}

	expectedPart1 := "13"
	expectedPart2 := "1"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test2(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       9,
		InputFile: "test2.txt",
	}

	expectedPart1 := "88"
	expectedPart2 := "36"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}
