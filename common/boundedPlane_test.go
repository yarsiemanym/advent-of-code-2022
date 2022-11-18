package common

import "testing"

func Test_NewBoundedPlaneFromPoints(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	if plane == nil {
		t.Error("plane is nil")
	} else if plane.span == nil {
		t.Error("plane.span is nil")
	} else if plane.span.start == nil {
		t.Error("plane.span.start is nil")
	} else if plane.span.start.x != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.x)
	} else if plane.span.start.y != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.y)
	} else if plane.span.end.x != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.x)
	} else if plane.span.end.y != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.y)
	}

	if len(plane.locations) != 7 {
		t.Errorf("Expected 7 but got %v.", len(plane.locations))
	}

	for _, row := range plane.locations {
		if len(row) != 7 {
			t.Errorf("Expected 7 but got %v.", len(row))
		}
	}
}

func Test_NewBoundedPlaneFromLines(t *testing.T) {
	lines := []*LineSegment{
		NewLineSegment(New2DPoint(3, 1), New2DPoint(1, -3)),
		NewLineSegment(New2DPoint(-3, -1), New2DPoint(-1, 3)),
	}

	plane := NewBoundedPlaneFromLines(lines)

	if plane == nil {
		t.Error("plane is nil")
	} else if plane.span == nil {
		t.Error("plane.span is nil")
	} else if plane.span.start == nil {
		t.Error("plane.span.start is nil")
	} else if plane.span.start.x != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.x)
	} else if plane.span.start.y != -3 {
		t.Errorf("Expected -3 but got %v.", plane.span.start.y)
	} else if plane.span.end.x != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.x)
	} else if plane.span.end.y != 3 {
		t.Errorf("Expected 3 but got %v.", plane.span.end.y)
	}

	if len(plane.locations) != 7 {
		t.Errorf("Expected 7 but got %v.", len(plane.locations))
	}

	for _, row := range plane.locations {
		if len(row) != 7 {
			t.Errorf("Expected 7 but got %v.", len(row))
		}
	}
}

func Test_boundedPlane_GetValueAt(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)
	plane.locations[1][1] = true

	value1 := plane.GetValueAt(New2DPoint(-2, -2)).(bool)
	value2 := plane.GetValueAt(New2DPoint(1, 1))

	if !value1 {
		t.Errorf("Expected true but got %v.", value1)
	}

	if value2 != nil {
		t.Error("value2 is not nil.")
	}
}

func Test_boundedPlane_SetValueAt(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)
	plane.SetValueAt(New2DPoint(-2, -2), true)

	value1 := plane.locations[1][1].(bool)
	value2 := plane.GetValueAt(New2DPoint(1, 1))

	if !value1 {
		t.Errorf("Expected true but got %v.", value1)
	}

	if value2 != nil {
		t.Error("value2 is not nil.")
	}
}

func Test_boundedPlane_GetVonNeumannNeighbors_Interior(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetVonNeumannNeighbors(New2DPoint(1, 1))

	if len(adjacentPoints) != 4 {
		t.Errorf("Expected 4 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].y)
	}

	if adjacentPoints[2] == nil {
		t.Error("adjacenPoints[2] is nil.")
	} else if adjacentPoints[2].x != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[2].x)
	} else if adjacentPoints[2].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[2].y)
	}

	if adjacentPoints[3] == nil {
		t.Error("adjacenPoints[3] is nil.")
	} else if adjacentPoints[3].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[3].x)
	} else if adjacentPoints[3].y != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[3].y)
	}
}

