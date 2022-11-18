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

// Only orthagonally adjacent, excluding self.
func (point *Point) GetVonNeumannNeighbors() []*Point {
	return []*Point{
		New2DPoint(point.x, point.y+1),
		New2DPoint(point.x-1, point.y),
		New2DPoint(point.x+1, point.y),
		New2DPoint(point.x, point.y-1),
	}
}

// Only orthagonally adjacent, excluding self.
func (point *Point) GetVonNeumannNeighborhood() []*Point {
	return []*Point{
		New2DPoint(point.x, point.y+1),
		New2DPoint(point.x-1, point.y),
		New2DPoint(point.x, point.y),
		New2DPoint(point.x+1, point.y),
		New2DPoint(point.x, point.y-1),
	}
}

// Both orthagonally and diagonally adjacent, including self.
func (point *Point) GetMooreNeighbors() []*Point {
	return []*Point{
		New2DPoint(point.x-1, point.y-1),
		New2DPoint(point.x, point.y-1),
		New2DPoint(point.x+1, point.y-1),
		New2DPoint(point.x-1, point.y),
		New2DPoint(point.x+1, point.y),
		New2DPoint(point.x-1, point.y+1),
		New2DPoint(point.x, point.y+1),
		New2DPoint(point.x+1, point.y+1),
	}
}

// Both orthagonally and diagonally adjacent, including self.
func (point *Point) GetMooreNeighborhood() []*Point {
	return []*Point{
		New2DPoint(point.x-1, point.y-1),
		New2DPoint(point.x, point.y-1),
		New2DPoint(point.x+1, point.y-1),
		New2DPoint(point.x-1, point.y),
		New2DPoint(point.x, point.y),
		New2DPoint(point.x+1, point.y),
		New2DPoint(point.x-1, point.y+1),
		New2DPoint(point.x, point.y+1),
		New2DPoint(point.x+1, point.y+1),
	}
}

func (point *Point) String() string {
	return fmt.Sprintf("(%d,%d,%d)", point.x, point.y, point.z)
}

func (point *Point) Hash() int64 {
	return int64((point.X())<<42 + (point.Y() << 21) + point.Z())
}
