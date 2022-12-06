package common

import "testing"

func Test_DoubleLinkedList_Append_EmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 1 {
		t.Errorf("Expected 1 but got %v.", list.Length())
	}

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

func Test_DoubleLinkedList_Append_NotEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)
	list.Append(2)
	list.Append(3)

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 3 {
		t.Errorf("Expected 3 but got %v.", list.Length())
	}

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

func Test_DoubleLinkedList_Prepend_EmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Prepend(1)

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 1 {
		t.Errorf("Expected 1 but got %v.", list.Length())
	}

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

func Test_DoubleLinkedList_Prepend_NotEmptyList(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Prepend(1)
	list.Prepend(2)
	list.Prepend(3)

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 3 {
		t.Errorf("Expected 3 but got %v.", list.Length())
	}

	if head == nil {
		t.Error("list.Head() returned nil.")
	} else if head.Value() == nil {
		t.Error("head.Value() returned nil.")
	} else if head.Value() != 3 {
		t.Errorf("Expected 3 but got %v.", head.Value())
	} else if head.Previous() != nil {
		t.Error("head.Previous() did not returned nil.")
	} else if head.Next() == nil {
		t.Error("head.Next() returned nil.")
	}

	if tail == nil {
		t.Error("list.Tail() returned nil.")
	} else if tail.Value() == nil {
		t.Error("tail.Value() returned nil.")
	} else if tail.value != 1 {
		t.Errorf("Expected 1 but got %v.", tail.Value())
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

func Test_DoubleLinkedList_InsertAfter(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)
	list.Append(2)
	list.InsertAfter(3, list.Head())

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 3 {
		t.Errorf("Expected 3 but got %v.", list.Length())
	}

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
	} else if tail.value != 2 {
		t.Errorf("Expected 2 but got %v.", tail.Value())
	} else if tail.Previous() == nil {
		t.Error("tail.Previous() returned nil.")
	} else if tail.Next() != nil {
		t.Error("tail.Next() did not returned nil.")
	}

	if head.Next() != tail.Previous() {
		t.Error("head and tail are not linked by the same body.")
	} else if head.Next().Value() == nil {
		t.Error("body.Value() returned nil.")
	} else if head.Next().Value() != 3 {
		t.Errorf("Expected 3 but got %v.", head.Next().Value())
	}
}

func Test_DoubleLinkedList_InsertBefore(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)
	list.Append(2)
	list.InsertBefore(3, list.Tail())

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 3 {
		t.Errorf("Expected 3 but got %v.", list.Length())
	}

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
	} else if tail.value != 2 {
		t.Errorf("Expected 2 but got %v.", tail.Value())
	} else if tail.Previous() == nil {
		t.Error("tail.Previous() returned nil.")
	} else if tail.Next() != nil {
		t.Error("tail.Next() did not returned nil.")
	}

	if head.Next() != tail.Previous() {
		t.Error("head and tail are not linked by the same body.")
	} else if head.Next().Value() == nil {
		t.Error("body.Value() returned nil.")
	} else if head.Next().Value() != 3 {
		t.Errorf("Expected 3 but got %v.", head.Next().Value())
	}
}

func Test_DoubleLinkedList_Remove(t *testing.T) {
	list := NewDoubleLinkedList()

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(list.Head().Next())

	head := list.Head()
	tail := list.Tail()

	if list.Length() != 2 {
		t.Errorf("Expected 2 but got %v.", list.Length())
	}

	if head == nil {
		t.Error("list.Head() returned nil.")
	} else if head.Value() == nil {
		t.Error("head.Value() returned nil.")
	} else if head.Value() != 1 {
		t.Errorf("Expected 1 but got %v.", head.Value())
	} else if head.Previous() != nil {
		t.Error("head.Previous() did not returned nil.")
	} else if head.Next() != tail {
		t.Error("head.Next() did not return tail.")
	}

	if tail == nil {
		t.Error("list.Tail() returned nil.")
	} else if tail.Value() == nil {
		t.Error("tail.Value() returned nil.")
	} else if tail.value != 3 {
		t.Errorf("Expected 3 but got %v.", tail.Value())
	} else if tail.Previous() != head {
		t.Error("tail.Previous() did not return head.")
	} else if tail.Next() != nil {
		t.Error("tail.Next() did not returned nil.")
	}
}
