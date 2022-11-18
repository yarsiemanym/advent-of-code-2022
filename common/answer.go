package common

import "time"

type Answer struct {
	Year          int
	Day           int
	Part1         string
	Part2         string
	ExecutionTime time.Duration
}
