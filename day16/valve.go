package day16

import (
	"regexp"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type valve struct {
	label     string
	index     int
	flowRate  int
	tunnelsTo []string
}

func parseValve(text string) any {
	if text == "" {
		return nil
	}

	valveRegexp := regexp.MustCompile(`Valve (\w+) has flow rate=(\d+); tunnels? leads? to valves? ([\w, ]+)`)
	matches := valveRegexp.FindStringSubmatch(text)

	label := matches[1]

	flowRate, err := strconv.Atoi(matches[2])
	common.Check(err)

	tunnelsTo := common.Split(matches[3], ", ")

	return &valve{
		label:     label,
		flowRate:  flowRate,
		tunnelsTo: tunnelsTo,
	}
}
