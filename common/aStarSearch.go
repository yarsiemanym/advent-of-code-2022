package common

import (
	"math"

	log "github.com/sirupsen/logrus"

	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)

type State interface {
	Key() string
	Cost() int
}

type HeuristicFunction func(current State, goal State) int
type PossibleNextStatesFunction func(current State) []State

type AStarSearch struct {
	start     State
	goal      State
	heuristic HeuristicFunction
	next      PossibleNextStatesFunction
	ceiling   int
}

func NewAStarSearch(start State, goal State, heuristic HeuristicFunction, next PossibleNextStatesFunction, ceiling int) *AStarSearch {
	return &AStarSearch{
		start:     start,
		goal:      goal,
		heuristic: heuristic,
		next:      next,
		ceiling:   ceiling,
	}
}

func (search *AStarSearch) Search() []State {
	log.Trace("Begin A* search.")

	allStates := map[string]State{}
	cameFrom := map[string]string{}
	costSoFar := map[string]int{}
	frontier := priorityqueue.New().WithComparator(&comparator{})
	frontier.Add(priorityQueueItem{
		priority: 0,
		value:    search.start,
	})

	allStates[search.start.Key()] = nil
	cameFrom[search.start.Key()] = ""
	costSoFar[search.start.Key()] = 0

	for frontier.Size() > 0 {
		current := frontier.Poll().(priorityQueueItem).value.(State)
		log.Tracef("Current state is %s.", current.Key())

		for _, next := range search.next(current) {
			allStates[next.Key()] = next
			currentCost, exists := costSoFar[next.Key()]

			log.Tracef("Checking cost of moving to state %s.", next.Key())

			if !exists {
				log.Trace("No cost known yet.")
				currentCost = math.MaxInt
			}

			newCost := costSoFar[current.Key()] + next.Cost()
			log.Tracef("newCost = %d", newCost)

			if newCost > search.ceiling {
				frontier.Remove(priorityQueueItem{
					priority: costSoFar[next.Key()],
					value:    next,
				})
			} else if newCost < currentCost {
				log.Tracef("New shortest path to %s found.", next.Key())
				costSoFar[next.Key()] = newCost
				cameFrom[next.Key()] = current.Key()
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
		here = allStates[cameFrom[here.Key()]]
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
