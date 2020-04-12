package problem041

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFullItinerary(t *testing.T) {
	testFlights := [][][]string{
		{{"SFO", "HKO"}, {"YYZ", "SFO"}, {"YUL", "YYZ"}, {"HKO", "ORD"}},
		{{"SFO", "COM"}, {"COM", "YYZ"}},
		{{"A", "B"}, {"A", "C"}, {"B", "C"}, {"C", "A"}},
	}
	startingAiportCodes := []string{"YUL", "COM", "A"}
	correctAnswers := [][]string{
		{"YUL", "YYZ", "SFO", "HKO", "ORD"},
		nil,
		nil,
	}
	for i := range testFlights {
		result := FullItinerary(testFlights[i], startingAiportCodes[i])
		if !reflect.DeepEqual(result, correctAnswers[i]) {
			t.Log(fmt.Sprintf("%s: %s != %s", testFlights[i], result, correctAnswers[i]))
			t.FailNow()
		}
	}
}
