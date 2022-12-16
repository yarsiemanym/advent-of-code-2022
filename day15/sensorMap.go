package day15

import (
	"strconv"
	"strings"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type sensorMap struct {
	sensors     []*sensor
	inspectionY int
	searchSpace int
}

func (thisSensorMap *sensorMap) Inspect() int {
	coveredPoints := map[common.Point]int{}

	for _, sensor := range thisSensorMap.sensors {
		if sensor.position.Y()-sensor.CoverageRadius() <= thisSensorMap.inspectionY &&
			sensor.position.Y()+sensor.CoverageRadius() >= thisSensorMap.inspectionY {

			distance := common.AbsInt(sensor.position.Y() - thisSensorMap.inspectionY)
			extraDistance := sensor.CoverageRadius() - distance

			for x := sensor.position.X() - extraDistance; x <= sensor.position.X()+extraDistance; x++ {
				point := common.New2DPoint(x, thisSensorMap.inspectionY)

				if *point != *sensor.closestBeacon {
					coveredPoints[*point] = coveredPoints[*point] + 1
				}
			}
		}
	}

	count := 0

	for range coveredPoints {
		count++
	}

	return count
}

func (thisSensorMap *sensorMap) FindTuningFrequency() uint64 {
	frequency := uint64(0)

	for y := 0; y <= thisSensorMap.searchSpace; y++ {
		for x := 0; x <= thisSensorMap.searchSpace; x++ {
			point := common.New2DPoint(x, y)
			isCovered := false
			for _, sensor := range thisSensorMap.sensors {
				if sensor.Covers(point) {
					isCovered = true
					break
				}
			}

			if !isCovered {
				frequency = (uint64(point.X()) * uint64(4000000)) + uint64(point.Y())
			}
		}
	}

	return frequency
}

func parseSensorMap(text string) *sensorMap {

	tokens := common.Split(text, "\n\n")

	sensorLines := common.Split(tokens[0], "\n")
	sensors := []*sensor{}
	minY, maxY := 0, 0

	for _, line := range sensorLines {
		result := parseSensor(line)

		if result == nil {
			continue
		}

		thisSensor := result.(*sensor)
		sensors = append(sensors, thisSensor)

		minY = common.MinInt(minY, thisSensor.position.Y(), thisSensor.closestBeacon.Y())
		maxY = common.MaxInt(maxY, thisSensor.position.Y(), thisSensor.closestBeacon.Y())
	}

	inspectionY, err := strconv.Atoi(strings.Trim(tokens[1], " \n"))
	common.Check(err)

	return &sensorMap{
		sensors:     sensors,
		inspectionY: inspectionY,
		searchSpace: 2 * inspectionY,
	}
}
