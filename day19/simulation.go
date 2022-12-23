package day19

import (
	log "github.com/sirupsen/logrus"
)

type simulation struct {
	blueprint    *blueprint
	activeRobots map[material]int
	stockpile    map[material]int
	time         int
	timeLimit    int
}

func (simulation *simulation) Run() int {

	log.Debugf("Simulating blueprint %d.", simulation.blueprint.id)

	for ; simulation.time <= simulation.timeLimit; simulation.time++ {
		log.Debugf("Minute %d", simulation.time)
		newRobot, newRobotType := simulation.TryCreateRobot()

		simulation.CollectResources()

		if newRobot {
			simulation.ActivateNewRobot(newRobotType)
		}
	}

	return simulation.stockpile[geode]
}

func (simulation *simulation) TimeRemaining() int {
	return simulation.timeLimit - simulation.time
}

func (simulation *simulation) TryCreateRobot() (bool, material) {
	for robotType := geode; robotType >= ore; robotType-- {
		robot := simulation.blueprint.robots[robotType]
		if simulation.CanBuild(robot) && simulation.ShouldBuild(robot) {
			log.Debugf("Starting to create a %s robot.", robotType)

			for material, quantity := range robot.requiredMaterials {
				simulation.stockpile[material] -= quantity
			}

			return true, robotType
		}
	}

	return false, none
}

func (simulation *simulation) CanBuild(robot *robot) bool {
	canBuild := true
	for material, quantity := range robot.requiredMaterials {
		if simulation.stockpile[material] < quantity {
			canBuild = false
			break
		}
	}

	return canBuild
}

func (simulation *simulation) ShouldBuild(robot *robot) bool {
	// TODO
	return false
}

func (simulation *simulation) CollectResources() {
	for material := ore; material <= geode; material++ {
		quantity := simulation.activeRobots[material]
		simulation.stockpile[material] += quantity
		log.Debugf("Added %d %s to the stockpile. There are now %d.", quantity, material, simulation.stockpile[material])
	}
}

func (simulation *simulation) ActivateNewRobot(robotType material) {
	simulation.activeRobots[robotType]++
	log.Debugf("Activating a new %s robot. There are now %d.", robotType, simulation.activeRobots[robotType])
}

func NewSimulation(blueprint *blueprint, timeLimit int) *simulation {
	return &simulation{
		blueprint: blueprint,
		activeRobots: map[material]int{
			ore:      1,
			clay:     0,
			obsidian: 0,
			geode:    0,
		},
		stockpile: map[material]int{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geode:    0,
		},
		time:      1,
		timeLimit: timeLimit,
	}
}
