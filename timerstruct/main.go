package main

import (
	"fmt"
	"time"

	Timer "github.com/goconc/timerstruct/timer"
)

func main() {
	aTimer := Timer.New()

	fmt.Println("starting timer.")
	go aTimer.StartTimer()

	go func() {
		select {
		case <-time.After(time.Duration(800) * time.Millisecond):
			fmt.Println("reset timer")
			aTimer.Reset <- true
		}
	}()

	for {
	}
}
