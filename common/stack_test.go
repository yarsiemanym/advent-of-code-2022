package common

import "testing"

func Test_stack_Push_Empty(t *testing.T) {
	stack := &Stack{
		elements: []interface{}{},
		size:     0,
	}

	stack.Push(5)

	if stack.size != 1 {
		t.Errorf("Expected 1 but got %v.", stack.size)
	}

	if len(stack.elements) != 1 {
		t.Errorf("Expected 1 but got %v.", len(stack.elements))
	}

	if stack.elements[0] != 5 {
		t.Errorf("Expected 5 but got %v.", stack.elements[0])
	}
}

func Test_stack_Push_NotEmpty(t *testing.T) {
	stack := &Stack{
		elements: []interface{}{4},
		size:     1,
	}

	stack.Push(5)

	if stack.size != 2 {
		t.Errorf("Expected 2 but got %v.", stack.size)
	}

	if len(stack.elements) != 2 {
		t.Errorf("Expected 2 but got %v.", len(stack.elements))
	}

	if stack.elements[0] != 4 {
		t.Errorf("Expected 4 but got %v.", stack.elements[0])
	}

	if stack.elements[1] != 5 {
		t.Errorf("Expected 5 but got %v.", stack.elements[1])
	}
}

func Test_stack_Pop_Empty(t *testing.T) {
	stack := &Stack{
		elements: []interface{}{},
		size:     0,
	}

	value := stack.Pop()

	if value != nil {
		t.Error("value is not nil.")
	}

	if stack.size != 0 {
		t.Errorf("Expected 0 but got %v.", stack.size)
	}

	if len(stack.elements) != 0 {
		t.Errorf("Expected 0 but got %v.", len(stack.elements))
	}
}

func Test_stack_Pop_NotEmpty(t *testing.T) {
	stack := &Stack{
		elements: []interface{}{4},
		size:     1,
	}

	value := stack.Pop()

	if value == nil {
		t.Error("value is nil.")
	} else if value.(int) != 4 {
		t.Errorf("Expected 4 but got %v.", value.(int))
	}

	if stack.size != 0 {
		t.Errorf("Expected 0 but got %v.", stack.size)
	}

	if len(stack.elements) != 0 {
		t.Errorf("Expected 0 but got %v.", len(stack.elements))
	}
}
