package vt100

import "fmt"

func MoveCursorToHome() {
	fmt.Print("\x1b[H")
}

func MoveCursorTo(x int, y int) {
	fmt.Printf("\x1b[%d;%df", y, x)
}

func SaveCursorPosition() {
	fmt.Print("\x1b[s")
}

func RestoreCursorPosition() {
	fmt.Print("\x1b[u")
}
