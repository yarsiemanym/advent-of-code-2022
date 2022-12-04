package day04

import (
	"fmt"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type cleaningAssignment struct {
	Start int
	End   int
}

func (thisAssignment *cleaningAssignment) String() string {
	return fmt.Sprintf("%d-%d", thisAssignment.Start, thisAssignment.End)
}

func (thisAssignment *cleaningAssignment) Contains(otherAssignment *cleaningAssignment) bool {
	log.Debugf("Checking if assignment %s contains %s.", thisAssignment, otherAssignment)

	contains := otherAssignment.End <= thisAssignment.End && otherAssignment.Start >= thisAssignment.Start

	if contains {
		log.Debugf("Assignment %s contains %s.", thisAssignment, otherAssignment)
	}

	return contains
}

func (thisAssignment *cleaningAssignment) Overlaps(otherAssignment *cleaningAssignment) bool {
	log.Debugf("Checking if assignments %s and %s overlap.", thisAssignment, otherAssignment)

	contains := otherAssignment.End >= thisAssignment.Start && otherAssignment.Start <= thisAssignment.End

	if contains {
		log.Debugf("Assignments %s and %s overlap.", thisAssignment, otherAssignment)
	}

	return contains
}

func parseCleaningAssignment(text string) *cleaningAssignment {
	if text == "" {
		return nil
	}

	tokens := common.Split(text, "-")

	start, err := strconv.Atoi(tokens[0])
	common.Check(err)

	end, err := strconv.Atoi(tokens[1])
	common.Check(err)

	return &cleaningAssignment{
		Start: start,
		End:   end,
	}
}
