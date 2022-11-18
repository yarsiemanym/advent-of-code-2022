package common

type Stack struct {
	elements []interface{}
	size     int
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
		size:     0,
	}
}

func (stack *Stack) Size() int {
	return stack.size
}

func (stack *Stack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *Stack) Push(element interface{}) {
	stack.elements = append(stack.elements, element)
	stack.size++
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}

	size := stack.Size()
	element := stack.elements[size-1]
	stack.elements = stack.elements[:size-1]
	stack.size--

	return element
}
