package main

import (
	"fmt"
	"time"
)

func main() {
	go countNumbers() // start a new goroutine
	go countLetters() // start another new goroutine

	// wait for 2 seconds to allow the goroutines to run
	time.Sleep(2 * time.Second)
}

func countNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func countLetters() {
	for i := 'a'; i <= 'e'; i++ {
		fmt.Println("Letter:", string(i))
		time.Sleep(500 * time.Millisecond)
	}
}
