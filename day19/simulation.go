package day19

import (
	"fmt"

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
		time:      1,
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
	for robotType := geode; robotType >= ore; robotType-- {
		currentState.DepthFirstSearch(robotType, &bestGeodes)
	}
	return bestGeodes
}

func (currentState simulationState) DepthFirstSearch(nextRobotType material, bestGeodes *int) {

	if !currentState.ShouldBuild(nextRobotType) {
		return
	}

	if currentState.IsFutile(bestGeodes) {
		return
	}

	for currentState.TimeRemaining() >= 0 {

		if currentState.CanBuild(nextRobotType) {
			nextState := currentState.Build(nextRobotType).CollectMaterials().ActivateRobot(nextRobotType).Tick()

			for robotType := geode; robotType >= ore; robotType-- {
				nextState.DepthFirstSearch(robotType, bestGeodes)
			}

			return
		} else {
			currentState = currentState.CollectMaterials().Tick()
		}
	}

	*bestGeodes = common.MaxInt(*bestGeodes, currentState.stockpiledMaterials[geode])
}

func (state simulationState) CanBuild(robotType material) bool {

	robot := state.blueprint.robots[robotType]
	for material := ore; material <= geode; material++ {
		if state.stockpiledMaterials[material] < robot.requiredMaterials[material] {
			return false
		}
	}
	return true
}

func (state simulationState) ShouldBuild(robotType material) bool {
	shouldBuild := false

	switch robotType {
	case ore:
		oreStockpile := state.stockpiledMaterials[ore]
		oreProduction := state.activeRobots[ore]
		oreDemand := common.MaxInt(state.blueprint.robots[clay].requiredMaterials[ore],
			state.blueprint.robots[obsidian].requiredMaterials[ore],
			state.blueprint.robots[geode].requiredMaterials[ore])
		shouldBuild = oreProduction < oreDemand && (oreProduction*state.TimeRemaining())+oreStockpile < state.TimeRemaining()*oreDemand
	case clay:
		clayStockpile := state.stockpiledMaterials[clay]
		clayProduction := state.activeRobots[clay]
		clayDemand := state.blueprint.robots[obsidian].requiredMaterials[clay]
		shouldBuild = clayProduction < clayDemand && (clayProduction*state.TimeRemaining())+clayStockpile < state.TimeRemaining()*clayDemand
	case obsidian:
		obsidianStockpile := state.stockpiledMaterials[obsidian]
		obsidianProduction := state.activeRobots[obsidian]
		obsidianDemand := state.blueprint.robots[geode].requiredMaterials[obsidian]
		shouldBuild = obsidianProduction < obsidianDemand && (obsidianProduction*state.TimeRemaining())+obsidianStockpile < state.TimeRemaining()*obsidianDemand
	case geode:
		shouldBuild = true
	default:
		shouldBuild = false
	}

	return shouldBuild
}

func (state simulationState) IsFutile(bestGeodes *int) bool {
	return *bestGeodes >= state.stockpiledMaterials[geode]+common.TrapezoidalSum(state.activeRobots[geode], state.activeRobots[geode]+state.TimeRemaining()+1)
}

func (state simulationState) Build(robotType material) simulationState {
	newState := state.Clone()
	robot := state.blueprint.robots[robotType]
	for material := ore; material <= geode; material++ {
		newState.stockpiledMaterials[material] -= robot.requiredMaterials[material]
	}
	return newState
}

func (state simulationState) CollectMaterials() simulationState {
	newState := state.Clone()
	for material := ore; material <= geode; material++ {
		newState.stockpiledMaterials[material] += state.activeRobots[material]
	}
	return newState
}

func (state simulationState) Tick() simulationState {
	newState := state.Clone()
	newState.time++
	return newState
}

func (state simulationState) ActivateRobot(robotType material) simulationState {
	newState := state.Clone()
	newState.activeRobots[robotType]++
	return newState
}

func (state *simulationState) TimeRemaining() int {
	return state.timeLimit - state.time
}

func (state simulationState) Clone() simulationState {
	return simulationState{
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

func (state simulationState) String() string {
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
