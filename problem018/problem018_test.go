package problem018

import (
	"reflect"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run([]int{10, 5, 2, 7, 8, 7}, 3)
	expectedResult := []int{10, 7, 8, 8}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Log(result)
		t.FailNow()
	}
}
