package day19

import (
	"regexp"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type blueprint struct {
	id     int
	robots map[material]*robot
}

func parseBlueprint(text string) any {
	if text == "" {
		return nil
	}

	pattern := `Blueprint (\d+): Each ore robot costs (\d+) ore. Each clay robot costs (\d+) ore. Each obsidian robot costs (\d+) ore and (\d+) clay. Each geode robot costs (\d+) ore and (\d+) obsidian.`
	blueprintRegexp := regexp.MustCompile(pattern)
	matches := blueprintRegexp.FindStringSubmatch(text)

	id, err := strconv.Atoi(matches[1])
	common.Check(err)

	oreRobotCostOre, err := strconv.Atoi(matches[2])
	common.Check(err)

	clayRobotCostOre, err := strconv.Atoi(matches[3])
	common.Check(err)

	obsidianRobotCostOre, err := strconv.Atoi(matches[4])
	common.Check(err)

	obsidianRobotCostClay, err := strconv.Atoi(matches[5])
	common.Check(err)

	geodeRobotCostOre, err := strconv.Atoi(matches[6])
	common.Check(err)

	geodeRobotCostObsidian, err := strconv.Atoi(matches[7])
	common.Check(err)

	return &blueprint{
		id: id,
		robots: map[material]*robot{
			ore: {
				gathersMaterial: ore,
				requiredMaterials: map[material]int{
					ore:      oreRobotCostOre,
					clay:     0,
					obsidian: 0,
					geode:    0,
				},
			},
			clay: {
				gathersMaterial: clay,
				requiredMaterials: map[material]int{
					ore:      clayRobotCostOre,
					clay:     0,
					obsidian: 0,
					geode:    0,
				},
			},
			obsidian: {
				gathersMaterial: obsidian,
				requiredMaterials: map[material]int{
					ore:      obsidianRobotCostOre,
					clay:     obsidianRobotCostClay,
					obsidian: 0,
					geode:    0,
				},
			},
			geode: {
				gathersMaterial: geode,
				requiredMaterials: map[material]int{
					ore:      geodeRobotCostOre,
					clay:     0,
					obsidian: geodeRobotCostObsidian,
					geode:    0,
				},
			},
		},
	}
}
