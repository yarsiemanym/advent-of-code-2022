package day14

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

const (
	empty   = '.'
	rock    = '#'
	opening = '+'
	sand    = 'O'
)

var source = common.New2DPoint(500, 0)

type cave struct {
	pointMap map[common.Point]rune
	bottom   int
	left     int
	right    int
}

func (thisCave *cave) FillWithSand() int {
	keepGoing := true
	count := 0

	for keepGoing {
		count++
		log.Debugf("Dropping unit of sand #%d.", count)
		keepGoing = thisCave.DropUnitOfSand()
	}

	return count - 1
}

func (thisCave *cave) DropUnitOfSand() bool {
	current := common.New2DPoint(500, 0)
	atRest := false
	flowingIntoTheAbyss := false

	for !atRest && !flowingIntoTheAbyss {

		if current.Y() >= thisCave.bottom {
			log.Debug("Falling into the abyss.")
			flowingIntoTheAbyss = true
			continue
		}

		next := common.New2DPoint(current.X(), current.Y()+1)
		if thisCave.IsEmptyAt(next) {
			log.Debug("Falling straight down.")
			current = next
			continue
		}

		next = common.New2DPoint(current.X()-1, current.Y()+1)
		if thisCave.IsEmptyAt(next) {
			log.Debug("Falling down and to the left.")
			current = next
			continue
		}

		next = common.New2DPoint(current.X()+1, current.Y()+1)
		if thisCave.IsEmptyAt(next) {
			log.Debug("Falling down and to the right.")
			current = next
			continue
		}

		thisCave.SettleSandAt(current)
		atRest = true
	}

	return !flowingIntoTheAbyss
}

func (thisCave *cave) IsEmptyAt(point *common.Point) bool {
	return thisCave.ValueAt(point) == empty
}

func (thisCave *cave) ValueAt(point *common.Point) rune {
	value, exists := thisCave.pointMap[*point]

	if !exists {
		return empty
	} else {
		return value
	}
}

func (thisCave *cave) SettleSandAt(point *common.Point) {
	log.Debugf("Coming to a rest %s.", point)
	thisCave.pointMap[*point] = sand
}

func (thisCave *cave) String() string {
	output := ""

	for y := 0; y <= thisCave.bottom; y++ {
		for x := thisCave.left; x <= thisCave.right; x++ {
			point := common.New2DPoint(x, y)
			value := thisCave.ValueAt(point)
			output += string(value)
		}
		output += "\n"
	}

	return output
}

func parseCave(text string) *cave {
	lines := common.Split(text, "\n")
	pointMap := map[common.Point]rune{}
	pointMap[*source] = opening
	bottom, left, right := source.Y(), source.X(), source.X()

	for _, line := range lines {
		pointStrings := common.Split(line, " -> ")

		for index := 1; index < len(pointStrings); index++ {
			start := parsePoint(pointStrings[index-1])
			end := parsePoint(pointStrings[index])
			xDirection := common.Sign(end.X() - start.X())
			yDirection := common.Sign(end.Y() - start.Y())

			for x, y := start.X(), start.Y(); ; {
				point := common.New2DPoint(x, y)
				log.Debugf("Placing rock at point %s.", point)
				pointMap[*point] = rock
				bottom = common.MaxInt(bottom, point.Y())
				left = common.MinInt(left, point.X())
				right = common.MaxInt(right, point.X())

				if x == end.X() && y == end.Y() {
					break
				} else {
					if x != end.X() {
						x += xDirection
					}

					if y != end.Y() {
						y += yDirection
					}
				}
			}
		}
	}

	return &cave{
		pointMap: pointMap,
		bottom:   bottom,
		left:     left,
		right:    right,
	}
}

func parsePoint(text string) *common.Point {
	log.Debugf("Parsing point \"%s\".", text)
	tokens := common.Split(text, ",")
	x, err := strconv.Atoi(tokens[0])
	common.Check(err)
	y, err := strconv.Atoi(tokens[1])
	common.Check(err)
	return common.New2DPoint(x, y)
}
