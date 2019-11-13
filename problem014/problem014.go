package problem014

import (
	"math"
	"math/rand"
	"time"
)

type tester func(float64, float64) bool

func Run(concurrency int, totalRepeat int) float64 {
	subResults := make(chan struct{ successCount, trialCount int }, concurrency)
	isInsideCircle := func(x float64, y float64) bool { return math.Pow(x, 2)+math.Pow(y, 2) < 1 }
	repeat := totalRepeat / concurrency
	for i := 0; i < concurrency; i++ {
		go func() {
			successCount, trialCount := simulate(isInsideCircle, repeat)
			subResults <- struct{ successCount, trialCount int }{successCount: successCount, trialCount: trialCount}
		}()
	}
	totalSuccessCount, totalTrialCount, finished := 0, 0, 0
	for finished < concurrency {
		select {
		case result, _ := <-subResults:
			totalSuccessCount += result.successCount
			totalTrialCount += result.trialCount
			finished++
		}
	}
	approxPi := float64(totalSuccessCount*4) / float64(totalTrialCount)
	return math.Round(1000*approxPi) / 1000
}

func simulate(criteria tester, loop int) (int, int) {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	successCount, trialCount := 0, 0
	for i := 0; i < loop; i++ {
		trialCount++
		if criteria(generator.Float64(), generator.Float64()) {
			successCount++
		}
	}
	return successCount, trialCount
}
