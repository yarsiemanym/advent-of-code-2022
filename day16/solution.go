package day16

import (
	"fmt"
	"math"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseValve)

	valves := map[string]valve{}
	for _, result := range results {
		valve := *result.(*valve)
		valves[valve.label] = valve
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(valves),
		Part2: solvePart2(valves),
	}
}

func solvePart1(valves map[string]valve) string {
	log.Debug("Solving part 1.")

	valves["AA"] = valve{
		label:     valves["AA"].label,
		tunnelsTo: valves["AA"].tunnelsTo,
		flowRate:  valves["AA"].flowRate,
		isOpen:    valves["AA"].isOpen,
		isVisited: true,
	}

	start := &simulationState{
		label:             "AA",
		time:              0,
		timeLimit:         30,
		action:            "start",
		pressurePerMinute: 0,
		totalPressure:     0,
		valves:            valves,
	}

	goal := &simulationState{
		label:         "*",
		time:          30,
		action:        "stop",
		totalPressure: 0,
	}

	ceiling := math.MaxInt
	search := common.NewAStarSearch(start, goal, heuristicFunction, possibleNextStatesFunction, ceiling)
	states := search.Search()
	finalState := states[len(states)-2].(*simulationState)
	pressure := finalState.totalPressure

	spew.Dump(states)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(pressure)
}

func solvePart2(map[string]valve) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}

type simulationState struct {
	label             string
	time              int
	timeLimit         int
	action            string
	pressurePerMinute int
	totalPressure     int
	valves            map[string]valve
}

func (thisState *simulationState) Clone() *simulationState {
	clonedValves := map[string]valve{}

	for label, valve := range thisState.valves {
		clonedValves[label] = valve
	}

	return &simulationState{
		label:             thisState.label,
		valves:            clonedValves,
		time:              thisState.time,
		timeLimit:         thisState.timeLimit,
		action:            thisState.action,
		pressurePerMinute: thisState.pressurePerMinute,
		totalPressure:     thisState.totalPressure,
	}
}

func (thisState *simulationState) TimeRemaining() int {
	return thisState.timeLimit - thisState.time
}

func (thisState *simulationState) Valve() valve {
	return thisState.valves[thisState.label]
}

func (thisState *simulationState) SumOfUnopenedFlowRate() int {
	sumOfUnopenedFlowRate := 0

	for _, valve := range thisState.valves {
		if !valve.isOpen {
			sumOfUnopenedFlowRate += valve.flowRate
		}
	}

	return sumOfUnopenedFlowRate
}

func (thisState *simulationState) String() string {
	return fmt.Sprintf("(%d %s %s %d)", thisState.time, thisState.action, thisState.label, thisState.totalPressure)
}

func (thisState *simulationState) Cost() int {
	thisValve := thisState.Valve()
	cost := 0

	switch thisState.action {
	case "open":
		cost += (thisState.SumOfUnopenedFlowRate() * thisState.TimeRemaining()) - (thisValve.flowRate * thisState.time)
	case "move":
		cost += thisState.SumOfUnopenedFlowRate() * thisState.TimeRemaining()
	case "idle":

	default:

	}

	return cost
}

func (thisState *simulationState) Key() string {
	return thisState.String()
}

func heuristicFunction(current common.State, goal common.State) int {
	/* currentState := current.(*simulationState)
	currentValve := currentState.Valve()

	if currentValve.isVisited {
		return currentState.HighestUnopenedFlowRate() * currentState.time
	} else {
		return 0
	} */

	return 0
}

func possibleNextStatesFunction(current common.State) []common.State {
	currentState := current.(*simulationState)

	log.Debugf("Current state is %s.", currentState)
	log.Debugf("Pressure per minute is %d", currentState.pressurePerMinute)
	log.Debugf("Total pressure is %d.", currentState.totalPressure)
	log.Debugf("Constructing possible next states from %s.", current)

	currentValve := currentState.valves[currentState.label]
	nextStates := []common.State{}

	if currentState.time < currentState.timeLimit {

		allValvesAreOpen := true

		for _, valve := range currentState.valves {
			if valve.flowRate > 0 && !valve.isOpen {
				allValvesAreOpen = false
			}
		}

		// All valves are on. Sit idle till time is up.
		if allValvesAreOpen {
			nextState := currentState.Clone()
			nextState.time += 1
			nextState.action = "idle"
			nextState.totalPressure += nextState.pressurePerMinute
			nextStates = append(nextStates, nextState)
			log.Debugf("%s -> %s costs %d", currentState, nextState, nextState.Cost())
		}

		// Open current valve
		if currentValve.flowRate > 0 && !currentValve.isOpen {
			nextState := currentState.Clone()
			nextState.valves[currentValve.label] = valve{
				label:     currentValve.label,
				tunnelsTo: currentValve.tunnelsTo,
				flowRate:  currentValve.flowRate,
				isOpen:    true,
				isVisited: currentValve.isVisited,
			}
			nextState.time += 1
			nextState.action = "open"
			nextState.totalPressure += nextState.pressurePerMinute
			nextState.pressurePerMinute += currentValve.flowRate
			nextStates = append(nextStates, nextState)
			log.Debugf("%s -> %s costs %d", currentState, nextState, nextState.Cost())
		}

		// Move to an adjacent valve
		for _, adjacentLabel := range currentValve.tunnelsTo {
			adjacentValve := currentState.valves[adjacentLabel]

			nextState := currentState.Clone()
			nextState.valves[adjacentLabel] = valve{
				label:     adjacentValve.label,
				tunnelsTo: adjacentValve.tunnelsTo,
				flowRate:  adjacentValve.flowRate,
				isOpen:    adjacentValve.isOpen,
				isVisited: true,
			}
			nextState.time += 1
			nextState.label = adjacentLabel
			nextState.action = "move"
			nextState.totalPressure += nextState.pressurePerMinute
			nextStates = append(nextStates, nextState)
			log.Debugf("%s -> %s costs %d", currentState, nextState, nextState.Cost())
		}
	} else {
		// Time is up. Stop where you are.
		nextState := currentState.Clone()
		nextState.label = "*"
		nextState.action = "stop"
		nextState.totalPressure = 0
		nextStates = append(nextStates, nextState)
		log.Debugf("%s -> %s costs %d", currentState, nextState, nextState.Cost())
	}

	return nextStates
}
