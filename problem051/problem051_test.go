package problem051

import (
	"fmt"
	"math"
	"testing"
)

func TestShuffle(t *testing.T) {
	const cardsCount = 52
	deck := make([]int, 0, cardsCount)
	for i := 0; i < cap(deck); i++ {
		deck = append(deck, i)
	}
	n, _ := whatAPerfectRandom(cardsCount)
	targetNumber := n - 1

	targetNumberAppearsFirst := 0
	loopCount := 10000
	for i := 0; i < loopCount; i++ {
		err := Shuffle(&deck)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
		if deck[0] == targetNumber {
			targetNumberAppearsFirst++
		}
	}

	if math.Abs(float64(loopCount)/float64(cardsCount)-float64(targetNumberAppearsFirst)) > float64(loopCount)/10.0 {
		t.Error(fmt.Sprintf("%d appears first %d out of %d times", targetNumber, targetNumberAppearsFirst, loopCount))
		t.Fail()
	}
}
