package day11

import (
	"regexp"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type monkey struct {
	id              int
	items           *common.Queue
	operation       *operation
	divisibleBy     int
	trueCaseTarget  int
	falseCaseTarget int
	activity        int
}

func (thisMonkey *monkey) TakeTurn(allMonkeys []*monkey, reliefFactor int, leastCommonMultiple uint64) {
	log.Debugf("Monkey %d is taking its turn.", thisMonkey.id)

	for !thisMonkey.items.IsEmpty() {
		item := thisMonkey.items.Pop().(*item)

		thisMonkey.Inspect(item)
		thisMonkey.GetBoredWith(item, reliefFactor, leastCommonMultiple)
		targetMonkey := thisMonkey.ChooseTarget(item, allMonkeys)
		thisMonkey.ThrowItem(item, targetMonkey)
	}
}

func (thisMonkey *monkey) Inspect(item *item) {
	log.Debugf("Monkey %d is inspecting item %d", thisMonkey.id, item.worryLevel)
	item.worryLevel = thisMonkey.operation.Execute(item.worryLevel)
	log.Debugf("worryLevel = %d", item.worryLevel)
	thisMonkey.activity++
}

func (thisMonkey *monkey) GetBoredWith(item *item, reliefFactor int, leastCommonMultiple uint64) {
	log.Debugf("Monkey %d is getting bored with item %d", thisMonkey.id, item.worryLevel)
	item.worryLevel /= uint64(reliefFactor)
	item.worryLevel %= leastCommonMultiple
	log.Debugf("worryLevel = %d", item.worryLevel)
}

func (thisMonkey *monkey) ChooseTarget(item *item, allMonkeys []*monkey) *monkey {
	var targetMonkey *monkey
	log.Debugf("Testing if item %d is divisible by %d.", item.worryLevel, thisMonkey.divisibleBy)
	if item.worryLevel%uint64(thisMonkey.divisibleBy) == 0 {
		log.Debugf("It is. Choosing monkey %d as target.", thisMonkey.trueCaseTarget)
		targetMonkey = allMonkeys[thisMonkey.trueCaseTarget]
	} else {
		log.Debugf("It is not. Choosing monkey %d as target.", thisMonkey.falseCaseTarget)
		targetMonkey = allMonkeys[thisMonkey.falseCaseTarget]
	}

	return targetMonkey
}

func (thisMonkey *monkey) ThrowItem(item *item, targetMonkey *monkey) {
	log.Debugf("Throwing item %d to monkey %d.", item.worryLevel, targetMonkey.id)
	targetMonkey.items.Push(item)
}

func parseMonkey(text string) interface{} {
	lines := common.Split(text, "\n")

	idText := lines[0]
	idRegexp := regexp.MustCompile(`Monkey (\d+):`)
	idMatches := idRegexp.FindStringSubmatch(idText)
	id, err := strconv.Atoi(idMatches[1])
	common.Check(err)

	startingItemsText := lines[1]
	startingItemsRegexp := regexp.MustCompile(`Starting items: ((\d+(, )?)+)`)
	startingItemsMatches := startingItemsRegexp.FindStringSubmatch(startingItemsText)
	tokens := common.Split(startingItemsMatches[1], ", ")
	startingItems := common.NewQueue()

	for _, token := range tokens {
		worryLevel, err := strconv.ParseUint(token, 10, 64)
		common.Check(err)
		item := &item{
			worryLevel: worryLevel,
		}
		startingItems.Push(item)
	}

	operationText := lines[2]
	operation := parseOperation(operationText)

	testText := lines[3]
	testRegexp := regexp.MustCompile(`Test: divisible by (\d+)`)
	testMatches := testRegexp.FindStringSubmatch(testText)
	divisibleBy, err := strconv.Atoi(testMatches[1])
	common.Check(err)

	trueCaseText := lines[4]
	trueCaseRegexp := regexp.MustCompile(`If true: throw to monkey (\d+)`)
	trueCaseMatches := trueCaseRegexp.FindStringSubmatch(trueCaseText)
	trueCaseTarget, err := strconv.Atoi(trueCaseMatches[1])
	common.Check(err)

	falseCaseText := lines[5]
	falseCaseRegexp := regexp.MustCompile(`If false: throw to monkey (\d+)`)
	falseCaseMatches := falseCaseRegexp.FindStringSubmatch(falseCaseText)
	falseCaseTarget, err := strconv.Atoi(falseCaseMatches[1])
	common.Check(err)

	return &monkey{
		id:              id,
		items:           startingItems,
		operation:       operation,
		divisibleBy:     divisibleBy,
		trueCaseTarget:  trueCaseTarget,
		falseCaseTarget: falseCaseTarget,
		activity:        0,
	}
}
