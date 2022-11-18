package common

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

const PEEK_MAX_DEFAULT = 20

func Peek(text string, max int) string {
	if max < 0 {
		max = 0
	}

	if len(text) <= max {
		return text
	} else {
		return text[0:max] + "..."
	}
}

func Split(text string, delimiter string) []string {
	log.Debugf("Splitting text on \"%v\".", strings.Replace(delimiter, "\n", "\\n", -1))
	log.Debugf("Text is %v characters long.", len(text))
	log.Tracef("Text is \"%v\".", Peek(text, PEEK_MAX_DEFAULT))
	tokens := strings.Split(text, delimiter)
	log.Debugf("Resulting in %v tokens.", len(tokens))

	if log.GetLevel() == log.TraceLevel {
		for index, token := range tokens {
			log.Tracef("tokens[%v] = \"%v\"", index, Peek(token, PEEK_MAX_DEFAULT))
		}
	}

	return tokens
}
