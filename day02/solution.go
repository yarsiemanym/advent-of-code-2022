package day02

import (
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {
	results := common.ParseFile(puzzle.InputFile, "\n", parsePart1Round)

	var part1Rounds []*part1Round
	for _, result := range results {
		part1Rounds = append(part1Rounds, result.(*part1Round))
	}

	results = common.ParseFile(puzzle.InputFile, "\n", parsePart2Round)

	var part2Rounds []*part2Round
	for _, result := range results {
		part2Rounds = append(part2Rounds, result.(*part2Round))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(part1Rounds),
		Part2: solvePart2(part2Rounds),
	}
}

func solvePart1(rounds []*part1Round) string {
	finalScore := 0

	for _, round := range rounds {
		finalScore += round.YourScore()
	}

	return strconv.Itoa(finalScore)
}

func solvePart2(rounds []*part2Round) string {
	finalScore := 0

	for _, round := range rounds {
		finalScore += round.YourScore()
	}

	return strconv.Itoa(finalScore)
}
