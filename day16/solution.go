package day16

import (
	"sort"
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

	pressureCache := map[pressureCacheKey]int{}
	pathCache := map[pathCacheKey][]string{}
	importantValves := findImportantValves(valves)
	valveState := 0
	populatePressureCache("AA", 0, 30, importantValves, valves, valveState, pressureCache, pathCache)

	highestPressure := 0

	for _, pressure := range pressureCache {
		if pressure > highestPressure {
			highestPressure = pressure
		}
	}

	log.Debug("Part 1 solved.")
	return strconv.Itoa(highestPressure)
}

func solvePart2(valves map[string]*valve) string {
	log.Debug("Solving part 2.")

	pressureCache := map[pressureCacheKey]int{}
	pathCache := map[pathCacheKey][]string{}
	importantValves := findImportantValves(valves)
	valveState := 0
	populatePressureCache("AA", 0, 26, importantValves, valves, valveState, pressureCache, pathCache)

	pressures := map[int]int{}

	for key, pressure := range pressureCache {
		valveState := key.startingValveState
		highestPressure, exists := pressures[valveState]

		if !exists {
			highestPressure = 0
		}

		if pressure > highestPressure {
			pressures[valveState] = pressure
		}
	}

	highestPressure := 0

	for thisValveState, thisPressure := range pressures {
		for otherValveState, otherPressure := range pressures {
			if thisValveState&otherValveState == 0 && thisPressure+otherPressure > highestPressure {
				highestPressure = thisPressure + otherPressure
			}
		}
	}

	log.Debug("Part 2 solved.")
	return strconv.Itoa(highestPressure)
}

type pressureCacheKey struct {
	startLabel         string
	startingPressure   int
	startingValveState int
	timeRemaining      int
}

func populatePressureCache(startLabel string, startingPressure int, timeRemaining int, importantValveLabels []string,
	allValves map[string]*valve, valveState int, pressureCache map[pressureCacheKey]int, pathCache map[pathCacheKey][]string) {

	cacheKey := pressureCacheKey{
		startLabel:         startLabel,
		startingPressure:   startingPressure,
		timeRemaining:      timeRemaining,
		startingValveState: valveState,
	}

	_, exists := pressureCache[cacheKey]

	if exists {
		return
	}

	if timeRemaining <= 0 {
		pressureCache[cacheKey] = startingPressure
		return
	}

	highestPressure := startingPressure
	for index, toLabel := range importantValveLabels {
		valveMask := 1 << index
		isOpen := valveState&valveMask != 0
		if !isOpen {
			path := findPath(startLabel, toLabel, allValves, pathCache)
			newValveState := valveState | valveMask
			newTimeRemaining := timeRemaining - len(path) - 1
			toValve := allValves[toLabel]
			populatePressureCache(toLabel, (newTimeRemaining*toValve.flowRate)+startingPressure, newTimeRemaining,
				importantValveLabels, allValves, newValveState, pressureCache, pathCache)
		}
	}

	pressureCache[cacheKey] = highestPressure
}

func findImportantValves(valves map[string]*valve) []string {
	importantValves := []string{}

	for label, valve := range valves {
		if valve.flowRate > 0 {
			importantValves = append(importantValves, label)
		}
	}

	sort.StringSlice.Sort(importantValves)
	return importantValves
}

type pathCacheKey struct {
	start string
	goal  string
}

func findPath(startValveLabel string, goalValveLabel string, valves map[string]*valve,
	pathCache map[pathCacheKey][]string) []string {

	cacheKey := pathCacheKey{
		start: startValveLabel,
		goal:  goalValveLabel,
	}
	cachedValue, exists := pathCache[cacheKey]

	if exists {
		return cachedValue
	}

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

	pathCache[cacheKey] = path

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
