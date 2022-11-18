package common

import "testing"

func Test_LineSegment_Slope_2D_Poistive(t *testing.T) {
	line := LineSegment{
		start: New2DPoint(1, 2),
		end:   New2DPoint(5, 4),
	}

	slope := line.Slope()

	if slope.x != 2 {
		t.Errorf("Expected 2 but got %v.", slope.x)
	}

	if slope.y != 1 {
		t.Errorf("Expected 1 but got %v.", slope.y)
	}
}

func Test_LineSegment_Slope_2D_Negative(t *testing.T) {
	line := LineSegment{
		start: New2DPoint(5, 4),
		end:   New2DPoint(1, 2),
	}

	slope := line.Slope()

	if slope.x != -2 {
		t.Errorf("Expected -2 but got %v.", slope.x)
	}

	if slope.y != -1 {
		t.Errorf("Expected -1 but got %v.", slope.y)
	}
}

func Test_LineSegment_Slope_3D_Positive(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 1,
			y: 2,
			z: 3,
		},
		end: &Point{
			x: 13,
			y: 10,
			z: 9,
		},
	}

	slope := line.Slope()

	if slope.x != 6 {
		t.Errorf("Expected 6 but got %v.", slope.x)
	}

	if slope.y != 4 {
		t.Errorf("Expected 4 but got %v.", slope.y)
	}

	if slope.z != 3 {
		t.Errorf("Expected 3 but got %v.", slope.z)
	}
}

func Test_LineSegment_Slope_3D_Negative(t *testing.T) {
	line := LineSegment{
		start: &Point{
			x: 13,
			y: 10,
			z: 9,
		},
		end: &Point{
			x: 1,
			y: 2,
			z: 3,
		},
	}

	slope := line.Slope()

	if slope.x != -6 {
		t.Errorf("Expected -6 but got %v.", slope.x)
	}

	if slope.y != -4 {
		t.Errorf("Expected -4 but got %v.", slope.y)
	}

	if slope.z != -3 {
		t.Errorf("Expected -3 but got %v.", slope.z)
	}
}
