package day04

import "github.com/yarsiemanym/advent-of-code-2022/common"

type cleaningAssignmentPair struct {
	First  *cleaningAssignment
	Second *cleaningAssignment
}

func (thisPair *cleaningAssignmentPair) OneContainsTheOther() bool {
	return thisPair.First.Contains(thisPair.Second) || thisPair.Second.Contains(thisPair.First)
}

func (thisPair *cleaningAssignmentPair) HasOverlap() bool {
	return thisPair.First.Overlaps(thisPair.Second)
}

func parseCleaningAssignmentPair(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, ",")

	return &cleaningAssignmentPair{
		First:  parseCleaningAssignment(tokens[0]),
		Second: parseCleaningAssignment(tokens[1]),
	}
}
