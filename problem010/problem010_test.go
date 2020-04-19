package problem010

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	ok := Schedule(func() interface{} { return 1 }, 1000)
	ok2 := Schedule(func() interface{} { return 2 }, 500)
	finishedJobs := 0
	for i := 0; i < 2; i++ {
		select {
		case x := <-ok:
			if x != 1 || finishedJobs != 1 {
				t.FailNow()
			}
			finishedJobs++
		case x := <-ok2:
			if x != 2 || finishedJobs != 0 {
				t.FailNow()
			}
			finishedJobs++
		case <-time.After(2000 * time.Millisecond):
			t.Log("Timeout")
			t.FailNow()
		}
	}
	if finishedJobs != 2 {
		t.Log(fmt.Sprintf("%d job(s) failed", 2-finishedJobs))
		t.FailNow()
	}
}
