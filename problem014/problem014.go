package problem014

import (
	"math"
	"math/rand"
	"time"
)

type tester func(float64, float64) bool
type generator func() float64

func Run(totalRepeat int) float64 {
	const concurrency = 8
	rand.Seed(time.Now().UnixNano())
	circleTester := func(x float64, y float64) bool { return math.Pow(x, 2)+math.Pow(y, 2) < 1 }
	xGenerator := func() float64 { return rand.Float64() }
	subResults := make(chan struct{ successCount, trialCount int }, concurrency)
	for i := 0; i < concurrency; i++ {
		margin := i
		yGenerator := func() float64 { return (float64(margin) + rand.Float64()) / float64(concurrency) }
		go func(testerFunction tester, xGenerator generator, yGenerator generator, repeat int) {
			successCount, trialCount := simulate(circleTester, xGenerator, yGenerator, repeat)
			subResults <- struct{ successCount, trialCount int }{successCount: successCount, trialCount: trialCount}
		}(circleTester, xGenerator, yGenerator, totalRepeat/concurrency)
	}
	finished := 0
	totalSuccessCount, totalTrialCount := 0, 0
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

func simulate(testerFunction tester, xGenerator generator, yGenerator generator, repeat int) (int, int) {
	successCount, trialCount := 0, 0
	for i := 0; i < repeat; i++ {
		if testerFunction(xGenerator(), yGenerator()) {
			successCount++
		}
		trialCount++
	}
	return successCount, trialCount
}
