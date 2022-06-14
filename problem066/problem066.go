package problem066

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type resultPair struct {
	first  int
	second int
}

func biasedToss() int {
	if rand.Int()%10 == 0 {
		return 0
	}
	return 1
}

func FairToss() int {
	tossResults := make(chan *resultPair, 10)

	const tossCount int = 2 * 5
	for {
		for i := 0; i < tossCount; i++ {
			go func() {
				tossResults <- &resultPair{biasedToss(), biasedToss()}
			}()
		}

		for i := 0; i < tossCount; i++ {
			select {
			case tossResult := <-tossResults:
				if tossResult.first != tossResult.second {
					return tossResult.first
				}
			}
		}
	}
}
