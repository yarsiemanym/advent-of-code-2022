package day17

import (
	"testing"
)

func Test_Rock_Generate_1(t *testing.T) {
	generator := rockGenerator{
		counter: 0,
	}

	rock := generator.NextRock()

	if len(rock.fallingBits) != 1 {
		t.Errorf("Expected 1 but got %d.", len(rock.fallingBits))
	} else if len(rock.fallingBits[0]) != 4 {
		t.Errorf("Expected 4 but got %d.", len(rock.fallingBits))
	} else if rock.fallingBits[0][0] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][0])
	} else if rock.fallingBits[0][1] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][1])
	} else if rock.fallingBits[0][2] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][2])
	} else if rock.fallingBits[0][3] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][3])
	} else if generator.counter != 1 {
		t.Errorf("Expected 1 but got %d.", generator.counter)
	}
}

func Test_Rock_Generate_2(t *testing.T) {
	generator := rockGenerator{
		counter: 1,
	}

	rock := generator.NextRock()

	if len(rock.fallingBits) != 3 {
		t.Errorf("Expected 3 but got %d.", len(rock.fallingBits))
	} else if len(rock.fallingBits[0]) != 3 {
		t.Errorf("Expected 3 but got %d.", len(rock.fallingBits[0]))
	} else if rock.fallingBits[0][0] != emptySpace {
		t.Errorf("Expected %d but got %c.", emptySpace, rock.fallingBits[0][0])
	} else if rock.fallingBits[0][1] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][1])
	} else if rock.fallingBits[0][2] != emptySpace {
		t.Errorf("Expected %d but got %c.", emptySpace, rock.fallingBits[0][2])
	} else if len(rock.fallingBits[1]) != 3 {
		t.Errorf("Expected 3 but got %d.", len(rock.fallingBits[1]))
	} else if rock.fallingBits[1][0] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[1][0])
	} else if rock.fallingBits[1][1] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[1][1])
	} else if rock.fallingBits[1][2] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[1][2])
	} else if len(rock.fallingBits[2]) != 3 {
		t.Errorf("Expected 3 but got %d.", len(rock.fallingBits[2]))
	} else if rock.fallingBits[2][0] != emptySpace {
		t.Errorf("Expected %d but got %c.", emptySpace, rock.fallingBits[2][0])
	} else if rock.fallingBits[2][1] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[2][1])
	} else if rock.fallingBits[2][2] != emptySpace {
		t.Errorf("Expected %d but got %c.", emptySpace, rock.fallingBits[2][2])
	} else if generator.counter != 2 {
		t.Errorf("Expected 2 but got %d.", generator.counter)
	}
}

func Test_Rock_Generate_3(t *testing.T) {
	generator := rockGenerator{
		counter: 5,
	}

	rock := generator.NextRock()

	if len(rock.fallingBits) != 1 {
		t.Errorf("Expected 1 but got %d.", len(rock.fallingBits))
	} else if len(rock.fallingBits[0]) != 4 {
		t.Errorf("Expected 4 but got %d.", len(rock.fallingBits))
	} else if rock.fallingBits[0][0] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][0])
	} else if rock.fallingBits[0][1] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][1])
	} else if rock.fallingBits[0][2] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][2])
	} else if rock.fallingBits[0][3] != fallingRock {
		t.Errorf("Expected %d but got %c.", fallingRock, rock.fallingBits[0][3])
	} else if generator.counter != 6 {
		t.Errorf("Expected 6 but got %d.", generator.counter)
	}
}
