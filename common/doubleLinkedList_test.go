package common

import "testing"

func Test_DoubleLinkedList_AppendToEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)

	head := list.Head()
	tail := list.Tail()

	if head == nil {
		t.Error("list.Head() returned nil.")
	} else if head.Value() == nil {
		t.Error("head.Value() returned nil.")
	} else if head.Value() != 1 {
		t.Errorf("Expected 1 but got %v.", head.Value())
	} else if head.Previous() != nil {
		t.Error("head.Previous() did not returned nil.")
	} else if head.Next() != nil {
		t.Error("head.Next() did not returned nil.")
	}

	if tail == nil {
		t.Error("list.Tail() returned nil.")
	} else if tail.Value() == nil {
		t.Error("tail.Value() returned nil.")
	} else if tail.value != 1 {
		t.Errorf("Expected 1 but got %v.", tail.Value())
	} else if tail.Previous() != nil {
		t.Error("tail.Previous() did not returned nil.")
	} else if tail.Next() != nil {
		t.Error("tail.Next() did not returned nil.")
	}

	if head != tail {
		t.Error("head and tail are different nodes.")
	}
}

func Test_DoubleLinkedList_AppendToNotEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	head := list.Head()
	tail := list.Tail()

	if head == nil {
		t.Error("list.Head() returned nil.")
	} else if head.Value() == nil {
		t.Error("head.Value() returned nil.")
	} else if head.Value() != 1 {
		t.Errorf("Expected 1 but got %v.", head.Value())
	} else if head.Previous() != nil {
		t.Error("head.Previous() did not returned nil.")
	} else if head.Next() == nil {
		t.Error("head.Next() returned nil.")
	}

	if tail == nil {
		t.Error("list.Tail() returned nil.")
	} else if tail.Value() == nil {
		t.Error("tail.Value() returned nil.")
	} else if tail.value != 3 {
		t.Errorf("Expected 3 but got %v.", tail.Value())
	} else if tail.Previous() == nil {
		t.Error("tail.Previous() returned nil.")
	} else if tail.Next() != nil {
		t.Error("tail.Next() did not returned nil.")
	}

	if head.Next() != tail.Previous() {
		t.Error("head and tail are not linked by the same body.")
	} else if head.Next().Value() == nil {
		t.Error("body.Value() returned nil.")
	} else if head.Next().Value() != 2 {
		t.Errorf("Expected 2 but got %v.", head.Next().Value())
	}
}
