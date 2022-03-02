package problem061

import (
	"testing"
)

type argumentSet struct {
	base     int
	exponent int
	result   int
}

func assert(t *testing.T, actual int, expected int) bool {
	t.Logf("Actual: %d / Expected: %d", actual, expected)
	if actual != expected {
		return false
	}
	return true
}

func getBenchmark(pow func(int, int) int, base int, exponent int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pow(base, exponent)
		}
	}
}

func TestPow(t *testing.T) {
	testCases := []argumentSet{
		{
			base:     8,
			exponent: 19,
		},
	}

	for _, testCase := range testCases {
		if !assert(t, PowFaster(testCase.base, testCase.exponent), PowSimple(testCase.base, testCase.exponent)) {
			t.FailNow()
		}
		testing.Benchmark(getBenchmark(PowFaster, testCase.base, testCase.exponent))
		testing.Benchmark(getBenchmark(PowSimple, testCase.base, testCase.exponent))
	}
}
