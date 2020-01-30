package problem035

import (
	"reflect"
	"testing"
)

func TestSegregate(t *testing.T) {
	input := []rune{'G', 'B', 'R', 'R', 'B', 'R', 'G', 'R', 'R', 'B', 'G', 'R', 'B'}
	Segregate(input)
	if !reflect.DeepEqual(input, []rune{'R', 'R', 'R', 'R', 'R', 'R', 'G', 'G', 'G', 'B', 'B', 'B', 'B'}) {
		t.Fail()
	}
}
