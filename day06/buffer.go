package day06

type buffer struct {
	DataStream string
}

func (thisBuffer *buffer) StartOfPacket() int {
	return thisBuffer.DetectUniqueSequence(4)
}

func (thisBuffer *buffer) StartOfMessage() int {
	return thisBuffer.DetectUniqueSequence(14)
}

func (thisBuffer *buffer) DetectUniqueSequence(length int) int {
	for index := length; index < len(thisBuffer.DataStream); index++ {
		potentialMarker := thisBuffer.DataStream[index-length : index]

		charMap := map[rune]int{}

		for _, char := range potentialMarker {
			charMap[char] = charMap[char] + 1
		}

		isUniqueSequence := true

		for _, count := range charMap {
			if count > 1 {
				isUniqueSequence = false
				break
			}
		}

		if isUniqueSequence {
			return index
		}
	}

	return -1
}
