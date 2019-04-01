package main

import "fmt"

func main() {
	go func() {
		for {
			fmt.Println("inside go routine")
		}
	}()

	fmt.Println("outside go routine")
}
