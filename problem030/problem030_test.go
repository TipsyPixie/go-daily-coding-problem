package problem030

import "testing"

func TestMeasureTrappedWater(t *testing.T) {
	testResult := MeasureTrappedWater([]int{2, 1, 2})
	if testResult != 1 {
		t.Fail()
	}

	testResult = MeasureTrappedWater([]int{1, 5, 2})
	if testResult != 0 {
		t.Fail()
	}

	testResult = MeasureTrappedWater([]int{3, 0, 1, 3, 0, 5})
	if testResult != 8 {
		t.Fail()
	}
}
