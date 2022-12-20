package day17

const (
	emptySpace  = '.'
	fallingRock = '@'
	settledRock = '#'
)

type rock struct {
	fallingBits [][]rune
	settledBits [][]rune
}

func (rock *rock) Height() int {
	return len(rock.fallingBits)
}

func (rock *rock) Width() int {
	return len(rock.fallingBits[0])
}

var rockTypes = []*rock{
	{
		fallingBits: [][]rune{
			{fallingRock, fallingRock, fallingRock, fallingRock},
		},
		settledBits: [][]rune{
			{settledRock, settledRock, settledRock, settledRock},
		},
	},
	{
		fallingBits: [][]rune{
			{emptySpace, fallingRock, emptySpace},
			{fallingRock, fallingRock, fallingRock},
			{emptySpace, fallingRock, emptySpace},
		},
		settledBits: [][]rune{
			{emptySpace, settledRock, emptySpace},
			{settledRock, settledRock, settledRock},
			{emptySpace, settledRock, emptySpace},
		},
	},
	{
		fallingBits: [][]rune{
			{emptySpace, emptySpace, fallingRock},
			{emptySpace, emptySpace, fallingRock},
			{fallingRock, fallingRock, fallingRock},
		},
		settledBits: [][]rune{
			{emptySpace, emptySpace, settledRock},
			{emptySpace, emptySpace, settledRock},
			{settledRock, settledRock, settledRock},
		},
	},
	{
		fallingBits: [][]rune{
			{fallingRock},
			{fallingRock},
			{fallingRock},
			{fallingRock},
		},
		settledBits: [][]rune{
			{settledRock},
			{settledRock},
			{settledRock},
			{settledRock},
		},
	},
	{
		fallingBits: [][]rune{
			{fallingRock, fallingRock},
			{fallingRock, fallingRock},
		},
		settledBits: [][]rune{
			{settledRock, settledRock},
			{settledRock, settledRock},
		},
	},
}

type rockGenerator struct {
	counter int
}

func NewRockGenerator() *rockGenerator {
	return &rockGenerator{
		counter: 0,
	}
}

func (generator *rockGenerator) RockCount() int {
	return generator.counter
}

func (generator *rockGenerator) NextRockIndex() int {
	return generator.counter % len(rockTypes)
}

func (generator *rockGenerator) NextRock() *rock {
	rock := rockTypes[generator.NextRockIndex()]
	generator.counter++
	return rock
}
