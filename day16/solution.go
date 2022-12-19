package day16

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

func Solve(puzzle *common.Puzzle) common.Answer {

	results := common.ParseFile(puzzle.InputFile, "\n", parseValve)

	valves := map[string]*valve{}
	for index, result := range results {
		valve := result.(*valve)
		valve.index = index
		valves[valve.label] = valve
	}

	return common.Answer{
		Year:  puzzle.Year,
		Day:   puzzle.Day,
		Part1: solvePart1(valves),
		Part2: solvePart2(valves),
	}
}

func solvePart1(valves map[string]*valve) string {
	log.Debug("Solving part 1.")

	importantValves := findImportantValves(valves)
	valveState := uint64(0)
	highestPressure := findHighestPressure("AA", 0, 30, importantValves, valves, valveState)

	log.Debug("Part 1 solved.")
	return strconv.Itoa(highestPressure)
}

func solvePart2(valves map[string]*valve) string {
	log.Debug("Solving part 2.")

	// TODO

	log.Debug("Part 2 solved.")
	return "Not Implemented"
}

func findHighestPressure(fromLabel string, pressure int, timeLimit int, importantValveLabels []string, allValves map[string]*valve, valveState uint64) int {
	log.Debug("BEGIN visit()")
	log.Debugf("fromLabel = %s", fromLabel)
	log.Debugf("pressure = %d", pressure)
	log.Debugf("timeLimit = %d", timeLimit)

	if timeLimit <= 0 {
		return pressure
	}

	highestPressure := pressure
	for _, toLabel := range importantValveLabels {
		toValve := allValves[toLabel]
		valveMask := uint64(1) << toValve.index
		if valveState&valveMask == 0 {
			path := findPath(fromLabel, toLabel, allValves)
			newValveState := valveState | valveMask
			timeRemaining := timeLimit - len(path) - 1
			newPressure := findHighestPressure(toLabel, (timeRemaining*toValve.flowRate)+pressure, timeRemaining, importantValveLabels, allValves, newValveState)
			highestPressure = common.MaxInt(highestPressure, newPressure)
		}
	}

	return highestPressure
}

func findImportantValves(valves map[string]*valve) []string {
	importantValves := []string{}

	for label, valve := range valves {
		if valve.flowRate > 0 {
			importantValves = append(importantValves, label)
		}
	}

	log.Debugf("importantValves = %s", importantValves)

	return importantValves
}

func findPath(startValveLabel string, goalValveLabel string, valves map[string]*valve) []string {
	start := &valveState{
		label:  startValveLabel,
		valves: valves,
	}

	goal := &valveState{
		label:  goalValveLabel,
		valves: valves,
	}

	search := common.NewAStarSearch(start, goal, heuristicFunction, possibleNextStatesFunction, len(valves))
	states := search.Search()
	path := []string{}

	for _, state := range states[1:] {
		path = append(path, state.(*valveState).label)
	}

	return path
}

type valveState struct {
	label  string
	valves map[string]*valve
}

func (thisState *valveState) Cost() int {
	return 1
}

func (thisState *valveState) Key() string {
	return thisState.label
}

func heuristicFunction(current common.State, goal common.State) int {
	return 0
}

func possibleNextStatesFunction(current common.State) []common.State {
	currentState := current.(*valveState)
	currentValve := currentState.valves[currentState.label]
	nextStates := []common.State{}

	for _, adjacentLabel := range currentValve.tunnelsTo {
		nextState := &valveState{
			label:  adjacentLabel,
			valves: currentState.valves,
		}
		nextStates = append(nextStates, nextState)
	}

	return nextStates
}
