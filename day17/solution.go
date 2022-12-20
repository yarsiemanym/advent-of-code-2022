package day17

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
	"github.com/yarsiemanym/advent-of-code-2022/vt100"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	jetsPattern := common.ParseFile(puzzle.InputFile, "\n", func(text string) any { return text })[0].(string)

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(jetsPattern),
		Part2: solvePart2(jetsPattern),
	}
}

func solvePart1(jetsPattern string) string {
	log.Debug("Solving part 1.")

	chamber := NewChamber(jetsPattern)
	ticks := 0

	vt100.ClearScreen()
	vt100.MoveCursorToHome()
	print(chamber, ticks)

	for rocks := 0; rocks < 2022; rocks = chamber.NumberOfSettledRocks() {
		chamber.Tick()
		ticks++
		print(chamber, ticks)
	}

	print(chamber, ticks)
	vt100.MoveCursorDown(54)

	height := chamber.HeightOfSettledRocks()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(height)
}

func solvePart2(jetsPattern string) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}

func print(chamber *chamber, ticks int) {
	vt100.SaveCursorPosition()
	output := chamber.String()
	output = strings.ReplaceAll(output, "|", vt100.Sprint("|", vt100.RedForegroundAttribute))
	output = strings.ReplaceAll(output, "-", vt100.Sprint("-", vt100.RedForegroundAttribute))
	output = strings.ReplaceAll(output, "+", vt100.Sprint("+", vt100.RedForegroundAttribute))
	output = strings.ReplaceAll(output, "#", vt100.Sprint("#", vt100.BlueForegroundAttribute))
	output = strings.ReplaceAll(output, "@", vt100.Sprint("@", vt100.GreenForegroundAttribute))
	vt100.Printf("Tick:\t%d\n", []any{ticks})
	vt100.Printf("Rocks:\t%d\n", []any{chamber.NumberOfSettledRocks()})
	vt100.Printf("Height:\t%d\n", []any{chamber.HeightOfSettledRocks()})
	vt100.Print(output)
	vt100.RestoreCursorPosition()
}
