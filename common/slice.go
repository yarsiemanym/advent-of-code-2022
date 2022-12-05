package common

func Contains[T comparable](slice []T, value T) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}

	return false
}

func ContainsPointer[T comparable](slice []*T, value *T) bool {
	for _, element := range slice {
		if *element == *value {
			return true
		}
	}

	return false
}

func Intersection[T comparable](slice1 []T, slice2 []T) []T {
	intersection := []T{}

	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				intersection = append(intersection, element1)
			}
		}
	}

	return intersection
}

func IntersectionPointers[T comparable](slice1 []*T, slice2 []*T) []*T {
	intersection := []*T{}

	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if *element1 == *element2 {
				intersection = append(intersection, element1)
			}
		}
	}

	return intersection
}

func Union[T comparable](slice1 []T, slice2 []T) []T {
	counts := map[T]int{}

	for _, element1 := range slice1 {
		counts[element1] = counts[element1] + 1
	}

	for _, element2 := range slice2 {
		counts[element2] = counts[element2] + 1
	}

	union := []T{}

	for key := range counts {
		union = append(union, key)
	}

	return union
}

func UnionPointers[T comparable](slice1 []*T, slice2 []*T) []*T {
	counts := map[T]int{}

	for _, element1 := range slice1 {
		counts[*element1] = counts[*element1] + 1
	}

	for _, element2 := range slice2 {
		counts[*element2] = counts[*element2] + 1
	}

	union := []*T{}

	for key := range counts {
		element := new(T)
		*element = key
		union = append(union, element)
	}

	return union
}

func Exclusive[T comparable](slice1 []T, slice2 []T) []T {
	counts := map[T]int{}

	for _, element1 := range slice1 {
		counts[element1] = counts[element1] + 1
	}

	for _, element2 := range slice2 {
		counts[element2] = counts[element2] + 1
	}

	exclusive := []T{}

	for key, value := range counts {
		if value == 1 {
			exclusive = append(exclusive, key)
		}
	}

	return exclusive
}

func ExclusivePointers[T comparable](slice1 []*T, slice2 []*T) []*T {
	counts := map[T]int{}

	for _, element1 := range slice1 {
		counts[*element1] = counts[*element1] + 1
	}

	for _, element2 := range slice2 {
		counts[*element2] = counts[*element2] + 1
	}

	exclusive := []*T{}

	for key, value := range counts {
		if value == 1 {
			element := new(T)
			*element = key
			exclusive = append(exclusive, element)
		}
	}

	return exclusive
}
