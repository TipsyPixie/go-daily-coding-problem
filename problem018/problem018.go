package problem018

type queue struct {
	values []int
}

func NewQueue(size int) *queue {
	return &queue{values: make([]int, 0, size)}
}

func (q *queue) enqueue(value int) {
	q.values = append(q.values, value)
}
func (q *queue) dequeue() int {
	value := q.values[0]
	q.values = q.values[1:]
	return value
}

func (q *queue) pop() int {
	value := q.values[len(q.values)-1]
	q.values = q.values[:len(q.values)-1]
	return value
}

func (q *queue) isEmpty() bool {
	return len(q.values) == 0
}

func (q *queue) peekHead() int {
	return q.values[0]
}

func (q *queue) peekTail() int {
	return q.values[len(q.values)-1]
}

func Run(input []int, subArraySize int) []int {
	indexQueue := NewQueue(subArraySize + 1)
	result := make([]int, 0, len(input)/subArraySize)
	for index, value := range input {
		if !indexQueue.isEmpty() && indexQueue.peekHead() <= index-subArraySize {
			indexQueue.dequeue()
		}
		for !indexQueue.isEmpty() && input[indexQueue.peekTail()] <= value {
			indexQueue.pop()
		}
		indexQueue.enqueue(index)
		if index >= subArraySize-1 {
			result = append(result, input[indexQueue.peekHead()])
		}
	}
	return result
}
