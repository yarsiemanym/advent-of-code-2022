package day08

import (
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type forest struct {
	Trees *common.BoundedPlane
}

func (thisForest *forest) GetVisibleTreePoints() []*common.Point {
	visibleTreePoints := []*common.Point{}

	for _, point := range thisForest.Trees.GetAllPoints() {

		if thisForest.IsTreeVisibleAt(point) {
			visibleTreePoints = append(visibleTreePoints, point)
		}
	}

	return visibleTreePoints
}

func (thisForest *forest) IsTreeVisibleAt(thisPoint *common.Point) bool {

	log.Debugf("Checking tree at %s for visibility.", thisPoint)

	thisTree := thisForest.Trees.GetValueAt(thisPoint).(*tree)
	log.Debugf("height = %d", thisTree.Height)

	isVisible := thisForest.Trees.IsEdge(thisPoint)

	if !isVisible {
		isVisible = true

		for y := thisPoint.Y() + 1; y <= thisForest.Trees.Span().End().Y(); y++ {
			otherPoint := common.New2DPoint(thisPoint.X(), y)
			otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

			if !otherTree.IsShorterThan(thisTree) {
				isVisible = false
				break
			}
		}
	}

	if !isVisible {
		isVisible = true

		for x := thisPoint.X() + 1; x <= thisForest.Trees.Span().End().X(); x++ {
			otherPoint := common.New2DPoint(x, thisPoint.Y())
			otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

			if !otherTree.IsShorterThan(thisTree) {
				isVisible = false
				break
			}
		}
	}

	if !isVisible {
		isVisible = true

		for y := thisPoint.Y() - 1; y >= thisForest.Trees.Span().Start().Y(); y-- {
			otherPoint := common.New2DPoint(thisPoint.X(), y)
			otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

			if !otherTree.IsShorterThan(thisTree) {
				isVisible = false
				break
			}
		}
	}

	if !isVisible {
		isVisible = true

		for x := thisPoint.X() - 1; x >= thisForest.Trees.Span().Start().X(); x-- {
			otherPoint := common.New2DPoint(x, thisPoint.Y())
			otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

			if !otherTree.IsShorterThan(thisTree) {
				isVisible = false
				break
			}
		}
	}

	log.Debugf("isVisible = %t", isVisible)

	return isVisible
}

func (thisForest *forest) GetHighestScenicScore() int {
	highestScenicScore := 0

	for _, point := range thisForest.Trees.GetAllPoints() {

		thisTreesScenicScore := thisForest.ScoreTreeAt(point)

		if thisTreesScenicScore > highestScenicScore {
			highestScenicScore = thisTreesScenicScore
		}
	}

	return highestScenicScore
}

func (thisForest *forest) ScoreTreeAt(thisPoint *common.Point) int {
	log.Debugf("Determining scenic score of tree at %s.", thisPoint)

	thisTree := thisForest.Trees.GetValueAt(thisPoint).(*tree)
	log.Debugf("height = %d", thisTree.Height)

	upScore, rightScore, downScore, leftScore := 0, 0, 0, 0

	for y := thisPoint.Y() + 1; y <= thisForest.Trees.Span().End().Y(); y++ {
		otherPoint := common.New2DPoint(thisPoint.X(), y)
		otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

		upScore++

		if !otherTree.IsShorterThan(thisTree) {
			break
		}
	}

	for x := thisPoint.X() + 1; x <= thisForest.Trees.Span().End().X(); x++ {
		otherPoint := common.New2DPoint(x, thisPoint.Y())
		otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

		rightScore++

		if !otherTree.IsShorterThan(thisTree) {
			break
		}
	}

	for y := thisPoint.Y() - 1; y >= thisForest.Trees.Span().Start().Y(); y-- {
		otherPoint := common.New2DPoint(thisPoint.X(), y)
		otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

		downScore++

		if !otherTree.IsShorterThan(thisTree) {
			break
		}
	}

	for x := thisPoint.X() - 1; x >= thisForest.Trees.Span().Start().X(); x-- {
		otherPoint := common.New2DPoint(x, thisPoint.Y())
		otherTree := thisForest.Trees.GetValueAt(otherPoint).(*tree)

		leftScore++

		if !otherTree.IsShorterThan(thisTree) {
			break
		}
	}

	score := upScore * rightScore * downScore * leftScore

	log.Debugf("score = %d", score)

	return score
}

func parseForest(text string) *forest {
	if text == "" {
		return nil
	}

	lines := common.Split(text, "\n")

	forest := &forest{
		Trees: common.NewBoundedPlane(len(lines)-1, len(lines[0])),
	}

	for rowIndex, line := range lines {
		if line == "" {
			continue
		}

		for colIndex, char := range line {
			forest.Trees.SetValueAt(common.New2DPoint(colIndex, rowIndex), parseTree(char))
		}
	}

	return forest
}
