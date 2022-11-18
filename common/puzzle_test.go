package common

import (
	"os"
	"testing"
)

func Test_EnsureInputExists(t *testing.T) {
	InitSession()

	puzzle := Puzzle{
		Year: 2020,
		Day:  1,
	}

	expectedInputFile := "./day01/input.txt"
	os.Remove(expectedInputFile)

	actualInputFile := puzzle.EnsureInputFileExists()

	if actualInputFile != expectedInputFile {
		t.Errorf("Expected \"%v\" but got \"%v\".", expectedInputFile, actualInputFile)
	}

	_, err := os.Open(actualInputFile)

	if os.IsNotExist(err) {
		t.Errorf("Input file \"%v\" does not exist.", actualInputFile)
	}

	os.Remove("day01/input.txt")
	os.Remove("day01")
}

func Test_IsUnlocked_True(t *testing.T) {
	puzzle := Puzzle{
		Year: 2020,
		Day:  1,
	}

	isUnlocked := puzzle.IsUnlocked()

	if !isUnlocked {
		t.Errorf("Expected true but got %v.", isUnlocked)
	}
}

func Test_IsUnlocked_False(t *testing.T) {
	puzzle := Puzzle{
		Year: 2100,
		Day:  1,
	}

	isUnlocked := puzzle.IsUnlocked()

	if isUnlocked {
		t.Errorf("Expected false but got %v.", isUnlocked)
	}
}
