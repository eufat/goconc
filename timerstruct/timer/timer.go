package timer

import (
	"fmt"
	"time"
)

type Timer struct {
	Reset chan bool
}

func New() Timer {
	Reset := make(chan bool)

	t := Timer{
		Reset,
	}

	return t
}

func (t *Timer) StartTimer() {
	const timeout = 100 * time.Millisecond
	tim := time.NewTimer(timeout)

	for {
		start := time.Now()

		select {
		case <-t.Reset:
			fmt.Println("triggered by reset.")

			elapsedTrig := time.Since(start)
			fmt.Println("elapsed t: ", elapsedTrig)

			tim.Reset(timeout)
			continue
		case <-tim.C:
			fmt.Println("triggered by timer.")
		}

		elapsed := time.Since(start)
		fmt.Printf("elapsed: %v \n", elapsed)

		tim.Reset(timeout)
	}
}
