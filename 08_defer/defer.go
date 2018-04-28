package main

import "fmt"

func count() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func main() {
	//defer arguments evaluated immediately, call executed on return
	defer fmt.Println("world!")

	fmt.Println("Hello")

	count()
}
