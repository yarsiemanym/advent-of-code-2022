package day15

import (
	"testing"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Test_Sensor_FindFrontier_1(t *testing.T) {
	sensor := &sensor{
		position:      common.New2DPoint(3, 3),
		closestBeacon: common.New2DPoint(3, 2),
	}

	/*
			 F
		    FBF
		   F#S#F
		    F#F
			 F
	*/

	edges := sensor.Frontier()

	if len(edges) != 4 {
		t.Errorf("Expected 4 but got %d.", len(edges))
	} else if edges[0].Start().X() != 1 {
		t.Errorf("Expected 3 but got %d.", edges[0].Start().X())
	} else if edges[0].Start().Y() != 3 {
		t.Errorf("Expected 1 but got %d.", edges[0].Start().Y())
	} else if edges[0].End().X() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[0].End().X())
	} else if edges[0].End().Y() != 1 {
		t.Errorf("Expected 1 but got %d.", edges[0].End().Y())
	} else if edges[1].Start().X() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[1].Start().X())
	} else if edges[1].Start().Y() != 1 {
		t.Errorf("Expected 1 but got %d.", edges[1].Start().Y())
	} else if edges[1].End().X() != 5 {
		t.Errorf("Expected 5 but got %d.", edges[1].End().X())
	} else if edges[1].End().Y() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[1].End().Y())
	} else if edges[2].Start().X() != 5 {
		t.Errorf("Expected 5 but got %d.", edges[2].Start().X())
	} else if edges[2].Start().Y() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[2].Start().Y())
	} else if edges[2].End().X() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[2].End().X())
	} else if edges[2].End().Y() != 5 {
		t.Errorf("Expected 5 but got %d.", edges[2].End().Y())
	} else if edges[3].Start().X() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[3].Start().X())
	} else if edges[3].Start().Y() != 5 {
		t.Errorf("Expected 5 but got %d.", edges[3].Start().Y())
	} else if edges[3].End().X() != 1 {
		t.Errorf("Expected 1 but got %d.", edges[3].End().X())
	} else if edges[3].End().Y() != 3 {
		t.Errorf("Expected 3 but got %d.", edges[3].End().Y())
	}
}
