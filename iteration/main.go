package main

import "fmt"

func main() {
	iterationChannel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			iterationChannel <- i
		}

		close(iterationChannel)
	}()

	for i := range iterationChannel {
		fmt.Println(i)
	}
}
