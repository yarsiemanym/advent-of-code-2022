package day15

import (
	"regexp"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type sensor struct {
	position      *common.Point
	closestBeacon *common.Point
}

func (thisSensor *sensor) CoverageRadius() int {
	return thisSensor.position.ManhattanDistance(thisSensor.closestBeacon)
}

func (thisSensor *sensor) Covers(point *common.Point) bool {
	return thisSensor.position.ManhattanDistance(point) <= thisSensor.CoverageRadius()
}

func (thisSensor *sensor) Frontier() []*common.LineSegment {
	frontier := []*common.LineSegment{}

	corner1 := common.New2DPoint(thisSensor.position.X()-thisSensor.CoverageRadius()-1, thisSensor.position.Y())
	corner2 := common.New2DPoint(thisSensor.position.X(), thisSensor.position.Y()-thisSensor.CoverageRadius()-1)
	corner3 := common.New2DPoint(thisSensor.position.X()+thisSensor.CoverageRadius()+1, thisSensor.position.Y())
	corner4 := common.New2DPoint(thisSensor.position.X(), thisSensor.position.Y()+thisSensor.CoverageRadius()+1)

	edge1 := common.NewLineSegment(corner1, corner2)
	edge2 := common.NewLineSegment(corner2, corner3)
	edge3 := common.NewLineSegment(corner3, corner4)
	edge4 := common.NewLineSegment(corner4, corner1)

	frontier = append(frontier, edge1, edge2, edge3, edge4)

	return frontier
}

func parseSensor(text string) any {

	if text == "" {
		return nil
	}

	sensorRegexp := regexp.MustCompile(`Sensor at x=([-\d]+), y=([-\d]+): closest beacon is at x=([-\d]+), y=([-\d]+)`)
	matches := sensorRegexp.FindStringSubmatch(text)

	sensorX, err := strconv.Atoi(matches[1])
	common.Check(err)

	sensorY, err := strconv.Atoi(matches[2])
	common.Check(err)

	beaconX, err := strconv.Atoi(matches[3])
	common.Check(err)

	beaconY, err := strconv.Atoi(matches[4])
	common.Check(err)

	return &sensor{
		position:      common.New2DPoint(sensorX, sensorY),
		closestBeacon: common.New2DPoint(beaconX, beaconY),
	}
}
