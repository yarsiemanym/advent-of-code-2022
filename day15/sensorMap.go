package day15

import (
	"log"
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
	candidates := thisSensorMap.FindUncoveredPoints()

	if len(candidates) == 0 {
		log.Panic("No candidates.")
	} else if len(candidates) != 1 {
		log.Panicf("Too many candidates: %d", len(candidates))
	}

	distressBeaconLocation := candidates[0]

	tuningFrequency := (uint64(distressBeaconLocation.X()) * uint64(4000000)) + uint64(distressBeaconLocation.Y())
	return tuningFrequency
}

func (thisSensorMap *sensorMap) FindUncoveredPoints() []*common.Point {
	uncoveredPoints := []*common.Point{}

	for _, thisSensor := range thisSensorMap.sensors {
		for _, point := range thisSensor.Frontier() {
			if thisSensorMap.IsInSearchSpace(point) {
				covered := false

				for _, otherSensor := range thisSensorMap.sensors {
					if thisSensor != otherSensor && otherSensor.Covers(point) {
						covered = true
					}
				}

				if !covered {
					uncoveredPoints = common.UnionPointers(uncoveredPoints, []*common.Point{point})
				}
			}
		}
	}

	return uncoveredPoints
}

func (thisSensorMap *sensorMap) IsInSearchSpace(point *common.Point) bool {
	return point.X() >= 0 && point.X() <= thisSensorMap.searchSpace && point.Y() >= 0 && point.Y() <= thisSensorMap.searchSpace
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
