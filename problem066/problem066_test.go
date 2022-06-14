package problem066

import "testing"

const leeway float32 = 0.1

func TestFairToss(t *testing.T) {
	for _, count := range []int{10000, 100000, 1000000} {
		headCount := 0
		for i := 0; i < count; i++ {
			if coin := FairToss(); coin == 0 {
				headCount += 1
			}
		}

		if headCount < int(float32(count)*(0.5-leeway)) || headCount > int(float32(count)*(0.5+leeway)) {
			t.Logf("%v heads out of %v", headCount, count)
			t.FailNow()
		}
	}
}
