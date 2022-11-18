package common

type LinkedList struct {
	length int
	head   *node
	tail   *node
}

type node struct {
	value interface{}
	next  *node
}

func (list *LinkedList) Length() int {
	return list.length
}

func (list *LinkedList) Append(value interface{}) {
	node := &node{
		value: value,
	}

	if list.head == nil {
		list.head = node
		list.tail = node
	} else {
		list.tail.next = node
		list.tail = node
	}

	list.length++
}
