package common

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
