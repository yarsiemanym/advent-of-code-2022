package day09

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

var up *common.Point = common.New2DPoint(0, 1)
var right *common.Point = common.New2DPoint(1, 0)
var down *common.Point = common.New2DPoint(0, -1)
var left *common.Point = common.New2DPoint(-1, 0)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseMotion)

	var motions []*motion
	for _, result := range results {
		motions = append(motions, result.(*motion))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(motions),
		Part2: solvePart2(motions),
	}
}

func solvePart1(motions []*motion) string {
	log.Debug("Solving part 1.")

	headPosition := common.New2DPoint(0, 0)
	tailPosition := common.New2DPoint(0, 0)
	positionsVisitedByTail := map[common.Point]int{}
	positionsVisitedByTail[*tailPosition] = 1

	for _, motion := range motions {

		log.Debugf("Moving head %d steps %c.", motion.Steps, motion.Direction)

		for step := 0; step < motion.Steps; step++ {
			headPosition = move(headPosition, motion.Direction)
			delta := headPosition.Subtract(tailPosition)

			if common.AbsInt(delta.X()) > 1 || common.AbsInt(delta.Y()) > 1 {
				tailPosition = determineNewTailPosition(headPosition, tailPosition)
				positionsVisitedByTail[*tailPosition] = positionsVisitedByTail[*tailPosition] + 1
			}

			log.Debugf("headPosition = %s", headPosition)
			log.Debugf("tailPosition = %s", tailPosition)
		}
	}

	tailPositionCount := 0

	for range positionsVisitedByTail {
		tailPositionCount++
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(tailPositionCount)
}

func solvePart2(motions []*motion) string {
	log.Debug("Solving part 2.")

	knots := []*common.Point{
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
		common.New2DPoint(0, 0),
	}

	positionsVisitedByTail := map[common.Point]int{}
	positionsVisitedByTail[*knots[9]] = 1

	for _, motion := range motions {

		log.Debugf("Moving head %d steps %c.", motion.Steps, motion.Direction)

		for step := 0; step < motion.Steps; step++ {
			knots[0] = move(knots[0], motion.Direction)

			for knotIndex := 1; knotIndex < len(knots); knotIndex++ {
				headPosition := knots[knotIndex-1]
				tailPosition := knots[knotIndex]
				delta := headPosition.Subtract(tailPosition)

				if common.AbsInt(delta.X()) > 1 || common.AbsInt(delta.Y()) > 1 {
					tailPosition = determineNewTailPosition(headPosition, tailPosition)
					knots[knotIndex] = tailPosition

					if knotIndex == len(knots)-1 {
						positionsVisitedByTail[*tailPosition] = positionsVisitedByTail[*tailPosition] + 1
					}
				}
			}

			for knotIndex := 0; knotIndex < len(knots); knotIndex++ {
				log.Debugf("knots[%d] = %s", knotIndex, knots[knotIndex])
			}
		}
	}

	tailPositionCount := 0

	for range positionsVisitedByTail {
		tailPositionCount++
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(tailPositionCount)
}

func move(start *common.Point, direction rune) *common.Point {
	var end *common.Point = nil

	switch direction {
	case 'U':
		end = start.Add(up)
	case 'R':
		end = start.Add(right)
	case 'D':
		end = start.Add(down)
	case 'L':
		end = start.Add(left)
	default:
		log.Panicf("Invalid direction %c.", direction)
	}

	return end
}

func determineNewTailPosition(headPosition *common.Point, tailPosition *common.Point) *common.Point {
	delta := headPosition.Subtract(tailPosition)

	if delta.Y() > 1 || (delta.Y() == 1 && common.AbsInt(delta.X()) > 1) {
		tailPosition = move(tailPosition, 'U')
	}

	if delta.X() > 1 || (delta.X() == 1 && common.AbsInt(delta.Y()) > 1) {
		tailPosition = move(tailPosition, 'R')
	}

	if delta.Y() < -1 || (delta.Y() == -1 && common.AbsInt(delta.X()) > 1) {
		tailPosition = move(tailPosition, 'D')
	}

	if delta.X() < -1 || (delta.X() == -1 && common.AbsInt(delta.Y()) > 1) {
		tailPosition = move(tailPosition, 'L')
	}

	return tailPosition
}
