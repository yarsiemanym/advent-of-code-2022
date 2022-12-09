package common

import (
	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)

type State interface {
	Cost() int
}

type HeuristicFunction func(current State, goal State) int
type PossibleNextStatesFunction func(current State) []State

type AStarSearch struct {
	start     State
	goal      State
	heuristic HeuristicFunction
	next      PossibleNextStatesFunction
}

func NewAStarSearch(start State, goal State, heuristic HeuristicFunction, next PossibleNextStatesFunction) *AStarSearch {
	return &AStarSearch{
		start:     start,
		goal:      goal,
		heuristic: heuristic,
		next:      next,
	}
}

func (search *AStarSearch) Search() []State {
	cameFrom := map[State]State{}
	costSoFar := map[State]int{}

	frontier := priorityqueue.New().WithComparator(&comparator{})
	frontier.Add(priorityQueueItem{
		priority: 0,
		value:    search.start,
	})

	cameFrom[search.start] = search.start
	costSoFar[search.start] = 0

	for frontier.Size() > 0 {
		current := frontier.Poll().(priorityQueueItem).value.(State)

		if current == search.goal {
			break
		}

		for _, next := range search.next(current) {
			currentCost, exists := costSoFar[current]
			newCost := costSoFar[current] + next.Cost()

			if !exists || newCost < currentCost {
				costSoFar[next] = newCost
				cameFrom[next] = current
				frontier.Add(priorityQueueItem{
					priority: newCost + search.heuristic(next, search.goal),
					value:    next,
				})
			}

		}
	}

	path := []State{}
	here := search.goal

	for here != nil {
		path = append([]State{here}, path...)
		here = cameFrom[here]
	}

	return path
}

type priorityQueueItem struct {
	priority int
	value    interface{}
}

type comparator struct{}

func (comparator *comparator) Compare(v1 interface{}, v2 interface{}) (int, error) {
	item1 := v1.(priorityQueueItem)
	item2 := v2.(priorityQueueItem)

	if item1.priority < item2.priority {
		return -1, nil
	} else if item1.priority > item2.priority {
		return 1, nil
	} else {
		return 0, nil
	}
}
