package day02

import (
	"strconv"

	log "github.com/sirupsen/logrus"

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
	log.Debug("Solving part 1.")

	finalScore := 0

	for index, round := range rounds {
		log.Debugf("Scoring round %d.", index)
		score := round.YourScore()
		log.Debugf("score = %d", score)
		finalScore += score
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(finalScore)
}

func solvePart2(rounds []*part2Round) string {
	log.Debug("Solving part 2.")

	finalScore := 0

	for index, round := range rounds {
		log.Debugf("Scoring round %d.", index)
		score := round.YourScore()
		log.Debugf("score = %d", score)
		finalScore += round.YourScore()
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(finalScore)
}
