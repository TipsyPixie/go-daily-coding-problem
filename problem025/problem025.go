package problem025

type state struct {
	success      bool
	defaultShift *state
	shiftMap     map[rune]*state
}

func (currentState *state) defaultTo(nextState *state) *state {
	currentState.defaultShift = nextState
	return currentState
}

func (currentState *state) register(indicator rune, nextState *state) *state {
	currentState.shiftMap[indicator] = nextState
	return currentState
}

func (currentState *state) shift(indicator rune) *state {
	if currentState.shiftMap[indicator] == nil {
		return currentState.defaultShift
	}
	return currentState.shiftMap[indicator]
}

func newState(success bool) *state {
	return &state{
		success:      success,
		defaultShift: nil,
		shiftMap:     make(map[rune]*state),
	}
}

func compileRegex(regex string) *state {
	characters := []rune(regex)
	doomedState, blessedState := newState(false), newState(true)
	doomedState.defaultTo(doomedState)
	blessedState.defaultTo(blessedState)
	initialState := newState(false).defaultTo(doomedState)
	cursor := initialState
	for index, character := range characters {
		if index == len(characters)-1 {
			switch character {
			case '.':
				cursor.defaultTo(newState(true).defaultTo(doomedState))
			case '*':
				cursor.success = true
			default:
				cursor.defaultTo(doomedState)
				cursor.register(character, newState(true).defaultTo(doomedState))
			}
		} else {
			nextGoodState := newState(false)
			if nextCharacter := characters[index+1]; nextCharacter == '*' {
				nextGoodState = cursor
			}
			switch character {
			case '.':
				cursor.defaultTo(nextGoodState)
				cursor = nextGoodState
			case '*':
			default:
				cursor.register(character, nextGoodState)
				cursor = nextGoodState
			}
		}
	}

	return initialState
}

func TestRegex(regex string, value string) bool {
	fsm := compileRegex(regex)
	characters := []rune(value)
	for _, character := range characters {
		fsm = fsm.shift(character)
	}
	return fsm.success
}
