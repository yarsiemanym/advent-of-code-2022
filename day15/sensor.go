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
