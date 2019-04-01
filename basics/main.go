package main

import "fmt"

func printSome() {
	fmt.Println("inside go routine")
}

func main() {
	go printSome()
	fmt.Println("outside go routine")

	for {
	}
}
