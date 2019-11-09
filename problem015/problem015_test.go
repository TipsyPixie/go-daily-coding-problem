package problem015

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	arg := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultSet := make(map[byte]int)
	const loop = 10000
	for i := 0; i < loop; i++ {
		result, err := Run(bytes.NewReader(arg))
		//t.Log(result)
		if err != nil {
			t.FailNow()
		}
		resultSet[result]++
	}
	for _, element := range arg {
		noLessThan := loop * (100 - 5*len(arg)) / (100 * len(arg))
		noMoreThan := loop * (100 + 5*len(arg)) / (100 * len(arg))
		if resultSet[element] < noLessThan || resultSet[element] > noMoreThan {
			t.Fail()
		}
	}
}
