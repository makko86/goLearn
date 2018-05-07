//Interfaces
//set of method signatures
//value of type <interface> cann hold any value tha implements those methods
//no need to declare the implementation of an interface. just do it
package main

import (
	"fmt"
	"math"
)

//add "-er" to name of function
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 { //actual implementation
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 { //actual implementation
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	//MyFloat and *Vertx implemente Abser
	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())

	//a = v won't work because Vertex DOES NOT implement Abser
}
