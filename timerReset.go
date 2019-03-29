package main

import (
	"fmt"
	"time"
)

func main() {
	const timeout = 100 * time.Millisecond
	t := time.NewTimer(timeout)
	reset := make(chan bool)

	go func() {
		select {
		case <-time.After(time.Duration(550) * time.Millisecond):
			fmt.Println("reset timer")
			reset <- true
		}
	}()

	for i := 1; i <= 10; i++ {
		start := time.Now()

		select {
		case <-reset:
			fmt.Println("triggered by reset.")

			elapsedTrig := time.Since(start)
			fmt.Println("elapsed t: ", elapsedTrig)

			t.Reset(timeout)
			continue
		case <-t.C:
			fmt.Println("triggered by timer.")
		}

		elapsed := time.Since(start)
		fmt.Printf("(%v) elapsed: %v \n", i, elapsed)

		t.Reset(timeout)
	}
}
