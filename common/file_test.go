package common

import (
	"strings"
	"testing"
)

func Test_ReadFile(t *testing.T) {
	path := "test.txt"
	text := ReadFile(path)

	if text != "1 blah\n2 foo\n" {
		t.Errorf("Expected \"1 blah\\n2 foo\\n\" but got \"%v\"", strings.Replace(text, "\n", "\\n", -1))
	}
}
