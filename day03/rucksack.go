package day03

import (
	"unicode"

	log "github.com/sirupsen/logrus"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type rucksack struct {
	Items []rune
}

func (thisRuksack *rucksack) Compartment1() []rune {
	length := len(thisRuksack.Items)
	compartment1 := thisRuksack.Items[0 : length/2]
	return compartment1
}

func (thisRuksack *rucksack) Compartment2() []rune {
	length := len(thisRuksack.Items)
	compartment2 := thisRuksack.Items[length/2 : length]
	return compartment2
}

func (thisRuksack *rucksack) FindPackingErrors() []rune {
	packingErrors := map[rune]int{}

	log.Debugf("Inpecting rucksack \"%s\".", string(thisRuksack.Items))

	for _, item1 := range thisRuksack.Compartment1() {
		for _, item2 := range thisRuksack.Compartment2() {
			if item1 == item2 {
				log.Debugf("Error detected: '%c'.", item1)
				packingErrors[item1] += packingErrors[item1] + 1
				break
			}
		}
	}

	itemTypes := []rune{}

	for itemType := range packingErrors {
		itemTypes = append(itemTypes, itemType)
	}

	log.Debugf("itemTypes = \"%s\"", string(itemTypes))

	return itemTypes
}

func determinePriority(itemTypes []rune) int {

	log.Debugf("Determining priority of item types \"%s\".", string(itemTypes))

	priority := 0

	for _, itemType := range itemTypes {
		if unicode.IsUpper(itemType) {
			priority += int(itemType) - int('A') + 27
		} else {
			priority += int(itemType) - int('a') + 1
		}
	}

	log.Debugf("priority = %d", priority)

	return priority
}

func determineBadge(rucksacks []*rucksack) rune {

	intersection := rucksacks[0].Items

	for index := 1; index < len(rucksacks); index++ {
		intersection = common.Intersection(intersection, rucksacks[index].Items)
	}

	return intersection[0]
}

func parseRucksack(text string) interface{} {
	if text == "" {
		return nil
	}

	return &rucksack{
		Items: []rune(text),
	}
}
