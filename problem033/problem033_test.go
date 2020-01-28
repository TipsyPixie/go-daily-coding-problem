package problem033

import (
	"reflect"
	"testing"
)

func TestGetMedians(t *testing.T) {
	sequence := []int{2, 1, 5, 7, 2, 0, 5}
	medians := GetMedians(sequence)
	if !reflect.DeepEqual(medians, []float32{2, 1.5, 2, 3.5, 2, 2, 2}) {
		t.Log(medians)
		t.Fail()
	}
}
