package common

import (
	"sort"
	"testing"
)

func Test_Point_Move(t *testing.T) {
	start := &Point{
		x: 1,
		y: 2,
		z: 3,
	}

	slope := &Point{
		x: 3,
		y: 4,
		z: -1,
	}

	end := start.Add(slope)

	if end.x != 4 {
		t.Errorf("Expected 4 but got %v.", end.x)
	}

	if end.y != 6 {
		t.Errorf("Expected 6 but got %v.", end.y)
	}

	if end.z != 2 {
		t.Errorf("Expected 2 but got %v.", end.z)
	}
}

func Test_Point_ManhattanDistance(t *testing.T) {
	point1 := New3DPoint(1, 2, -3)
	point2 := New3DPoint(-4, 0, 2)
	distance := point1.ManhattanDistance(point2)

	if distance != 12 {
		t.Errorf("Expected 12 but got %d.", distance)
	}
}

func Test_Point_Hash(t *testing.T) {
	point := New3DPoint(1, 2, 3)
	hash := point.Hash()

	if hash != 4398050705411 {
		t.Errorf("Expected 4398050705411 but got %d.", hash)
	}
}

func Test_Point_SortByHash(t *testing.T) {
	point1 := New3DPoint(10, 2, -6)
	point2 := New3DPoint(-23, 42, 93)
	point3 := New3DPoint(-23, -5, 9)
	point4 := New3DPoint(-23, -5, -5)
	points := []*Point{
		point1,
		point2,
		point3,
		point4,
	}

	sort.Slice(points, func(i int, j int) bool {
		point1 := points[i]
		point2 := points[j]
		return point1.Hash() < point2.Hash()
	})

	if len(points) != 4 {
		t.Errorf("Expected 4 but got %d.", len(points))
	}

	if points[0] != point4 {
		t.Errorf("Expected \"%s\" but got \"%s\".", point4, points[0])
	}

	if points[1] != point3 {
		t.Errorf("Expected \"%s\" but got \"%s\".", point3, points[1])
	}

	if points[2] != point2 {
		t.Errorf("Expected \"%s\" but got \"%s\".", point2, points[2])
	}

	if points[3] != point1 {
		t.Errorf("Expected \"%s\" but got \"%s\".", point1, points[3])
	}
}

func Test_Point_RotateX(t *testing.T) {
	point := New3DPoint(1, 2, 3)
	rotatedPoint := point.RotateXClockwise()

	if rotatedPoint == nil {
		t.Error("rotatedPoint is nil.")
	}

	if rotatedPoint.X() != 1 {
		t.Errorf("Expected 1 but got %d.", rotatedPoint.X())
	}

	if rotatedPoint.Y() != 3 {
		t.Errorf("Expected 3 but got %d.", rotatedPoint.Y())
	}

	if rotatedPoint.Z() != -2 {
		t.Errorf("Expected -2 but got %d.", rotatedPoint.Z())
	}
}

func Test_Point_RotateY(t *testing.T) {
	point := New3DPoint(1, 2, 3)
	rotatedPoint := point.RotateYClockwise()

	if rotatedPoint == nil {
		t.Error("rotatedPoint is nil.")
	}

	if rotatedPoint.X() != -3 {
		t.Errorf("Expected -3 but got %d.", rotatedPoint.X())
	}

	if rotatedPoint.Y() != 2 {
		t.Errorf("Expected 2 but got %d.", rotatedPoint.Y())
	}

	if rotatedPoint.Z() != 1 {
		t.Errorf("Expected 1 but got %d.", rotatedPoint.Z())
	}
}

func Test_Point_RotateZ(t *testing.T) {
	point := New3DPoint(1, 2, 3)
	rotatedPoint := point.RotateZClockwise()

	if rotatedPoint == nil {
		t.Error("rotatedPoint is nil.")
	}

	if rotatedPoint.X() != 2 {
		t.Errorf("Expected 2 but got %d.", rotatedPoint.X())
	}

	if rotatedPoint.Y() != -1 {
		t.Errorf("Expected -1 but got %d.", rotatedPoint.Y())
	}

	if rotatedPoint.Z() != 3 {
		t.Errorf("Expected 3 but got %d.", rotatedPoint.Z())
	}
}
