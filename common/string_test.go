package common

import "testing"

func Test_Peek_LenLessThanMax(t *testing.T) {
	text := "1234567890"
	result := Peek(text, 20)

	if result != text {
		t.Errorf("Expected \"%v\" but got \"%v\".", text, result)
	}
}

func Test_Peek_LenEqualToMax(t *testing.T) {
	text := "1234567890"
	result := Peek(text, 10)

	if result != text {
		t.Errorf("Expected \"%v\" but got \"%v\".", text, result)
	}
}

func Test_Peek_LenEqualToGreaterThanMax(t *testing.T) {
	text := "1234567890"
	expected := "12345..."
	result := Peek(text, 5)

	if result != expected {
		t.Errorf("Expected \"%v\" but got \"%v\".", expected, result)
	}
}

func Test_Split_NoDelimiters(t *testing.T) {
	text := "Line1"
	tokens := Split(text, "\n")

	if tokens == nil {
		t.Error("Returned nil.")
	}

	if len(tokens) != 1 {
		t.Errorf("Expected 1 tokens but got %v tokens.", len(tokens))
	}

	if tokens[0] != "Line1" {
		t.Errorf("Expected `['Line1']` but got `[\"%v\"]`.", tokens[0])
	}
}

func Test_Split_WithDelimiters(t *testing.T) {
	text := "Line1\nLine2"
	tokens := Split(text, "\n")

	if tokens == nil {
		t.Error("Returned nil.")
	}

	if len(tokens) != 2 {
		t.Errorf("Expected 2 tokens but got %v tokens.", len(tokens))
	}

	if tokens[0] != "Line1" || tokens[1] != "Line2" {
		t.Errorf("Expected `['Line1', 'Line2']` but got `[\"%v\", \"%v\"]`.", tokens[0], tokens[1])
	}
}
