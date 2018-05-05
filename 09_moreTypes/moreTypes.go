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

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//defaults lowbound = 0 highbound = length of slice
	s1 := []int{2, 3, 5, 7, 11, 13}

	s1 = s1[1:4]
	fmt.Println(s1)

	s1 = s1[:2]
	fmt.Println(s1)

	s1 = s1[1:]
	fmt.Println(s1)
}

func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13} //baut erst das array dann den slice
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceLengthCapacity() {
	fmt.Println("Slice length and capacity")
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	//Slice to zero length.
	s = s[:0]
	printSlice(s)

	//Extend lengtj
	s = s[:4]
	printSlice(s)

	//Drop first two values
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func makeSlice() {
	fmt.Println("make slice")
	a := make([]int, 5) //len(a) = 5
	printSliceS("a", a)

	b := make([]int, 0, 5) //len(b) = 0 cap(b) = 5
	printSliceS("b", b)

	c := b[:2]
	printSliceS("c", c)

	d := c[2:5]
	printSliceS("d", d)
}

func printSliceS(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

//append will resize the backing array as needed
func appendToSlice() {
	fmt.Println("Append to Slice")
	var s []int
	printSlice(s)

	//append works on nil slices
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	//add more than one element
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func rangeLoop() {
	fmt.Println("Range loop")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//i : index
	//v: copy of value
	//one can skip any of these with _
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
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
	sliceLiterals()
	sliceLengthCapacity()
	makeSlice()
	appendToSlice()
	rangeLoop()
}
