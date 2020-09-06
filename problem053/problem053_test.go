package problem053

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue(16)
	for i := 0; i < 8; i++ {
		q.Enqueue(i)
	}

	for i := 0; i < 8; i++ {
		if v := q.Dequeue(); v != i {
			t.Logf("%d != %d", v, i)
			t.Fail()
		}
	}
}
