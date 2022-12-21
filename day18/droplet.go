package day18

import (
	"math"
	"strconv"
	"time"

	"github.com/yarsiemanym/advent-of-code-2022/common"
	"github.com/yarsiemanym/advent-of-code-2022/vt100"
)

const (
	unknown = 0
	lava    = 1
	pocket  = 2
	water   = 3
)

type droplet struct {
	cubeMap map[common.Point]int
	span    *common.LineSegment
}

func (droplet *droplet) SurfaceArea() int {
	surfaceArea := 0

	for point := range droplet.cubeMap {
		surfaceArea += droplet.SurfaceAreaAt(&point)
	}

	return surfaceArea
}

func (droplet *droplet) SurfaceAreaAt(point *common.Point) int {
	surfaceArea := 0
	pointState := droplet.cubeMap[*point]

	if pointState != lava {
		return surfaceArea
	}

	neighboringPoints := point.Get3DVonNeumannNeighbors(1)

	for _, neighbor := range neighboringPoints {
		neighborState, exists := droplet.cubeMap[*neighbor]

		if !exists || neighborState == water {
			surfaceArea++
		}
	}

	return surfaceArea
}

func (droplet *droplet) DetectPockets() {
	for z := droplet.span.Start().Z(); z <= droplet.span.End().Z(); z++ {
		for y := droplet.span.Start().Y(); y <= droplet.span.End().Y(); y++ {
			for x := droplet.span.Start().X(); x <= droplet.span.End().X(); x++ {

				point := common.New3DPoint(x, y, z)
				_, exists := droplet.cubeMap[*point]

				if !exists {
					droplet.DetectPocketAt(point, []*common.Point{})
				}
			}
		}
	}

	// Second pass to make sure we didn't accidentally miscategorize any spaces as pockets.
	for z := droplet.span.End().Z(); z >= droplet.span.Start().Z(); z-- {
		for y := droplet.span.End().Y(); y >= droplet.span.Start().Y(); y-- {
			for x := droplet.span.End().X(); x >= droplet.span.Start().X(); x-- {

				point := common.New3DPoint(x, y, z)
				state, exists := droplet.cubeMap[*point]

				if !exists || state == pocket {
					droplet.DetectPocketAt(point, []*common.Point{})
				}
			}
		}
	}
}

func (droplet *droplet) DetectPocketAt(point *common.Point, ignorePoints []*common.Point) int {
	if point.X() < droplet.span.Start().X() || point.X() > droplet.span.End().X() ||
		point.Y() < droplet.span.Start().Y() || point.Y() > droplet.span.End().Y() ||
		point.Z() < droplet.span.Start().Z() || point.Z() > droplet.span.End().Z() {
		return water
	}

	neighboringPoints := point.Get3DVonNeumannNeighbors(1)
	ignorePoints = append(ignorePoints, point)

	for _, neighbor := range neighboringPoints {

		if common.ContainsPointer(ignorePoints, neighbor) {
			continue
		}

		neighborState, exists := droplet.cubeMap[*neighbor]

		if !exists {
			neighborState = droplet.DetectPocketAt(neighbor, ignorePoints)
		}

		if neighborState == water {
			droplet.cubeMap[*point] = water
			return water
		}
	}

	droplet.cubeMap[*point] = pocket
	return pocket
}

func (droplet *droplet) Print() {
	sleepInterval := 1000 * time.Millisecond

	for {

		for z := droplet.span.Start().Z(); z <= droplet.span.End().Z(); z++ {
			vt100.ClearScreen()
			vt100.MoveCursorToHome()
			vt100.Printf("Z=%d\n", []any{z})

			for y := droplet.span.Start().Y(); y <= droplet.span.End().Y(); y++ {
				for x := droplet.span.Start().X(); x <= droplet.span.End().X(); x++ {
					point := common.New3DPoint(x, y, z)
					state, exists := droplet.cubeMap[*point]

					if exists && state == lava {
						vt100.Print(" ", vt100.RedBackgroundAttribute)
					} else if exists && state == water {
						vt100.Print(" ", vt100.BlueBackgroundAttribute)
					} else {
						vt100.Print(" ")
					}
				}

				vt100.Println("")
			}

			time.Sleep(sleepInterval)
		}

		for z := droplet.span.End().Z(); z >= droplet.span.Start().Z(); z-- {
			vt100.ClearScreen()
			vt100.MoveCursorToHome()
			vt100.Printf("Z=%d\n", []any{z})

			for y := droplet.span.Start().Y(); y <= droplet.span.End().Y(); y++ {
				for x := droplet.span.Start().X(); x <= droplet.span.End().X(); x++ {
					point := common.New3DPoint(x, y, z)
					state, exists := droplet.cubeMap[*point]

					if exists && state == lava {
						vt100.Print(" ", vt100.RedBackgroundAttribute)
					} else if exists && state == water {
						vt100.Print(" ", vt100.BlueBackgroundAttribute)
					} else {
						vt100.Print(" ")
					}
				}

				vt100.Println("")
			}

			time.Sleep(sleepInterval)
		}
	}
}

func parseDroplet(text string) *droplet {
	if text == "" {
		return nil
	}

	minX, maxX := math.MaxInt, 0
	minY, maxY := math.MaxInt, 0
	minZ, maxZ := math.MaxInt, 0

	tokens := common.Split(text, "\n")
	cubeMap := map[common.Point]int{}

	for _, token := range tokens {
		cube := parsePoint(token)

		if cube != nil {
			minX = common.MinInt(minX, cube.X())
			maxX = common.MaxInt(maxX, cube.X())

			minY = common.MinInt(minY, cube.Y())
			maxY = common.MaxInt(maxY, cube.Y())

			minZ = common.MinInt(minZ, cube.Z())
			maxZ = common.MaxInt(maxZ, cube.Z())

			cubeMap[*cube] = lava
		}
	}

	span := common.NewLineSegment(common.New3DPoint(minX, minY, minZ), common.New3DPoint(maxX, maxY, maxZ))

	return &droplet{
		cubeMap: cubeMap,
		span:    span,
	}
}

func parsePoint(text string) *common.Point {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, ",")

	x, err := strconv.Atoi(tokens[0])
	common.Check(err)

	y, err := strconv.Atoi(tokens[1])
	common.Check(err)

	z, err := strconv.Atoi(tokens[2])
	common.Check(err)

	return common.New3DPoint(x, y, z)
}
