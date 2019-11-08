package problem014

import (
	"math"
	"math/rand"
	"time"
)

type tester func(float64, float64) bool

func Run(concurrency int, totalRepeat int) float64 {
	rand.Seed(time.Now().UnixNano())
	subResults := make(chan struct{ successCount, trialCount int }, concurrency)
	isInsideCircle := func(x float64, y float64) bool { return math.Pow(x, 2)+math.Pow(y, 2) < 1 }
	repeat := totalRepeat / concurrency
	xSamples, ySamples := sampleCoordinates(repeat * concurrency)
	for i := 0; i < concurrency; i++ {
		startsAt, endsAt := i*repeat, (i+1)*repeat
		go func() {
			successCount, trialCount := simulate(isInsideCircle, xSamples[startsAt:endsAt], ySamples[startsAt:endsAt], repeat)
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

func simulate(criteria tester, xValues []float64, yValues []float64, loop int) (int, int) {
	successCount, trialCount := 0, 0
	for i := 0; i < loop; i++ {
		trialCount++
		if criteria(xValues[i], yValues[i]) {
			successCount++
		}
	}
	return successCount, trialCount
}

func sampleCoordinates(loop int) ([]float64, []float64) {
	xValues := make([]float64, loop, loop)
	yValues := make([]float64, loop, loop)
	for i := 0; i < loop; i++ {
		xValues[i] = rand.Float64()
		yValues[i] = rand.Float64()
	}
	return xValues, yValues
}
