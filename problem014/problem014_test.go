package problem014

import (
	"math"
	"testing"
)

func TestRun(t *testing.T) {
	result := Run(1, 10000)
	if result < math.Pi*0.95 || result > math.Pi*1.05 {
		t.FailNow()
	}
}

func TestRun2(t *testing.T) {
	result := Run(8, 100000000)
	if result < math.Pi*0.999 || result > math.Pi*1.001 {
		t.FailNow()
	}
}
