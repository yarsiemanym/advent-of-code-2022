package common

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
	z int
}

func New2DPoint(x int, y int) *Point {
	return &Point{
		x: x,
		y: y,
		z: 0,
	}
}

func New3DPoint(x int, y int, z int) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

func (point *Point) X() int {
	return point.x
}

func (point *Point) Y() int {
	return point.y
}

func (point *Point) Z() int {
	return point.z
}

func (point *Point) Add(other *Point) *Point {
	return &Point{
		x: point.X() + other.X(),
		y: point.Y() + other.Y(),
		z: point.Z() + other.Z(),
	}
}

func (point *Point) Subtract(other *Point) *Point {
	return &Point{
		x: point.X() - other.X(),
		y: point.Y() - other.Y(),
		z: point.Z() - other.Z(),
	}
}

func (point *Point) ManhattanDistance(otherPoint *Point) int {
	return AbsInt(point.x-otherPoint.x) + AbsInt(point.y-otherPoint.y) + AbsInt(point.z-otherPoint.z)
}

func (point *Point) Distance(otherPoint *Point) float64 {
	diff := point.Subtract(otherPoint)
	dist := math.Sqrt(math.Pow(float64(diff.X()), 2) + math.Pow(float64(diff.Y()), 2) + math.Pow(float64(diff.Z()), 2))
	return dist
}

func (point *Point) RotateXClockwise() *Point {
	return New3DPoint(point.X(), point.Z(), 0-point.Y())
}

func (point *Point) RotateXCounterClockwise() *Point {
	return New3DPoint(point.X(), 0-point.Z(), point.Y())
}

func (point *Point) RotateYClockwise() *Point {
	return New3DPoint(0-point.Z(), point.Y(), point.X())
}

func (point *Point) RotateYCounterClockwise() *Point {
	return New3DPoint(point.Z(), point.Y(), 0-point.X())
}

func (point *Point) RotateZClockwise() *Point {
	return New3DPoint(point.Y(), 0-point.X(), point.Z())
}

func (point *Point) RotateZCounterClockwise() *Point {
	return New3DPoint(0-point.Y(), point.X(), point.Z())
}

// Only orthagonally adjacent.
func (center *Point) Get2DVonNeumannNeighbors(radius int) []*Point {
	neighborhood := center.Get2DVonNeumannNeighborhood(radius)
	neighborhood = ExclusivePointers(neighborhood, []*Point{center})
	return neighborhood
}

// Only orthagonally adjacent plus self.
func (center *Point) Get2DVonNeumannNeighborhood(radius int) []*Point {
	if radius <= 0 {
		return []*Point{
			New2DPoint(center.x, center.y),
		}
	}

	immediateNeighborhood := []*Point{
		New2DPoint(center.x, center.y+1),
		New2DPoint(center.x-1, center.y-1),
		New2DPoint(center.x, center.y),
		New2DPoint(center.x+1, center.y),
		New2DPoint(center.x+1, center.y+1),
	}

	extendedNeighborhood := []*Point{}

	for _, neighbor := range immediateNeighborhood {
		extendedNeighborhood = UnionPointers(extendedNeighborhood, neighbor.Get2DVonNeumannNeighborhood(radius-1))
	}

	extendedNeighborhood = UnionPointers(extendedNeighborhood, immediateNeighborhood)

	return extendedNeighborhood
}

// Both orthagonally and diagonally adjacent.
func (center *Point) Get2DMooreNeighbors(radius int) []*Point {
	neighborhood := center.Get2DMooreNeighborhood(radius)
	neighborhood = ExclusivePointers(neighborhood, []*Point{center})
	return neighborhood
}

// Both orthagonally and diagonally adjacent plus self.
func (center *Point) Get2DMooreNeighborhood(radius int) []*Point {
	neighbors := []*Point{}

	for x := center.x - radius; x <= center.x+radius; x++ {
		for y := center.y - radius; y <= center.y+radius; y++ {
			neighbors = append(neighbors, New2DPoint(x, y))
		}
	}

	return neighbors
}

// Only orthagonally adjacent.
func (center *Point) Get3DVonNeumannNeighbors(radius int) []*Point {
	neighborhood := center.Get3DVonNeumannNeighborhood(radius)
	neighborhood = ExclusivePointers(neighborhood, []*Point{center})
	return neighborhood
}

// Only orthagonally adjacent plus self.
func (center *Point) Get3DVonNeumannNeighborhood(radius int) []*Point {
	if radius <= 0 {
		return []*Point{}
	}

	immediateNeighbors := []*Point{
		New3DPoint(center.x, center.y+1, center.z),
		New3DPoint(center.x-1, center.y, center.z),
		New3DPoint(center.x+1, center.y, center.z),
		New3DPoint(center.x, center.y, center.z),
		New3DPoint(center.x, center.y-1, center.z),
		New3DPoint(center.x, center.y, center.z-1),
		New3DPoint(center.x, center.y, center.z+1),
	}

	neighbors := []*Point{}

	for _, neighbor := range immediateNeighbors {
		neighbors = UnionPointers(neighbors, neighbor.Get3DVonNeumannNeighborhood(radius-1))
	}

	neighbors = UnionPointers(neighbors, immediateNeighbors)

	return neighbors
}

// Both orthagonally and diagonally adjacent.
func (center *Point) Get3DMooreNeighbors(radius int) []*Point {
	neighborhood := center.Get3DMooreNeighborhood(radius)
	neighborhood = ExclusivePointers(neighborhood, []*Point{center})
	return neighborhood
}

// Both orthagonally and diagonally adjacent plus self.
func (center *Point) Get3DMooreNeighborhood(radius int) []*Point {
	neighbors := []*Point{}

	for x := center.x - radius; x <= center.x+radius; x++ {
		for y := center.y - radius; y <= center.y+radius; y++ {
			for z := center.z - radius; z <= center.z+radius; z++ {
				neighbors = append(neighbors, New2DPoint(x, y))
			}
		}
	}

	return neighbors
}

func (point *Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}

func (point *Point) Hash() int64 {
	return int64((point.X())<<42 + (point.Y() << 21) + point.Z())
}
