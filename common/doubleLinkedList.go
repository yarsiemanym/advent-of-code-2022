package common

import "log"

type DoubleLinkedListNode struct {
	value    interface{}
	next     *DoubleLinkedListNode
	previous *DoubleLinkedListNode
	list     *DoubleLinkedList
}

func (thisNode *DoubleLinkedListNode) Next() *DoubleLinkedListNode {
	return thisNode.next
}

func (thisNode *DoubleLinkedListNode) Previous() *DoubleLinkedListNode {
	return thisNode.previous
}

func (thisNode *DoubleLinkedListNode) Value() interface{} {
	return thisNode.value
}

type DoubleLinkedList struct {
	length int
	head   *DoubleLinkedListNode
	tail   *DoubleLinkedListNode
}

func (list *DoubleLinkedList) Length() int {
	return list.length
}

func (thisList *DoubleLinkedList) Head() *DoubleLinkedListNode {
	return thisList.head
}

func (thisList *DoubleLinkedList) Tail() *DoubleLinkedListNode {
	return thisList.tail
}

func (thisList *DoubleLinkedList) Append(value interface{}) {
	node := &DoubleLinkedListNode{
		value: value,
		list:  thisList,
	}

	if thisList.head == nil {
		thisList.head = node
		thisList.tail = node
	} else {
		node.previous = thisList.tail
		thisList.tail.next = node
		thisList.tail = node
	}

	thisList.length++
}

func (thisList *DoubleLinkedList) Prepend(value interface{}) {
	newNode := &DoubleLinkedListNode{
		value: value,
		list:  thisList,
	}

	if thisList.head == nil {
		thisList.head = newNode
		thisList.tail = newNode
	} else {
		newNode.next = thisList.head
		thisList.head.previous = newNode
		thisList.head = newNode
	}

	thisList.length++
}

func (thisList *DoubleLinkedList) InsertAfter(value interface{}, thisNode *DoubleLinkedListNode) {
	nextNode := thisNode.next
	newNode := &DoubleLinkedListNode{
		value:    value,
		next:     nextNode,
		previous: thisNode,
		list:     thisList,
	}
	thisNode.next = newNode
	nextNode.previous = newNode

	if newNode.next == nil {
		thisList.tail = newNode
	}

	thisList.length++
}

func (thisList *DoubleLinkedList) InsertBefore(value interface{}, thisNode *DoubleLinkedListNode) {
	previousNode := thisNode.previous
	newNode := &DoubleLinkedListNode{
		value:    value,
		next:     thisNode,
		previous: previousNode,
		list:     thisList,
	}
	thisNode.previous = newNode
	previousNode.next = newNode

	if newNode.previous == nil {
		thisList.head = newNode
	}

	thisList.length++
}

func (thisList *DoubleLinkedList) Remove(thisNode *DoubleLinkedListNode) {
	if thisNode.list != thisList {
		log.Panic("The specified DoubleLinkedListNode is not an element of this DoubleLinkedList.")
	}

	if thisNode.previous == nil {
		thisList.head = thisNode.next
	} else {
		thisNode.previous.next = thisNode.next
	}

	if thisNode.next == nil {
		thisList.tail = thisNode.previous
	} else {
		thisNode.next.previous = thisNode.previous
	}

	thisList.length--
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{
		length: 0,
		head:   nil,
		tail:   nil,
	}
}
