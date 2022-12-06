package day06

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       6,
		InputFile: "test1.txt",
	}

	expectedPart1 := "7"
	expectedPart2 := "19"
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
		Day:       6,
		InputFile: "test2.txt",
	}

	expectedPart1 := "5"
	expectedPart2 := "23"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test3(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       6,
		InputFile: "test3.txt",
	}

	expectedPart1 := "6"
	expectedPart2 := "23"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test4(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       6,
		InputFile: "test4.txt",
	}

	expectedPart1 := "10"
	expectedPart2 := "29"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}

func Test_Solve_Test5(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       6,
		InputFile: "test5.txt",
	}

	expectedPart1 := "11"
	expectedPart2 := "26"
	answer := Solve(puzzle)

	if answer.Part1 != expectedPart1 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart1, answer.Part1)
	}

	if answer.Part2 != expectedPart2 {
		t.Errorf("Expected: %v, Actual: %v", expectedPart2, answer.Part2)
	}
}
