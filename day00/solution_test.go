package day00

import (
	"strconv"
	"testing"
	"time"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Test_Solve_Test1(t *testing.T) {
	puzzle := &common.Puzzle{
		Year:      2022,
		Day:       0,
		InputFile: "test1.txt",
	}

	expectedName := "Joe Schmoe"
	birthday, err := time.Parse(common.ShortDateFormat, "1983-11-24")
	common.Check(err)
	expectedAge := int(time.Since(birthday).Hours() / 24 / 365)

	answer := Solve(puzzle)

	if answer.Part1 != expectedName {
		t.Errorf("Expected: %v, Actual:%v", expectedName, answer.Part1)
	}

	if answer.Part2 != strconv.Itoa(expectedAge) {
		t.Errorf("Expected: %v, Actual: %v", expectedAge, answer.Part2)
	}
}
