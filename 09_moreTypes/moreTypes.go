package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

//Struct literals
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} //Y = 0 implicit
	v3 = Vertex{}     //X = 0 and Y = 0 implicit
	p1 = &Vertex{1, 2}
)

func pointers() {
	i, j := 42, 2701

	p := &i //Zeiger auf i
	fmt.Println(p)
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4] //indizes 1-3
	fmt.Println(s)
}

func main() {
	fmt.Println("Calling pointers()")
	pointers()

	fmt.Println("Structs:")
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	//Pointers to structs
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, p1, v2, v3)

	//Arrays
	fmt.Println("Calling arrays")
	arrays()

	//Slices
	fmt.Println("Calling Slices:")
	slices()
}
