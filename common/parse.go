package common

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type ParseFunc func(string) interface{}

func ParseToken(token string, fn ParseFunc) interface{} {
	log.Debug("Parsing token.")
	log.Debugf("Token is %v characters long.", len(token))
	log.Tracef("Token is \"%v\".", Peek(token, PEEK_MAX_DEFAULT))
	result := fn(token)
	return result
}

func ParseText(text string, delimiter string, fn ParseFunc) []interface{} {
	log.Debug("Parsing text.")
	log.Debugf("Text is %v characters long.", len(text))
	log.Tracef("Text is \"%v\".", Peek(text, PEEK_MAX_DEFAULT))

	tokens := Split(text, delimiter)
	var results []interface{}

	for _, token := range tokens {
		token := strings.Trim(token, " \n\t")
		result := ParseToken(token, fn)

		if result != nil {
			results = append(results, result)
		}
	}

	return results
}

func ParseFile(path string, delimiter string, fn ParseFunc) []interface{} {
	text := ReadFile(path)
	results := ParseText(text, delimiter, fn)
	return results
}
