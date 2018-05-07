package main

import "fmt"

//fibbonaci is a function tha returns
//a function that returns an int.
//values of n and v are stored in between calls
func fibbonaci() func() int {
	n, v := 1, 0
	return func() int {
		n, v = v, n+v
		return n
	}
}

func main() {
	f := fibbonaci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
