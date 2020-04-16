package problem043

type stack struct {
	values        []int
	maxValueStack *stack
}

type StackEmptyError struct{}

func (stackEmptyError StackEmptyError) Error() string {
	return "StackEmptyError"
}

func NewStack(initialCapacity uint) *stack {
	return &stack{
		values: make([]int, 0, initialCapacity),
		maxValueStack: &stack{
			values:        make([]int, 0, initialCapacity),
			maxValueStack: nil,
		},
	}
}

func (thisStack *stack) peek() (int, error) {
	if len(thisStack.values) == 0 {
		return 0, StackEmptyError{}
	}
	return thisStack.values[len(thisStack.values)-1], nil
}

func (thisStack *stack) Push(value int) *stack {
	if thisStack.maxValueStack != nil {
		currentMaxValue, err := thisStack.maxValueStack.peek()
		if (err != nil && err.Error() == "StackEmptyError") || currentMaxValue <= value {
			thisStack.maxValueStack.Push(value)
		}
	}
	thisStack.values = append(thisStack.values, value)
	return thisStack
}

func (thisStack *stack) Pop() (int, error) {
	if len(thisStack.values) == 0 {
		return 0, StackEmptyError{}
	}
	value := thisStack.values[len(thisStack.values)-1]
	thisStack.values = thisStack.values[:len(thisStack.values)-1]

	if thisStack.maxValueStack != nil {
		currentMaxValue, _ := thisStack.maxValueStack.peek()
		if currentMaxValue == value {
			_, _ = thisStack.maxValueStack.Pop()
		}
	}
	return value, nil
}

func (thisStack *stack) Max() (int, error) {
	currentMaxValue, err := thisStack.maxValueStack.peek()
	if err != nil {
		return 0, err
	}
	return currentMaxValue, nil
}
