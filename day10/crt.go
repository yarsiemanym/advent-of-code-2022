package day10

import log "github.com/sirupsen/logrus"

const screenWidth int = 40

type crt struct {
	pixels [][]bool
}

func NewCRT() *crt {
	return &crt{
		pixels: [][]bool{
			make([]bool, screenWidth),
			make([]bool, screenWidth),
			make([]bool, screenWidth),
			make([]bool, screenWidth),
			make([]bool, screenWidth),
			make([]bool, screenWidth),
		},
	}
}

func (thisCRT *crt) MaybeLightPixel(cycle int, registerX int) {
	log.Debugf("Checking to see if pixel %d should be lit.", cycle)
	log.Debugf("Sprite location is %d.", registerX)

	if cycle%screenWidth >= registerX-1 && cycle%screenWidth <= registerX+1 {
		thisCRT.LightPixel(cycle)
	} else {
		log.Debugf("Pixel %d remains dark.", cycle)
	}
}

func (thisCRT *crt) LightPixel(index int) {
	log.Debugf("Lighting pixel %d.", index)
	thisCRT.pixels[index/screenWidth][index%screenWidth] = true
}

func (thisCRT *crt) Render() string {
	output := ""

	for _, line := range thisCRT.pixels {
		for _, pixel := range line {
			if pixel {
				output += "#"
			} else {
				output += "."
			}
		}

		output += "\n"
	}

	return output
}
