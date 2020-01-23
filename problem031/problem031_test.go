package problem031

import (
	"testing"
)

func TestMeasureEditDistance(t *testing.T) {
	var editDistance int

	editDistance = MeasureEditDistance("sitting", "kitten")
	if editDistance != 3 {
		t.Fail()
		t.Log(editDistance)
	}

	editDistance = MeasureEditDistance("Saturday", "Sunday")
	if editDistance != 3 {
		t.Fail()
		t.Log(editDistance)
	}
}
