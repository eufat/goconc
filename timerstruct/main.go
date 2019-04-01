package main

import (
	"fmt"
	"time"

	Timer "github.com/goconc/timerstruct/timer"
)

func main() {
	aTimer := Timer.New()

	fmt.Println("starting timer.")
	aTimer.StartTimer()

	go func() {
		select {
		case <-time.After(time.Duration(10000) * time.Millisecond):
			fmt.Println("resetting timer.")
			aTimer.Reset <- true
		}
	}()

	for {
	}
}
