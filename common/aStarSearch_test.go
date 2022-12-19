package common

import (
	"testing"
)

func Test_AStarSearch_Search_Adjacent(t *testing.T) {
	setupStates()
	start := allStates["AA"]
	end := allStates["DD"]
	search := NewAStarSearch(start, end, heuristicFunction, possibleNextStatesFunction, 10)
	path := search.Search()

	if len(path) != 2 {
		t.Errorf("Expected 2 but got %d.", len(path))
	} else if path[0].Key() != "AA" {
		t.Errorf("Expected AA but got \"%s\".", path[0].Key())
	} else if path[1].Key() != "DD" {
		t.Errorf("Expected DD but got \"%s\".", path[1].Key())
	}
}

func Test_AStarSearch_Search_NotAdjacent(t *testing.T) {
	setupStates()
	start := allStates["AA"]
	end := allStates["HH"]
	search := NewAStarSearch(start, end, heuristicFunction, possibleNextStatesFunction, 10)
	path := search.Search()

	if len(path) != 6 {
		t.Errorf("Expected 6 but got %d.", len(path))
	} else if path[0].Key() != "AA" {
		t.Errorf("Expected AA but got \"%s\".", path[0].Key())
	} else if path[1].Key() != "DD" {
		t.Errorf("Expected DD but got \"%s\".", path[1].Key())
	} else if path[2].Key() != "EE" {
		t.Errorf("Expected EE but got \"%s\".", path[2].Key())
	} else if path[3].Key() != "FF" {
		t.Errorf("Expected FF but got \"%s\".", path[3].Key())
	} else if path[4].Key() != "GG" {
		t.Errorf("Expected GG but got \"%s\".", path[4].Key())
	} else if path[5].Key() != "HH" {
		t.Errorf("Expected HH but got \"%s\".", path[5].Key())
	}
}

func setupStates() {
	allStates["AA"] = &testState{
		key:           "AA",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"DD", "II", "BB"},
	}

	allStates["BB"] = &testState{
		key:           "BB",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"CC", "AA"},
	}

	allStates["CC"] = &testState{
		key:           "CC",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"DD", "BB"},
	}

	allStates["DD"] = &testState{
		key:           "DD",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"CC", "AA", "EE"},
	}

	allStates["EE"] = &testState{
		key:           "EE",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"FF", "DD"},
	}

	allStates["FF"] = &testState{
		key:           "FF",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"EE", "GG"},
	}

	allStates["GG"] = &testState{
		key:           "GG",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"FF", "HH"},
	}

	allStates["HH"] = &testState{
		key:           "HH",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"GG"},
	}

	allStates["II"] = &testState{
		key:           "II",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"AA", "JJ"},
	}

	allStates["JJ"] = &testState{
		key:           "JJ",
		cost:          1,
		heuristic:     0,
		nextStateKeys: []string{"II"},
	}
}

var allStates = map[string]*testState{}

type testState struct {
	key           string
	cost          int
	heuristic     int
	nextStateKeys []string
}

func (state *testState) Cost() int {
	return state.cost
}

func (state *testState) Key() string {
	return state.key
}

func heuristicFunction(current State, goal State) int {
	return goal.(*testState).heuristic
}

func possibleNextStatesFunction(current State) []State {
	nextStates := []State{}

	for _, key := range current.(*testState).nextStateKeys {
		nextStates = append(nextStates, allStates[key])
	}

	return nextStates
}
