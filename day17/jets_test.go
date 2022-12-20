package day17

import (
	"testing"
)

func Test_Jets_Blast_1(t *testing.T) {
	pattern := ">><<"
	jets := parseJets(pattern)

	direction := jets.Blast()

	if direction != right {
		t.Errorf("Expected %c but got %c.", right, direction)
	} else if jets.counter != 1 {
		t.Errorf("Expected 1 but got %d.", jets.counter)
	}

	direction = jets.Blast()

	if direction != right {
		t.Errorf("Expected %c but got %c.", right, direction)
	} else if jets.counter != 2 {
		t.Errorf("Expected 2 but got %d.", jets.counter)
	}

	direction = jets.Blast()

	if direction != left {
		t.Errorf("Expected %c but got %c.", left, direction)
	} else if jets.counter != 3 {
		t.Errorf("Expected 3 but got %d.", jets.counter)
	}

	direction = jets.Blast()

	if direction != left {
		t.Errorf("Expected %c but got %c.", left, direction)
	} else if jets.counter != 4 {
		t.Errorf("Expected 4 but got %d.", jets.counter)
	}

	direction = jets.Blast()

	if direction != right {
		t.Errorf("Expected %c but got %c.", right, direction)
	} else if jets.counter != 5 {
		t.Errorf("Expected 5 but got %d.", jets.counter)
	}

	direction = jets.Blast()

	if direction != right {
		t.Errorf("Expected %c but got %c.", right, direction)
	} else if jets.counter != 6 {
		t.Errorf("Expected 6 but got %d.", jets.counter)
	}
}
