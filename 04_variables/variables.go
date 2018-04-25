package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//Integer ist 64 Bit groß auf 64 Bit Systemem
//Zero Values:
//0 for numeric types,
//false for the boolean type, and
//"" (the empty string) for strings.
var i, j int = 1, 2

//Variablen Deklaration "en Block"
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1 //64 Bit Integer; links shift für max
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//wenn Variablensofort initialisiert werden, muss kein Typ angegeben werden
	var c, python, java = true, false, "no!"
	//in einer Funktion auch ohne var
	k := 3
	fmt.Println(i, j, k, c, python, java)

	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)

	//Cast => immer explicit
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	const Pi = 3.14
	fmt.Println("Happy", Pi, "Day")
}
