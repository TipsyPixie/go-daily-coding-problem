package problem10

import "time"

type Job func() interface{}

func Schedule(f Job, n int) <-chan interface{} {
	returnChannel := make(chan interface{}, 1)
	go func() {
		select {
		case <-time.After(time.Duration(n) * time.Millisecond):
			returnChannel <- f()
		}
	}()
	return returnChannel
}
