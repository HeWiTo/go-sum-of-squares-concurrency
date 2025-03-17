package main

import "fmt"

func sumOfSquares(n int, squares chan int, quit chan bool) {
	i := 1

	for {
		select {
		case squares <- i * i:
			i++
			if i > n {
				quit <- true
				return
			}
		case <-quit:
			return
		}
	}
}

func main() {
	n := 5
	squares := make(chan int)
	quit := make(chan bool)

	// Run calculations concurrently
	go sumOfSquares(n, squares, quit)

	sum := 0
	// Take exactly n values from the channel squares and sum them up
	for i := 1; i <= n; i++ {
		sum += <-squares
	}
	// Make sure the goroutine sends a finish signal
	<-quit

	fmt.Printf("The sum of the squares from 1 to %d is %d\n", n, sum)
}
