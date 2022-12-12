package day12

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type heightMap struct {
	plane *common.BoundedPlane
	start *common.Point
	end   *common.Point
}

func NewHeightMap(height int, width int) *heightMap {
	return &heightMap{
		plane: common.NewBoundedPlane(height, width),
	}
}

func (thisHeightMap *heightMap) ShortestPathFromStart() []*common.Point {
	return thisHeightMap.ShortestPathFrom(thisHeightMap.start)
}

func (thisHeightMap *heightMap) ShortestPathFrom(start *common.Point) []*common.Point {
	startingState := &heightMapState{
		point: start,
		plane: thisHeightMap.plane,
	}

	endState := &heightMapState{
		point: thisHeightMap.end,
		plane: thisHeightMap.plane,
	}

	ceiling := (thisHeightMap.plane.Span().End().Y() - thisHeightMap.plane.Span().Start().Y() + 1) * (thisHeightMap.plane.Span().End().X() - thisHeightMap.plane.Span().Start().X() + 1)
	aStar := common.NewAStarSearch(startingState, endState, heuristicFunction, possibleNextStatesFunction, ceiling)
	states := aStar.Search()

	if states == nil {
		return nil
	}

	points := make([]*common.Point, len(states))

	for index, state := range states {
		points[index] = state.(*heightMapState).Point()
	}

	return points
}

func parseHeightMap(text string) *heightMap {
	lines := common.Split(text, "\n")
	height := len(lines) - 1
	width := len(lines[0])
	plane := common.NewBoundedPlane(height, width)
	start := common.New2DPoint(0, 0)
	end := common.New2DPoint(0, 0)

	for y, line := range lines {
		for x, char := range line {
			point := common.New2DPoint(x, y)

			if char == 'S' {
				start = point
				plane.SetValueAt(point, 'a')
			} else if char == 'E' {
				end = point
				plane.SetValueAt(point, 'z')
			} else {
				plane.SetValueAt(point, char)
			}
		}
	}

	return &heightMap{
		plane: plane,
		start: start,
		end:   end,
	}
}

/********************
 * A* Search Stuff
 *******************/

type heightMapState struct {
	point *common.Point
	plane *common.BoundedPlane
}

func (state *heightMapState) Cost() int {
	return 1
}

func (state *heightMapState) Key() string {
	return state.point.String()
}

func (state *heightMapState) Point() *common.Point {
	return state.point
}

func heuristicFunction(current common.State, goal common.State) int {
	return current.(*heightMapState).point.ManhattanDistance(goal.(*heightMapState).point)
}

func possibleNextStatesFunction(current common.State) []common.State {
	state := current.(*heightMapState)
	here := state.point

	log.Debugf("Determine next steps from %s.", here)

	nextStates := []common.State{}
	neighbors := current.(*heightMapState).plane.GetVonNeumannNeighbors(here)

	for _, neighbor := range neighbors {
		currentHeight := state.plane.GetValueAt(here).(rune)
		neighborHeight := state.plane.GetValueAt(neighbor).(rune)

		if neighborHeight <= currentHeight+1 {
			log.Debugf("Neighbor %s is has potential.", neighbor)
			nextState := &heightMapState{
				point: neighbor,
				plane: state.plane,
			}

			nextStates = append(nextStates, nextState)
		}
	}

	return nextStates
}
