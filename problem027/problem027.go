package problem027

type stack struct {
	buffer []rune
}

func newStack() *stack {
	return &stack{buffer: []rune{}}
}

func (s *stack) push(value rune) *stack {
	s.buffer = append(s.buffer, value)
	return s
}

func (s *stack) pop() rune {
	value := s.buffer[len(s.buffer)-1]
	s.buffer = s.buffer[:len(s.buffer)-1]
	return value
}

func (s *stack) isEmpty() bool {
	return len(s.buffer) == 0
}

func bracketMatched(bracket rune, pair rune) bool {
	switch pair {
	case ')':
		return bracket == '('
	case '}':
		return bracket == '{'
	case ']':
		return bracket == '['
	default:
		return false
	}
}

func ValidateBrackets(brackets string) bool {
	bracketStack := newStack()
	for _, bracket := range []rune(brackets) {
		switch bracket {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			bracketStack.push(bracket)
		case ')':
			fallthrough
		case '}':
			fallthrough
		case ']':
			if bracketStack.isEmpty() {
				return false
			}
			if !bracketMatched(bracketStack.pop(), bracket) {
				return false
			}
		default:
			return false
		}
	}
	return bracketStack.isEmpty()
}
