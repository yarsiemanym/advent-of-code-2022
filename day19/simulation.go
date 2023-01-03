package day19

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type simulationState struct {
	blueprint           *blueprint
	time                int
	timeLimit           int
	activeRobots        map[material]int
	stockpiledMaterials map[material]int
}

func NewSimulation(blueprint *blueprint, timeLimit int) *simulationState {
	return &simulationState{
		blueprint: blueprint,
		time:      0,
		timeLimit: timeLimit,
		activeRobots: map[material]int{
			ore:      1,
			clay:     0,
			obsidian: 0,
			geode:    0,
		},
		stockpiledMaterials: map[material]int{
			ore:      0,
			clay:     0,
			obsidian: 0,
			geode:    0,
		},
	}
}

func (currentState *simulationState) Run() int {
	bestGeodes := 0
	currentState.run(&bestGeodes)
	return bestGeodes
}

func (currentState *simulationState) run(bestGeodes *int) {
	if currentState.stockpiledMaterials[geode] > *bestGeodes {
		*bestGeodes = currentState.stockpiledMaterials[geode]
		log.Debugf("New best geodes: %d", *bestGeodes)
		log.Debug(currentState)
	}

	if currentState.time >= currentState.timeLimit {
		return
	}

	if *bestGeodes >= currentState.stockpiledMaterials[geode]+common.TriangularSum(currentState.activeRobots[geode], currentState.activeRobots[geode]+currentState.TimeRemaining()) {
		return
	}

	noBuildState := currentState.Clone()
	noBuildState.Tick()
	noBuildState.CollectMaterials()
	nextStates := []*simulationState{noBuildState}

	for robotType := geode; robotType >= ore; robotType-- {

		if !currentState.CanBuild(robotType) {
			continue
		}

		if !currentState.ShouldBuild(robotType) {
			break
		}

		buildState := currentState.Clone()
		buildState.Tick()
		buildState.Build(robotType)
		buildState.CollectMaterials()
		buildState.ActivateRobot(robotType)
		nextStates = append(nextStates, buildState)
	}

	for _, nextState := range nextStates {
		nextState.run(bestGeodes)
	}
}

func (state *simulationState) CanBuild(robotType material) bool {

	robot := state.blueprint.robots[robotType]
	for material := ore; material <= geode; material++ {
		if state.stockpiledMaterials[material] < robot.requiredMaterials[material] {
			return false
		}
	}

	return true
}

func (state *simulationState) ShouldBuild(robotType material) bool {
	switch robotType {
	case ore:
		return state.activeRobots[ore] < state.blueprint.robots[ore].requiredMaterials[ore] ||
			state.activeRobots[ore] < state.blueprint.robots[clay].requiredMaterials[ore] ||
			state.activeRobots[ore] < state.blueprint.robots[obsidian].requiredMaterials[ore] ||
			state.activeRobots[ore] < state.blueprint.robots[geode].requiredMaterials[ore]
	case clay:
		return state.activeRobots[clay] < state.blueprint.robots[obsidian].requiredMaterials[clay]
	case obsidian:
		return state.activeRobots[obsidian] < state.blueprint.robots[geode].requiredMaterials[obsidian]
	case geode:
		return true
	default:
		return false
	}
}

func (state *simulationState) Build(robotType material) {

	robot := state.blueprint.robots[robotType]
	for material := ore; material <= geode; material++ {
		state.stockpiledMaterials[material] -= robot.requiredMaterials[material]
	}
}

func (state *simulationState) CollectMaterials() {
	for material := ore; material <= geode; material++ {
		state.stockpiledMaterials[material] += state.activeRobots[material]
	}
}

func (state *simulationState) Tick() {
	state.time++
}

func (state *simulationState) ActivateRobot(robotType material) {
	state.activeRobots[robotType]++
}

func (state *simulationState) TimeRemaining() int {
	return state.timeLimit - state.time
}

func (state *simulationState) Clone() *simulationState {
	return &simulationState{
		blueprint: state.blueprint,
		time:      state.time,
		timeLimit: state.timeLimit,
		activeRobots: map[material]int{
			ore:      state.activeRobots[ore],
			clay:     state.activeRobots[clay],
			obsidian: state.activeRobots[obsidian],
			geode:    state.activeRobots[geode],
		},
		stockpiledMaterials: map[material]int{
			ore:      state.stockpiledMaterials[ore],
			clay:     state.stockpiledMaterials[clay],
			obsidian: state.stockpiledMaterials[obsidian],
			geode:    state.stockpiledMaterials[geode],
		},
	}
}

func (state *simulationState) String() string {
	output := fmt.Sprintf("Time %d of %d\n", state.time, state.timeLimit)

	output += "--- Active Robots ---\n"
	for material := ore; material <= geode; material++ {
		output += fmt.Sprintf("%d %s\n", state.activeRobots[material], material)
	}

	output += "--- Stockpiled Materials ---\n"
	for material := ore; material <= geode; material++ {
		output += fmt.Sprintf("%d %s\n", state.stockpiledMaterials[material], material)
	}

	return output
}
