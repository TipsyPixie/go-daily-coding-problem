package problem10

import (
    "testing"
)

func TestSchedule(t *testing.T) {
    ok := Schedule(func() interface{} { return 1 }, 3000)
    ok2 := Schedule(func() interface{} { return 2 }, 1000)
    finishedJobs := 0
    for i := 0; i < 2; i++ {
        select {
        case x := <-ok:
            if x != 1 {
                t.FailNow()
            }
            finishedJobs++
        case x := <-ok2:
            if x != 2 {
                t.FailNow()
            }
            finishedJobs++
        }
    }
    if finishedJobs != 2 {
        t.FailNow()
    }
}
