package problem042

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSubset(t *testing.T) {
	values := []int{12, 1, 61, 5, 9, 2}
	k := 24
	answer, correctAnswer := GetSubset(values, k), []int{12, 1, 9, 2}
	if !reflect.DeepEqual(answer, correctAnswer) {
		t.Log(fmt.Sprintf("%v, %d: %v != %v", values, k, answer, correctAnswer))
		t.FailNow()
	}
	j := 61
	answer2, correctAnswer2 := GetSubset(values, j), []int{61}
	if !reflect.DeepEqual(answer, correctAnswer) {
		t.Log(fmt.Sprintf("%v, %d: %v != %v", values, j, answer2, correctAnswer2))
		t.FailNow()
	}
}
