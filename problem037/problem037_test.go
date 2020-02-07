package problem037

import (
	"testing"
)

func TestSet_PowerSet(t *testing.T) {
	powerSet := NewSet(1, 2, 3).PowerSet()
	answer := NewSetOfSet(
		NewSet(1, 2, 3),
		NewSet(1, 2), NewSet(1, 3), NewSet(2, 3),
		NewSet(1), NewSet(2), NewSet(3),
		NewSet(),
	)
	if !powerSet.equal(answer) {
		t.Fail()
	}
}
