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

		(3,1) (2,2) (4,2) (1,3) (5,3) (2,4) (4,4) (3,5)
	*/

	frontier := sensor.Frontier()

	if len(frontier) != 8 {
		t.Errorf("Expected 8 but got %d.", len(frontier))
	} else if frontier[0].X() != 3 {
		t.Errorf("Expected 3 but got %d.", frontier[0].X())
	} else if frontier[0].Y() != 1 {
		t.Errorf("Expected 1 but got %d.", frontier[0].Y())
	} else if frontier[1].X() != 2 {
		t.Errorf("Expected 2 but got %d.", frontier[1].X())
	} else if frontier[1].Y() != 2 {
		t.Errorf("Expected 2 but got %d.", frontier[1].Y())
	} else if frontier[2].X() != 4 {
		t.Errorf("Expected 4 but got %d.", frontier[2].X())
	} else if frontier[2].Y() != 2 {
		t.Errorf("Expected 2 but got %d.", frontier[2].Y())
	} else if frontier[3].X() != 1 {
		t.Errorf("Expected 1 but got %d.", frontier[3].X())
	} else if frontier[3].Y() != 3 {
		t.Errorf("Expected 3 but got %d.", frontier[3].Y())
	} else if frontier[4].X() != 5 {
		t.Errorf("Expected 5 but got %d.", frontier[4].X())
	} else if frontier[4].Y() != 3 {
		t.Errorf("Expected 3 but got %d.", frontier[4].Y())
	} else if frontier[5].X() != 2 {
		t.Errorf("Expected 2 but got %d.", frontier[5].X())
	} else if frontier[5].Y() != 4 {
		t.Errorf("Expected 4 but got %d.", frontier[5].Y())
	} else if frontier[6].X() != 4 {
		t.Errorf("Expected 4 but got %d.", frontier[6].X())
	} else if frontier[6].Y() != 4 {
		t.Errorf("Expected 4 but got %d.", frontier[6].Y())
	} else if frontier[7].X() != 3 {
		t.Errorf("Expected 3 but got %d.", frontier[7].X())
	} else if frontier[7].Y() != 5 {
		t.Errorf("Expected 5 but got %d.", frontier[7].Y())
	}
}
