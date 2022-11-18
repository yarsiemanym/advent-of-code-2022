package common

import (
	"math"
	"sort"
)

func AbsInt(value int) int {
	return int(math.Abs(float64(value)))
}

func SumInt(values ...int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func AverageInt(values ...int) int {
	return SumInt(values...) / len(values)
}

func MedianInt(values ...int) int {
	length := len(values)
	median := 0

	if length > 0 {
		sort.Ints(values)

		if length%2 != 0 {
			median = values[length/2]
		} else {
			median = AverageInt(values[(length/2)-1], values[length/2])
		}
	}

	return median
}

func MaxInt(values ...int) int {
	maxValue := math.MinInt

	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

func MinInt(values ...int) int {
	minValue := math.MaxInt

	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}

	return minValue
}

func Reduce(numbers ...int) []int {
	if len(numbers) == 0 {
		return []int{}
	} else if len(numbers) == 1 {
		return numbers
	}

	gcf := numbers[0]

	for i := 1; i < len(numbers); i++ {
		gcf = GreatestCommonFactor(gcf, numbers[i])
	}

	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] / gcf
	}

	return numbers
}

func GreatestCommonFactor(a int, b int) int {

	if a < 0 {
		a = AbsInt(a)
	}

	if b < 0 {
		b = AbsInt(b)
	}

	if a == 0 {
		return b
	}

	if b == 0 {
		return a
	}

	if a > b {
		a, b = b, a
	}

	if b%a == 0 {
		return a
	} else if a > b {
		return GreatestCommonFactor(a-b, b)
	} else {
		return GreatestCommonFactor(a, b-a)
	}
}

func PowInt(base int, exponent int) int {
	return int(math.Pow(float64(base), float64(exponent)))
}
