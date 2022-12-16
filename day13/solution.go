package day13

import (
	"sort"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n\n", parsePacketPair)

	var packetPairs []*packetPair
	for _, result := range results {
		packetPairs = append(packetPairs, result.(*packetPair))
	}

	results = common.ParseFile(puzzle.InputFile, "\n", parsePacket)

	var packets []*packet
	for _, result := range results {
		packets = append(packets, result.(*packet))
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(packetPairs),
		Part2: solvePart2(packets),
	}
}

func solvePart1(packetPairs []*packetPair) string {
	log.Debug("Solving part 1.")

	sumOfIndices := 0

	for index, packetPair := range packetPairs {
		if packetPair.IsOrderedCorrectly() {
			sumOfIndices += index + 1
		}
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(sumOfIndices)
}

func solvePart2(packets []*packet) string {
	log.Debug("Solving part 2.")

	dividerPackets := []*packet{
		parsePacket("[[2]]").(*packet),
		parsePacket("[[6]]").(*packet),
	}

	packets = append(packets, dividerPackets...)
	sort.Slice(packets, func(first int, second int) bool { return packets[first].IsLessThan(packets[second]) })
	indices := findPackets(packets, dividerPackets)
	product := (indices[0] + 1) * (indices[1] + 1)

	log.Debug("Part 2 solved.")
	return strconv.Itoa(product)
}
