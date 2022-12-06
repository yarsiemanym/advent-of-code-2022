package day06

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	buffer := &buffer{
		DataStream: common.ReadFile(puzzle.InputFile),
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(buffer),
		Part2: solvePart2(buffer),
	}
}

func solvePart1(buffer *buffer) string {
	log.Debug("Solving part 1.")

	position := buffer.StartOfPacket()

	log.Debug("Part 1 solved.")
	return strconv.Itoa(position)
}

func solvePart2(buffer *buffer) string {
	log.Debug("Solving part 2.")

	position := buffer.StartOfMessage()

	log.Debug("Part 2 solved.")
	return strconv.Itoa(position)
}
