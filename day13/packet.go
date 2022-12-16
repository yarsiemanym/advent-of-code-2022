package day13

import (
	"log"
	"reflect"
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

var intType = reflect.TypeOf(0)
var sliceType = reflect.TypeOf([]any{})

type packetPair struct {
	packet1 *packet
	packet2 *packet
}

func parsePacketPair(text string) any {
	tokens := common.Split(text, "\n")

	return &packetPair{
		packet1: parsePacket(tokens[0]).(*packet),
		packet2: parsePacket(tokens[1]).(*packet),
	}
}

func (thisPacketPair *packetPair) IsOrderedCorrectly() bool {
	return thisPacketPair.packet1.IsLessThan(thisPacketPair.packet2)
}

type packet struct {
	str  string
	data []any
}

func (thisPacket *packet) IsLessThan(otherPacket *packet) bool {
	return comparePacketData(thisPacket.data, otherPacket.data) == -1
}

func (thisPacket *packet) IsEqualTo(otherPacket *packet) bool {
	return comparePacketData(thisPacket.data, otherPacket.data) == 0
}

func comparePacketData(data1 []any, data2 []any) int {

	end := common.MaxInt(len(data1), len(data2))

	for index := 0; index < end; index++ {

		if index >= len(data1) && index < len(data2) {
			return -1
		}

		if index >= len(data2) && index < len(data1) {
			return 1
		}

		value1, value2 := data1[index], data2[index]

		if reflect.TypeOf(value1) == intType && reflect.TypeOf(value2) == intType {

			if value1.(int) == value2.(int) {
				continue
			} else if value1.(int) < value2.(int) {
				return -1
			} else if value1.(int) > value2.(int) {
				return 1
			}

		} else {

			var slice1 []any
			var slice2 []any

			if reflect.TypeOf(value1) == intType {
				slice1 = []any{value1}
			} else {
				slice1 = value1.([]any)
			}

			if reflect.TypeOf(value2) == intType {
				slice2 = []any{value2}
			} else {
				slice2 = value2.([]any)
			}

			comparison := comparePacketData(slice1, slice2)

			if comparison == 0 {
				continue
			} else {
				return comparison
			}
		}
	}

	return 0
}

func findPackets(packets []*packet, dividerPackets []*packet) []int {
	indices := []int{}

	for index, packet := range packets {
		for _, dividerPacket := range dividerPackets {
			if dividerPacket.IsEqualTo(packet) {
				indices = append(indices, index)
			}
		}
	}

	return indices
}

func parsePacket(text string) any {
	if text == "" {
		return nil
	}

	data, _ := parsePacketData(text[1:])
	return &packet{
		str:  text,
		data: data,
	}
}

func parsePacketData(text string) ([]any, int) {
	data := []any{}
	consumed := 0
	buffer := ""

	for index := 0; index < len(text); index++ {
		character := text[index]

		switch character {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			buffer += string(character)

		case '[':
			value, consumed := parsePacketData(text[index+1:])
			data = append(data, value)
			index += consumed + 1

		case ',':
			if buffer != "" {
				value, err := strconv.Atoi(buffer)
				common.Check(err)
				buffer = ""
				data = append(data, value)
			}

		case ']':
			if buffer != "" {
				value, err := strconv.Atoi(buffer)
				common.Check(err)
				buffer = ""
				data = append(data, value)
			}

			return data, consumed

		default:
			log.Panicf("Unexpected character '%c'.", character)
		}

		consumed++
	}

	return data, consumed
}
