package problem028

import (
	"reflect"
	"testing"
)

func TestJustify(t *testing.T) {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	result := Justify(words, 16)
	if !reflect.DeepEqual(result, []string{"the  quick brown", "fox  jumps  over", "the   lazy   dog"}) {
		t.Fail()
		t.Log(result)
	}
}
