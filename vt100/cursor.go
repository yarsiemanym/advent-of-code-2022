package vt100

import "fmt"

func MoveCursorToHome() {
	fmt.Print("\x1b[H")
}

func MoveCursorTo(x int, y int) {
	fmt.Printf("\x1b[%d;%df", y, x)
}

func MoveCursorUp(y int) {
	fmt.Printf("\x1b[%dA", y)
}

func MoveCursorDown(y int) {
	fmt.Printf("\x1b[%dB", y)
}

func MoveCursorRight(x int) {
	fmt.Printf("\x1b[%dC", x)
}

func MoveCursorLeft(x int) {
	fmt.Printf("\x1b[%dD", x)
}

func SaveCursorPosition() {
	fmt.Print("\x1b[s")
}

func RestoreCursorPosition() {
	fmt.Print("\x1b[u")
}
