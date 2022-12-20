package vt100

import "fmt"

func ClearScreen() {
	fmt.Print("\x1b[2J")
}
