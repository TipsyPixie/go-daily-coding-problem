package problem1

import (
    "testing"
)

func TestRun(t *testing.T) {
    result := Run([]uint{10, 15, 3, 7}, 17)
    if !result {
        t.FailNow()
    }
}
