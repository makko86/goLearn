package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

//buffered Channels block only when the buffer is full.
func bufferedChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//ch <- 3 blocks and deadlocks
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x //send x to channel
		x, y = y, x+y
	}
	//close the channel to tell the receiver we finished sending
	close(c)
}

//select waits on multiple channels and locks until anyof its cases can run, choosing one at random if multiple are ready
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return //exit the loop
		default:
			fmt.Println("Nothing to see here")
		}
	}
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //[7 2 8]
	go sum(s[len(s)/2:], c) //[-9 4 0]
	x, y := <-c, <-c        //receive fom c
	//send and receive block until the other side is ready
	fmt.Println(x, y, x+y)

	//buffered Channels block when buffer is full
	bufferedChan()
	fmt.Println()

	//c1 will be closed by sender when done => for range will loop until c1 is closed
	//only the sender should close a channel!
	c1 := make(chan int, 10)
	go fibonacci(cap(c1), c1)
	for i := range c1 {
		fmt.Println(i)
	}

	c2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c2)
		}
		quit <- 0
	}()
	fibonacciSelect(c2, quit)
}
