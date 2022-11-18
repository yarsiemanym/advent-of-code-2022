package common

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

type Puzzle struct {
	Year      int
	Day       int
	InputFile string
	solution  Solution
}

func (puzzle *Puzzle) SetSolution(solution Solution) {
	puzzle.solution = solution
}

func (puzzle *Puzzle) Solve() Answer {

	if !puzzle.IsUnlocked() {
		log.Fatalf("Day %v has not been unlocked.", puzzle.Day)
	}

	log.Info("Ensuring input file exists.")
	puzzle.EnsureInputFileExists()
	log.Info("Input file exists.")

	if puzzle.solution == nil {
		log.Fatal("Solution is not set.")
	}

	log.Info("Solving puzzle.")
	startTime := time.Now()
	answer := puzzle.solution(puzzle)
	elapsedTime := time.Since(startTime)
	answer.ExecutionTime = elapsedTime
	log.Info("Puzzle solved!")

	return answer
}

func (puzzle *Puzzle) IsUnlocked() bool {
	log.Debugf("Checking if day %v has been unlocked.", puzzle.Day)

	est, err := time.LoadLocation(("EST"))
	Check(err)

	var puzzleUnlockAt time.Time
	if puzzle.Day != 0 {
		puzzleUnlockAt = time.Date(puzzle.Year, 11, 30, 0, 0, 0, 0, est).Add(time.Hour * 24 * time.Duration(puzzle.Day))
	}
	log.Tracef("puzzleUnlockAt = \"%v\"", puzzleUnlockAt)

	isUnlocked := puzzleUnlockAt.Before(time.Now())
	log.Tracef("isUnlocked = %v", isUnlocked)

	return isUnlocked
}

func (puzzle *Puzzle) EnsureInputFileExists() string {
	log.Debug("Checking existence of input file.")

	if puzzle.InputFile == "" {
		puzzle.InputFile = fmt.Sprintf("./day%02d/input.txt", puzzle.Day)
	}

	target := puzzle.InputFile
	log.Tracef("target = \"%v\"", target)

	_, err := os.Open(target)

	if err != nil {
		log.Debug("Local copy of input file not found.")
		url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", puzzle.Year, puzzle.Day)
		DownloadFile(url, target)
	} else {
		log.Debug("Local copy of input file found.")
	}

	puzzle.InputFile = target
	return target
}
