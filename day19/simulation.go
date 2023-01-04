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
	currentState.DepthFirstSearch(&bestGeodes)
	return bestGeodes
}

func (currentState *simulationState) DepthFirstSearch(bestGeodes *int) {
	if currentState.TimeRemaining() <= 0 {
		return
	}

	if currentState.IsFutile(bestGeodes) {
		return
	}

	nextStates := []*simulationState{}

	for robotType := geode; robotType >= ore; robotType-- {
		buildState := currentState.Clone()
		buildState.Tick()

		if !buildState.CanBuild(robotType) || !buildState.ShouldBuild(robotType) {
			continue
		}

		buildState.Build(robotType)
		buildState.CollectMaterials()
		buildState.ActivateRobot(robotType)
		nextStates = append(nextStates, buildState)

		log.Debugf("T=%d, Building a(n) %s robot.", buildState.time, robotType)
	}

	noBuildState := currentState.Clone()
	noBuildState.Tick()
	noBuildState.CollectMaterials()
	nextStates = append(nextStates, noBuildState)

	for _, nextState := range nextStates {
		*bestGeodes = common.MaxInt(*bestGeodes, nextState.stockpiledMaterials[geode])
		nextState.DepthFirstSearch(bestGeodes)
	}
}

func (state *simulationState) CanBuild(robotType material) bool {

	robot := state.blueprint.robots[robotType]
	for material := ore; material <= geode; material++ {
		if state.stockpiledMaterials[material] < robot.requiredMaterials[material] {
			log.Debugf("T=%d, Can not build a(n) %s robot.", state.time, robotType)
			return false
		}
	}

	log.Debugf("T=%d, Can build a(n) %s robot.", state.time, robotType)
	return true
}

func (state *simulationState) ShouldBuild(robotType material) bool {
	shouldBuild := false

	switch robotType {
	case ore:
		oreStockpile := state.stockpiledMaterials[ore]
		oreProduction := state.activeRobots[ore]
		oreDemand := common.MaxInt( // state.blueprint.robots[ore].requiredMaterials[ore],
			state.blueprint.robots[clay].requiredMaterials[ore],
			state.blueprint.robots[obsidian].requiredMaterials[ore],
			state.blueprint.robots[geode].requiredMaterials[ore])
		shouldBuild = oreProduction < oreDemand && oreStockpile < oreDemand
	case clay:
		clayStockpile := state.stockpiledMaterials[clay]
		clayProduction := state.activeRobots[clay]
		clayDemand := state.blueprint.robots[obsidian].requiredMaterials[clay]
		shouldBuild = clayProduction < clayDemand && clayStockpile < clayDemand
	case obsidian:
		obsidianStockpile := state.stockpiledMaterials[obsidian]
		obsidianProduction := state.activeRobots[obsidian]
		obsidianDemand := state.blueprint.robots[geode].requiredMaterials[obsidian]
		shouldBuild = obsidianProduction < obsidianDemand && obsidianStockpile < obsidianDemand
	case geode:
		shouldBuild = true
	default:
		shouldBuild = false
	}

	if shouldBuild {
		log.Debugf("T=%d, Should build a(n) %s robot.", state.time, robotType)
	} else {
		log.Debugf("T=%d, Should not build a(n) %s robot.", state.time, robotType)
	}

	return shouldBuild
}

func (state *simulationState) IsFutile(bestGeodes *int) bool {
	return *bestGeodes >= state.stockpiledMaterials[geode]+common.TrapezoidalSum(state.activeRobots[geode], state.activeRobots[geode]+state.TimeRemaining())
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
