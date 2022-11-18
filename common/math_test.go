package common

import "testing"

func Test_SumInt(t *testing.T) {
	sum := SumInt(1, 2, 3, 4)

	if sum != 10 {
		t.Errorf("Expected 10 but got %d.", sum)
	}
}

func Test_MaxInt(t *testing.T) {
	max := MaxInt(3, 1, 4, 2)

	if max != 4 {
		t.Errorf("Expected 4 but got %d.", max)
	}
}

func Test_MinInt(t *testing.T) {
	min := MinInt(3, 1, 4, 2)

	if min != 1 {
		t.Errorf("Expected 1 but got %d.", min)
	}
}

func Test_GreatestCommonFactor_NotOne(t *testing.T) {
	gcf := GreatestCommonFactor(15, 6)

	if gcf != 3 {
		t.Errorf("Expected 3 but got %d.", gcf)
	}
}

func Test_GreatestCommonFactor_One(t *testing.T) {
	gcf := GreatestCommonFactor(19, 7)

	if gcf != 1 {
		t.Errorf("Expected 1 but got %d.", gcf)
	}
}

func Test_GreatestCommonFactor_Negative(t *testing.T) {
	gcf := GreatestCommonFactor(-15, 6)

	if gcf != 3 {
		t.Errorf("Expected 3 but got %d.", gcf)
	}
}

func Test_GreatestCommonFactor_Zero(t *testing.T) {
	gcf := GreatestCommonFactor(0, 5)

	if gcf != 5 {
		t.Errorf("Expected 5 but got %d.", gcf)
	}
}

func Test_Reduce(t *testing.T) {
	reducedNumbers := Reduce(-15, 6)

	if len(reducedNumbers) != 2 {
		t.Errorf("Expected 2 but got %d.", len(reducedNumbers))
	}

	if reducedNumbers[0] != -5 {
		t.Errorf("Expected -5 but got %d.", reducedNumbers[0])
	}

	if reducedNumbers[1] != 2 {
		t.Errorf("Expected 2 but got %d.", reducedNumbers[1])
	}
}

func Test_PowInt(t *testing.T) {
	result := PowInt(2, 3)

	if result != 8 {
		t.Errorf("Expected 8 but got %d.", result)
	}
}

func Test_MedianInt_Empty(t *testing.T) {
	median := MedianInt([]int{}...)

	if median != 0 {
		t.Errorf("Expected 0 but got %d.", median)
	}
}

func Test_MedianInt_Odd(t *testing.T) {
	values := []int{1, 2, 3}
	median := MedianInt(values...)

	if median != 2 {
		t.Errorf("Expected 2 but got %d.", median)
	}
}

func Test_MedianInt_Even(t *testing.T) {
	values := []int{1, 3, 5, 7}
	median := MedianInt(values...)

	if median != 4 {
		t.Errorf("Expected 4 but got %d.", median)
	}
}
