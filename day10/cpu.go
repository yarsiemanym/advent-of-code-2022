package day10

import (
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type cpu struct {
	cycle          int
	registerX      int
	signalStrength int
	callback       cpuCycleCallback
}

type cpuCycleCallback func(cycle int, registerX int)

func (thisCPU *cpu) Cycle() int {
	return thisCPU.cycle
}

func (thisCPU *cpu) RegisterX() int {
	return thisCPU.registerX
}

func (thisCPU *cpu) SignalStrength() int {
	return thisCPU.signalStrength
}

func (thisCPU *cpu) ExecuteInstructions(instructions []string) {

	log.Debug("Begin Cycle 0")
	thisCPU.Callback()
	log.Debug("End Cycle 0")

	for _, instruction := range instructions {
		thisCPU.ExecuteInstruction(instruction)
	}
}

func (thisCPU *cpu) ExecuteInstruction(instruction string) {
	if instruction == "" {
		return
	}

	tokens := common.Split(instruction, " ")
	command := tokens[0]

	switch command {
	case "addx":
		thisCPU.cycle++
		log.Debugf("Begin Cycle %d", thisCPU.cycle)
		thisCPU.AccumulateSignalStrength()
		thisCPU.Callback()
		log.Debugf("End Cycle %d", thisCPU.cycle)

		thisCPU.cycle++
		log.Debugf("Begin Cycle %d", thisCPU.cycle)
		thisCPU.AccumulateSignalStrength()
		parameter, err := strconv.Atoi(tokens[1])
		common.Check(err)
		thisCPU.registerX += parameter
		thisCPU.Callback()
		log.Debugf("End Cycle %d", thisCPU.cycle)

	case "noop":
		thisCPU.cycle++
		log.Debugf("Begin Cycle %d", thisCPU.cycle)
		thisCPU.AccumulateSignalStrength()
		thisCPU.Callback()
		log.Debugf("End Cycle %d", thisCPU.cycle)

	default:
		log.Panicf("Unsupported command \"%s\".", command)

	}
}

func (thisCPU *cpu) AccumulateSignalStrength() {
	if (thisCPU.cycle-20)%40 == 0 {
		thisCPU.signalStrength += thisCPU.cycle * thisCPU.registerX
	}
}

func (thisCPU *cpu) Callback() {
	if thisCPU.callback != nil {
		thisCPU.callback(thisCPU.cycle, thisCPU.registerX)
	}
}

func NewCPU(callback cpuCycleCallback) *cpu {
	return &cpu{
		cycle:     0,
		registerX: 1,
		callback:  callback,
	}
}
