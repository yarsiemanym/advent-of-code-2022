package day17

import "log"

const (
	left  = '<'
	right = '>'
)

type jets struct {
	counter int
	pattern []rune
}

func (jets *jets) Blast() rune {
	direction := jets.pattern[jets.counter%len(jets.pattern)]
	jets.counter++
	return direction
}

func (jets *jets) BlastCount() int {
	return jets.counter
}

func (jets *jets) NextBlastIndex() int {
	return jets.counter % len(jets.pattern)
}

func parseJets(text string) *jets {
	pattern := []rune{}

	for _, char := range text {

		switch char {
		case left, right:
			pattern = append(pattern, char)
		default:
			log.Panicf("Invalid direction '%c'.", char)
		}
	}

	return &jets{
		counter: 0,
		pattern: pattern,
	}
}
