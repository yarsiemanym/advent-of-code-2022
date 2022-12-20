package day17

import (
	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

var leftVector = common.New2DPoint(-1, 0)
var rightVector = common.New2DPoint(1, 0)
var downVector = common.New2DPoint(0, -1)

type chamber struct {
	spaces              [][]rune
	rockGenerator       *rockGenerator
	jets                *jets
	fallingRock         *rock
	fallingRockPosition *common.Point
	settledRockCounter  int
}

func NewChamber(jetsPattern string) *chamber {
	return &chamber{
		spaces: [][]rune{
			{emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace},
			{emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace},
			{emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace},
			{emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace},
		},
		rockGenerator:       NewRockGenerator(),
		jets:                parseJets(jetsPattern),
		fallingRock:         nil,
		fallingRockPosition: nil,
		settledRockCounter:  0,
	}
}

func (chamber *chamber) Tick() {
	newRock := chamber.EnsureFallingRockExists()

	if newRock {
		log.Debug("A new rock appears.")
		chamber.DrawFallingRock()
	} else {
		chamber.ActivateJets()
		chamber.ActivateGravity()
	}
}

func (chamber *chamber) ActivateJets() {
	direction := chamber.jets.Blast()
	log.Debugf("Jets pushing %c.", direction)
	var vector *common.Point = nil

	switch direction {
	case left:
		vector = leftVector
	case right:
		vector = rightVector
	default:
		log.Panicf("Unsupported direction '%c'.", direction)
	}

	chamber.MoveRock(vector)
}

func (chamber *chamber) ActivateGravity() {
	vector := downVector
	log.Debug("Gravity pulling v.")
	isSettled := !chamber.MoveRock(vector)

	if isSettled {
		log.Debug("Rock has settled.")
		chamber.EraseFallingRock()
		chamber.DrawSettledRock()
		chamber.settledRockCounter++
		chamber.fallingRock = nil
		chamber.fallingRockPosition = nil
	}
}

func (chamber *chamber) MoveRock(vector *common.Point) bool {
	rock := chamber.fallingRock
	collides := false

	for y := rock.Height() - 1; !collides && y >= 0; y-- {
		for x := 0; !collides && x < rock.Width(); x++ {
			relativePoint := common.New2DPoint(x, y)
			absolutePoint := chamber.fallingRockPosition.Add(common.New2DPoint(x, -y))
			nextAbsolutePoint := absolutePoint.Add(vector)

			if nextAbsolutePoint.Y() < 0 {
				log.Debugf("Collision detected with the bottom of the chamber.")
				collides = true
			} else if nextAbsolutePoint.X() < 0 || nextAbsolutePoint.X() >= chamber.Width() {
				log.Debugf("Collision detected with the side of the chamber.")
				collides = true
			} else if rock.fallingBits[relativePoint.Y()][relativePoint.X()] == fallingRock &&
				chamber.spaces[nextAbsolutePoint.Y()][nextAbsolutePoint.X()] == settledRock {
				log.Debugf("Collision detected with settled rock.")
				collides = true
			}
		}
	}

	if !collides {
		chamber.EraseFallingRock()
		chamber.fallingRockPosition = chamber.fallingRockPosition.Add(vector)
		chamber.DrawFallingRock()
	}

	return !collides
}

func (chamber *chamber) EraseFallingRock() {
	for y := chamber.fallingRock.Height() - 1; y >= 0; y-- {
		for x := 0; x < chamber.fallingRock.Width(); x++ {
			if chamber.fallingRock.fallingBits[y][x] != emptySpace {
				chamber.Draw(chamber.fallingRockPosition.X()+x, chamber.fallingRockPosition.Y()-y, emptySpace)
			}
		}
	}
}

func (chamber *chamber) DrawFallingRock() {
	for y := chamber.fallingRock.Height() - 1; y >= 0; y-- {
		for x := 0; x < chamber.fallingRock.Width(); x++ {
			if chamber.fallingRock.fallingBits[y][x] != emptySpace {
				chamber.Draw(chamber.fallingRockPosition.X()+x, chamber.fallingRockPosition.Y()-y, chamber.fallingRock.fallingBits[y][x])
			}
		}
	}
}

func (chamber *chamber) DrawSettledRock() {
	for y := chamber.fallingRock.Height() - 1; y >= 0; y-- {
		for x := 0; x < chamber.fallingRock.Width(); x++ {
			if chamber.fallingRock.settledBits[y][x] != emptySpace {
				chamber.Draw(chamber.fallingRockPosition.X()+x, chamber.fallingRockPosition.Y()-y, chamber.fallingRock.settledBits[y][x])
			}
		}
	}
}

func (chamber *chamber) Draw(x int, y int, symbol rune) {
	chamber.spaces[y][x] = symbol
}

func (chamber *chamber) Height() int {
	return len(chamber.spaces)
}

func (chamber *chamber) Width() int {
	return len(chamber.spaces[0])
}

func (chamber *chamber) EnsureFallingRockExists() bool {
	rockCreated := false

	if chamber.fallingRock == nil {
		chamber.fallingRock = chamber.rockGenerator.NextRock()
		rockCreated = true
	}

	if chamber.fallingRockPosition == nil {
		desiredHeight := chamber.HeightOfSettledRocks() + 3 + chamber.fallingRock.Height()
		chamber.GrowTo(desiredHeight)
		chamber.fallingRockPosition = common.New2DPoint(2, desiredHeight-1)
		rockCreated = true
	}

	return rockCreated
}

func (chamber *chamber) HeightOfSettledRocks() int {
	for y := len(chamber.spaces) - 1; y >= 0; y-- {
		row := chamber.spaces[y]
		for x := 0; x < len(row); x++ {
			space := row[x]
			if space == settledRock {
				return y + 1
			}
		}
	}

	return 0
}

func (chamber *chamber) NumberOfSettledRocks() int {
	return chamber.settledRockCounter
}

func (chamber *chamber) GrowTo(desiredHeight int) {
	for y := chamber.Height(); y < desiredHeight; y++ {
		chamber.spaces = append(chamber.spaces, []rune{emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace, emptySpace})
	}
}

func (chamber *chamber) String() string {
	output := ""

	for rowIndex := chamber.Height() - 1; rowIndex >= common.MaxInt(0, chamber.Height()-50); rowIndex-- {
		output += "|"
		row := chamber.spaces[rowIndex]
		for _, space := range row {
			output += string(space)
		}
		output += "|\n"
	}

	output += "+-------+\n"
	return output
}
