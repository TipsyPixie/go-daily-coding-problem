package problem2

import (
    "reflect"
    "testing"
)

func TestRun(t *testing.T) {
    if !reflect.DeepEqual(Run([]int{1, 2, 3, 4, 5}), []int{120, 60, 40, 30, 24}) {
        t.Fail()
    }
    if !reflect.DeepEqual(Run([]int{3, 2, 1}), []int{2, 3, 6}) {
        t.Fail()
    }
}
