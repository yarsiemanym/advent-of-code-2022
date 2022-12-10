package day10

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
	"github.com/yarsiemanym/advent-of-code-2022/vt100"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	text := common.ReadFile(puzzle.InputFile)
	instructions := common.Split(text, "\n")

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(instructions),
		Part2: solvePart2(instructions),
	}
}

func solvePart1(instructions []string) string {
	log.Debug("Solving part 1.")

	cpu := NewCPU(nil)
	cpu.ExecuteInstructions(instructions)
	signalStrength := cpu.SignalStrength()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(signalStrength)
}

func solvePart2(instructions []string) string {
	log.Debug("Solving part 2.")

	crt := NewCRT()
	cpu := NewCPU(drawPixelCallback(crt))

	cpu.ExecuteInstructions(instructions)
	output := ToVT100(crt.Render())

	log.Debug("Part 2 solved.")
	return output
}

func drawPixelCallback(crt *crt) cpuCycleCallback {
	return func(cycle int, registerX int) {
		crt.MaybeLightPixel(cycle, registerX)
	}
}

func ToVT100(output string) string {
	output = "\n" + output

	output = strings.Replace(output, "#", vt100.Sprint(" ", vt100.YellowBackgroundAttribute), -1)
	output = strings.Replace(output, ".", " ", -1)

	return output
}
