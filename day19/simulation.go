package day19

import (
	"fmt"
	"math"

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
		log.Info(currentState)
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

	if len(nextStates) == 0 {
		noBuildState := currentState.Clone()
		noBuildState.Tick()
		noBuildState.CollectMaterials()
		nextStates = append(nextStates, noBuildState)
	}

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

	timeToOreCrtiticalMass := common.MaxInt(state.TimeToOreCriticalMass(ore),
		state.TimeToOreCriticalMass(clay),
		state.TimeToOreCriticalMass(obsidian),
		state.TimeToOreCriticalMass(geode))

	switch robotType {
	case ore:
		shouldBuild = timeToOreCrtiticalMass > 0 && state.TimeRemaining() > 4
	case clay:
		shouldBuild = state.TimeToClayCriticalMass() > timeToOreCrtiticalMass && state.TimeRemaining() > 6
	case obsidian:
		shouldBuild = state.TimeToObsidianCriticalMass() > timeToOreCrtiticalMass && state.TimeRemaining() > 4
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

func (state *simulationState) TimeToOreCriticalMass(robotType material) int {
	oreDemand := state.blueprint.robots[robotType].requiredMaterials[ore]
	oreDeficit := oreDemand - state.stockpiledMaterials[ore]
	timeToCiticalMass := math.MaxInt
	robotCount := state.activeRobots[ore]
	if robotCount > 0 {
		timeToCiticalMass = oreDeficit / robotCount

		if timeToCiticalMass%robotCount != 0 {
			timeToCiticalMass++
		}
	}

	return common.MaxInt(timeToCiticalMass, 0)
}

func (state *simulationState) TimeToClayCriticalMass() int {
	clayDemand := state.blueprint.robots[obsidian].requiredMaterials[clay]
	clayDeficit := clayDemand - state.stockpiledMaterials[clay]
	timeToCiticalMass := math.MaxInt
	robotCount := state.activeRobots[clay]
	if robotCount > 0 {
		timeToCiticalMass = clayDeficit / robotCount

		if timeToCiticalMass%robotCount != 0 {
			timeToCiticalMass++
		}
	}

	return common.MaxInt(timeToCiticalMass, 0)
}

func (state *simulationState) TimeToObsidianCriticalMass() int {
	obsidianDemand := state.blueprint.robots[geode].requiredMaterials[obsidian]
	obsidianDeficit := obsidianDemand - state.stockpiledMaterials[obsidian]
	timeToCiticalMass := math.MaxInt
	robotCount := state.activeRobots[obsidian]
	if robotCount > 0 {
		timeToCiticalMass = obsidianDeficit / robotCount

		if timeToCiticalMass%robotCount != 0 {
			timeToCiticalMass++
		}
	}

	return common.MaxInt(timeToCiticalMass, 0)
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
