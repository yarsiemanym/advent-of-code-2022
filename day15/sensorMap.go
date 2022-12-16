package day15

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type sensorMap struct {
	sensors     []*sensor
	inspectionY int
	searchSpace int
}

func (thisSensorMap *sensorMap) Inspect() int {
	coveredPoints := map[common.Point]int{}
	count := 0

	for _, sensor := range thisSensorMap.sensors {
		if sensor.position.Y()-sensor.CoverageRadius() <= thisSensorMap.inspectionY &&
			sensor.position.Y()+sensor.CoverageRadius() >= thisSensorMap.inspectionY {

			yDistance := common.AbsInt(sensor.position.Y() - thisSensorMap.inspectionY)
			xDistance := sensor.CoverageRadius() - yDistance

			for x := sensor.position.X() - xDistance; x <= sensor.position.X()+xDistance; x++ {
				point := common.New2DPoint(x, thisSensorMap.inspectionY)

				if *point != *sensor.closestBeacon {
					coveredPoints[*point] = coveredPoints[*point] + 1
				}
			}
		}
	}

	for range coveredPoints {
		count++
	}

	return count
}

func (thisSensorMap *sensorMap) FindTuningFrequency() uint64 {
	edges := []*common.LineSegment{}

	for _, sensor := range thisSensorMap.sensors {
		edges = append(edges, sensor.Frontier()...)
	}

	intersections := map[common.Point]int{}

	for _, thisEdge := range edges {
		for _, otherEdge := range edges {
			if thisEdge != otherEdge {
				intersection := findIntersection(thisEdge, otherEdge)

				if intersection != nil && thisSensorMap.IsInSearchSpace(intersection) {
					intersections[*intersection] = intersections[*intersection] + 1
				}
			}
		}
	}

	var distressBeaconLocation *common.Point = nil

	for point := range intersections {
		iscovered := false
		for _, sensor := range thisSensorMap.sensors {
			if sensor.Covers(&point) {
				iscovered = true
				break
			}
		}

		if !iscovered {
			distressBeaconLocation = &point
			break
		}
	}

	log.Debugf("distressBeaconLocation = %s", distressBeaconLocation)

	tuningFrequency := (uint64(distressBeaconLocation.X()) * uint64(4000000)) + uint64(distressBeaconLocation.Y())
	return tuningFrequency
}

func (thisSensorMap *sensorMap) IsInSearchSpace(point *common.Point) bool {
	return point.X() >= 0 && point.X() <= thisSensorMap.searchSpace && point.Y() >= 0 && point.Y() <= thisSensorMap.searchSpace
}

func findIntersection(line1 *common.LineSegment, line2 *common.LineSegment) *common.Point {
	a1, a2 := line1.Start().X()-line1.End().X(), line2.Start().X()-line2.End().X()
	b1, b2 := line1.Start().Y()-line1.End().Y(), line2.Start().Y()-line2.End().Y()

	det := determinant(a1, a2, b1, b2)

	if det == 0 {
		return nil
	} else {
		d1, d2 := determinant(line1.Start().X(), line1.Start().Y(), line1.End().X(), line1.End().Y()), determinant(line2.Start().X(), line2.Start().Y(), line2.End().X(), line2.End().Y())
		x := determinant(d1, d2, a1, a2) / det
		y := determinant(d1, d2, b1, b2) / det
		return common.New2DPoint(x, y)
	}
}

func determinant(a1 int, a2 int, b1 int, b2 int) int {
	return a1*b2 - a2*b1
}

func parseSensorMap(text string) *sensorMap {

	tokens := common.Split(text, "\n\n")

	sensorLines := common.Split(tokens[0], "\n")
	sensors := []*sensor{}

	for _, line := range sensorLines {
		result := parseSensor(line)

		if result == nil {
			continue
		}

		thisSensor := result.(*sensor)
		sensors = append(sensors, thisSensor)
	}

	inspectionY, err := strconv.Atoi(strings.Trim(tokens[1], " \n"))
	common.Check(err)

	return &sensorMap{
		sensors:     sensors,
		inspectionY: inspectionY,
		searchSpace: 2 * inspectionY,
	}
}
