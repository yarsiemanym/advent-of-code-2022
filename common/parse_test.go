package common

import (
	"strconv"
	"testing"
)

type Thingamajig struct {
	intProp    int
	stringProp string
}

func Parse(text string) interface{} {
	if text == "" {
		return nil
	}

	tokens := Split(text, " ")

	intProp, err := strconv.Atoi(tokens[0])
	Check(err)

	stringProp := tokens[1]

	result := Thingamajig{
		intProp:    intProp,
		stringProp: stringProp,
	}

	return result
}

func Test_ParseToken(t *testing.T) {
	text := "1 blah"
	result := ParseToken(text, Parse)

	if result == nil {
		t.Error("Returned nil.")
	}

	thing := result.(Thingamajig)

	if thing.intProp != 1 {
		t.Errorf("Expected `thingamajig.intProp = 1` but got `thingamajig.intProp = %v`.", thing.intProp)
	}

	if thing.stringProp != "blah" {
		t.Errorf("Expected `thingamajig.intProp = \"blah\"` but got `thingamajig.intProp = \"%v\"`.", thing.stringProp)
	}
}

func Test_ParseText(t *testing.T) {
	text := "1 blah\n2 foo\n"
	results := ParseText(text, "\n", Parse)

	if results == nil {
		t.Error("Returned nil.")
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results but got %v.", len(results))
	}

	thing := results[0].(Thingamajig)

	if thing.intProp != 1 {
		t.Errorf("Expected `thingamajig.intProp = 1` but got `thingamajig.intProp = %v`.", thing.intProp)
	}

	if thing.stringProp != "blah" {
		t.Errorf("Expected `thingamajig.intProp = \"blah\"` but got `thingamajig.intProp = \"%v\"`.", thing.stringProp)
	}

	thing = results[1].(Thingamajig)

	if thing.intProp != 2 {
		t.Errorf("Expected `thingamajig.intProp = 2` but got `thingamajig.intProp = %v`.", thing.intProp)
	}

	if thing.stringProp != "foo" {
		t.Errorf("Expected `thingamajig.intProp = \"foo\"` but got `thingamajig.intProp = \"%v\"`.", thing.stringProp)
	}
}

func Test_ParseFile(t *testing.T) {
	path := "test.txt"
	results := ParseFile(path, "\n", Parse)

	if results == nil {
		t.Error("Returned nil")
	}

	if len(results) != 2 {
		t.Errorf("Expected 2 results but got %v.", len(results))
	}

	thing := results[0].(Thingamajig)

	if thing.intProp != 1 {
		t.Errorf("Expected `thingamajig.intProp = 1` but got `thingamajig.intProp = %v`.", thing.intProp)
	}

	if thing.stringProp != "blah" {
		t.Errorf("Expected `thingamajig.intProp = \"blah\"` but got `thingamajig.intProp = \"%v\"`.", thing.stringProp)
	}

	thing = results[1].(Thingamajig)

	if thing.intProp != 2 {
		t.Errorf("Expected `thingamajig.intProp = 2` but got `thingamajig.intProp = %v`.", thing.intProp)
	}

	if thing.stringProp != "foo" {
		t.Errorf("Expected `thingamajig.intProp = \"foo\"` but got `thingamajig.intProp = \"%v\"`.", thing.stringProp)
	}
}
