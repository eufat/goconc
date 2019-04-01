package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Timeout int
	Reset   chan bool
}

func New() Timer {
	Timeout := 100
	Reset := make(chan bool)

	t := Timer{
		Timeout,
		Reset,
	}

	return t
}

func (t *Timer) StartTimer() {
	timer := time.NewTimer(time.Duration(t.Timeout) * time.Millisecond)

	go func() {
		for {
			select {
			case <-t.Reset:
				timer.Reset(time.Duration(t.Timeout) * time.Millisecond)
				continue
			case <-timer.C:
				fmt.Println("timer trigerred.")
			}
		}
	}()
}
