package common

type Queue struct {
	elements []interface{}
	size     int
}

func NewQueue() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func (queue *Queue) Size() int {
	return queue.size
}

func (queue *Queue) IsEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) Push(element interface{}) {
	queue.elements = append(queue.elements, element)
	queue.size++
}

func (queue *Queue) Pop() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	element := queue.elements[0]
	queue.elements = queue.elements[1:]
	queue.size--

	return element
}

func (queue *Queue) Peek() interface{} {
	if queue.IsEmpty() {
		return nil
	}

	element := queue.elements[0]

	return element
}
