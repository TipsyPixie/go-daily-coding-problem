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
	for {
		tossResult := &resultPair{biasedToss(), biasedToss()}
		if tossResult.first != tossResult.second {
			return tossResult.first
		}
	}
}
