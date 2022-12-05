package common

type DoubleLinkedListNode struct {
	value    interface{}
	next     *DoubleLinkedListNode
	previous *DoubleLinkedListNode
}

func (node *DoubleLinkedListNode) Next() *DoubleLinkedListNode {
	return node.next
}

func (node *DoubleLinkedListNode) Previous() *DoubleLinkedListNode {
	return node.previous
}

func (node *DoubleLinkedListNode) Value() interface{} {
	return node.value
}

type DoubleLinkedList struct {
	length int
	head   *DoubleLinkedListNode
	tail   *DoubleLinkedListNode
}

func (list *DoubleLinkedList) Length() int {
	return list.length
}

func (list *DoubleLinkedList) Head() *DoubleLinkedListNode {
	return list.head
}

func (list *DoubleLinkedList) Tail() *DoubleLinkedListNode {
	return list.tail
}

func (list *DoubleLinkedList) Append(value interface{}) {
	node := &DoubleLinkedListNode{
		value: value,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		node.previous = list.tail
		list.tail.next = node
		list.tail = node
	}

	list.length++
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}
