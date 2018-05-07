package main

import "fmt"

//access concrete value; test type
func typeAssertion() {
	var i interface{} = "hallo"

	s := i.(string) //actual value of i is of type string -> all good
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //actual value of i IS NOT of type float64 -> f = 0 and ok = false
	fmt.Println(f, ok)

	//f = i.(float64) //panics
	//fmt.Println(f)
}

//Several type assertions in a row
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an int!", v*2)
	case string:
		fmt.Println("I'm a string!", v)
	default:
		fmt.Println("I don't know who I am :(", v)
	}
}

//Stringer interface
//A Strigner is a type that can describe itself as a string
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	typeAssertion()
	typeSwitch(21)
	typeSwitch("text")
	typeSwitch(true)

	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
