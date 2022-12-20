package vt100

import (
	"fmt"
	"strings"
)

func Sprint(text string, attributes ...string) string {
	attributeString := strings.Join(attributes, ";")
	output := fmt.Sprintf("\x1b[%sm%s\x1b[%sm", attributeString, text, ResetAttributes)
	return output
}

func Sprintf(format string, values []any, attributes ...string) string {
	text := fmt.Sprintf(format, values...)
	return Sprint(text, attributes...)
}

func Sprintln(text string, attributes ...string) string {
	text += "\n"
	return Sprint(text, attributes...)
}

func Printf(format string, values []any, attributes ...string) {
	text := fmt.Sprintf(format, values...)
	Print(text, attributes...)
}

func Print(text string, attributes ...string) {
	fmt.Print(Sprint(text, attributes...))
}

func Println(text string, attributes ...string) {
	text += "\n"
	Print(text, attributes...)
}
