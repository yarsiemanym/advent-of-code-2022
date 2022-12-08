package day08

import (
	"strconv"

	"github.com/yarsiemanym/advent-of-code-2022/common"
)

type tree struct {
	Height int
}

func (thisTree *tree) IsShorterThan(otherTree *tree) bool {
	return thisTree.Height < otherTree.Height
}

func parseTree(char rune) *tree {
	height, err := strconv.Atoi(string(char))
	common.Check(err)

	return &tree{
		Height: height,
	}
}
