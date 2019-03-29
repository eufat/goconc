package main

import (
	"fmt"
	"time"
)

func main() {
	const timeout = 100 * time.Millisecond
	t := time.NewTimer(timeout)

	for i := 1; i <= 10; i++ {
		start := time.Now()

		select {
		case <-t.C:
			fmt.Println("triggered by timer.")
		}

		elapsed := time.Since(start)
		fmt.Printf("(%v) elapsed: %v \n", i, elapsed)

		t.Reset(timeout)
	}
}
