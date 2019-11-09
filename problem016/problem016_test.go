package problem016

import "testing"

func TestOrderIdLogger(t *testing.T) {
	logger := NewOrderIdLogger(7)
	logger.Record("first")
	logger.Record("second")
	logger.Record("third")
	logger.Record("fourth")
	logger.Record("fifth")
	logger.Record("sixth")
	logger.Record("seventh")
	logger.Record("eighth")
	logger.Record("ninth")
	logger.Record("tenth")
	if logger.GetLast(1) != "tenth" || logger.GetLast(7) != "fourth" {
		t.FailNow()
	}
}
