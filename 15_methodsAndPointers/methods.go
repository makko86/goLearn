package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//define methods on types
//(v Vertex) = receiver argument
//could be written as regular function
//cannot modifiy values (Call-by-Value)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//can be defined on non-struct types
//but only on types declared in same package
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//pointer receivers
//can modify values
//no neede to be called with "&"
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Pointer arguments
//must be called with "&"
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//possible to Call-by-value on pointers
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//methods should either have value OR pointer receivers not both!

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())

	ScaleFunc(&v, 2)
	fmt.Println(v)

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
