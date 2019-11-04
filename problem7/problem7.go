package problem7

import "strconv"

type codeString string

func (code codeString) isDecodable() bool {
	switch {
	case len(code) == 1:
		asInteger, err := strconv.Atoi(string(code))
		return err == nil && asInteger > 0 && asInteger < 10
	case len(code) == 2:
		asInteger, err := strconv.Atoi(string(code))
		return err == nil && asInteger > 9 && asInteger < 27
	default:
		return false
	}
}

func Run(input codeString) int {
	answer := 0
	switch {
	case len(input) == 1:
		if input.isDecodable() {
			answer += 1
		}
	case len(input) == 2:
		if input.isDecodable() {
			answer += 1
		}
		if input[:1].isDecodable() && input[1:].isDecodable() {
			answer += 1
		}
	case len(input) > 2:
		if input[:1].isDecodable() {
			answer += Run(input[1:])
		}
		if input[:2].isDecodable() {
			answer += Run(input[2:])
		}
	}
	return answer
}
