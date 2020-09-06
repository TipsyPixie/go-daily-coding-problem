package problem053

type stack struct {
	capacity int
	data     []int
}

func newStack(capacity int) *stack {
	return &stack{
		capacity: capacity,
		data:     make([]int, 0, capacity),
	}
}

func (thisStack *stack) isEmpty() bool {
	return len(thisStack.data) == 0
}

func (thisStack *stack) push(v int) *stack {
	if len(thisStack.data) >= thisStack.capacity {
		panic("max length exceeded")
	}
	thisStack.data = append(thisStack.data, v)
	return thisStack
}

func (thisStack *stack) pop() int {
	if len(thisStack.data) == 0 {
		panic("stack empty")
	}
	v := thisStack.data[len(thisStack.data)-1]
	thisStack.data = thisStack.data[:len(thisStack.data)-1]
	return v
}

type queue struct {
	incomingStack *stack
	outgoingStack *stack
}

func NewQueue(capacity int) *queue {
	return &queue{
		incomingStack: newStack(capacity),
		outgoingStack: newStack(capacity),
	}
}

func (thisQueue *queue) Enqueue(v int) *queue {
	for !thisQueue.outgoingStack.isEmpty() {
		thisQueue.incomingStack.push(thisQueue.outgoingStack.pop())
	}
	thisQueue.incomingStack.push(v)
	return thisQueue
}

func (thisQueue *queue) Dequeue() int {
	for !thisQueue.incomingStack.isEmpty() {
		thisQueue.outgoingStack.push(thisQueue.incomingStack.pop())
	}
	return thisQueue.outgoingStack.pop()

}
