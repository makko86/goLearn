package main

import (
	"fmt"
	"math"
)

//functions may be used as function arguments
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//fucntion closures: function "bound" to variable
//"stores" its own value of bound variable
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	fmt.Println()
	//adder() "bound" to pos and neq
	pos, neq := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neq(-2*i))
	}
}