func Test_boundedPlane_GetVonNeumannNeighbors_Corner1(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetVonNeumannNeighbors(New2DPoint(-3, -3))

	if len(adjacentPoints) != 2 {
		t.Errorf("Expected 2 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != -3 {
		t.Errorf("Expected -3 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != -3 {
		t.Errorf("Expected -3 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[1].y)
	}
}

func Test_boundedPlane_GetVonNeumannNeighbors_Corner2(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetVonNeumannNeighbors(New2DPoint(3, 3))

	if len(adjacentPoints) != 2 {
		t.Errorf("Expected 2 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != 3 {
		t.Errorf("Expected 3 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != 3 {
		t.Errorf("Expected 3 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].y)
	}
}

func Test_boundedPlane_GetMooreNeighbors_Interior(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetMooreNeighbors(New2DPoint(1, 1))

	if len(adjacentPoints) != 8 {
		t.Errorf("Expected 8 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].y)
	}

	if adjacentPoints[2] == nil {
		t.Error("adjacenPoints[2] is nil.")
	} else if adjacentPoints[2].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[2].x)
	} else if adjacentPoints[2].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[2].y)
	}

	if adjacentPoints[3] == nil {
		t.Error("adjacenPoints[3] is nil.")
	} else if adjacentPoints[3].x != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[3].x)
	} else if adjacentPoints[3].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[3].y)
	}

	if adjacentPoints[4] == nil {
		t.Error("adjacenPoints[4] is nil.")
	} else if adjacentPoints[4].x != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[4].x)
	} else if adjacentPoints[4].y != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[4].y)
	}

	if adjacentPoints[5] == nil {
		t.Error("adjacenPoints[5] is nil.")
	} else if adjacentPoints[5].x != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[5].x)
	} else if adjacentPoints[5].y != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[5].y)
	}

	if adjacentPoints[6] == nil {
		t.Error("adjacenPoints[6] is nil.")
	} else if adjacentPoints[6].x != 1 {
		t.Errorf("Expected 1 but got %v.", adjacentPoints[6].x)
	} else if adjacentPoints[6].y != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[6].y)
	}

	if adjacentPoints[7] == nil {
		t.Error("adjacenPoints[7] is nil.")
	} else if adjacentPoints[7].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[7].x)
	} else if adjacentPoints[7].y != 0 {
		t.Errorf("Expected 0 but got %v.", adjacentPoints[7].y)
	}
}

func Test_boundedPlane_GetMooreNeighbors_Corner1(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetMooreNeighbors(New2DPoint(-3, -3))

	if len(adjacentPoints) != 3 {
		t.Errorf("Expected 3 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != -3 {
		t.Errorf("Expected -3 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[1].y)
	}

	if adjacentPoints[2] == nil {
		t.Error("adjacenPoints[2] is nil.")
	} else if adjacentPoints[2].x != -3 {
		t.Errorf("Expected -3 but got %v.", adjacentPoints[2].x)
	} else if adjacentPoints[2].y != -2 {
		t.Errorf("Expected -2 but got %v.", adjacentPoints[2].y)
	}
}

func Test_boundedPlane_GetMooreNeighbors_Corner2(t *testing.T) {
	points := []*Point{
		New2DPoint(3, 3),
		New2DPoint(-3, 3),
		New2DPoint(-3, -3),
		New2DPoint(3, -3),
	}

	plane := NewBoundedPlaneFromPoints(points)

	adjacentPoints := plane.GetMooreNeighbors(New2DPoint(3, 3))

	if len(adjacentPoints) != 3 {
		t.Errorf("Expected 3 but got %v.", len(adjacentPoints))
	}

	if adjacentPoints[0] == nil {
		t.Error("adjacenPoints[0] is nil.")
	} else if adjacentPoints[0].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[0].x)
	} else if adjacentPoints[0].y != 3 {
		t.Errorf("Expected 3 but got %v.", adjacentPoints[0].y)
	}

	if adjacentPoints[1] == nil {
		t.Error("adjacenPoints[1] is nil.")
	} else if adjacentPoints[1].x != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].x)
	} else if adjacentPoints[1].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[1].y)
	}

	if adjacentPoints[2] == nil {
		t.Error("adjacenPoints[2] is nil.")
	} else if adjacentPoints[2].x != 3 {
		t.Errorf("Expected 3 but got %v.", adjacentPoints[2].x)
	} else if adjacentPoints[2].y != 2 {
		t.Errorf("Expected 2 but got %v.", adjacentPoints[2].y)
	}
}
