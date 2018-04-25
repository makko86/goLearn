package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

//wenn gleicher Typ dann muss dieser nur beim letzten Parameter angegeben werden
func add2(x, y int) int {
	return x + y
}

//mehrere Ergebnisse als Rückgabe
func swap(x, y string) (string, string) {
	return y, x
}

//Benannte Rückgaben, "Naked Return"
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	//Funktionsaufruf
	fmt.Println("Let's add 47 and 11:", add(47, 11))
	fmt.Println("Let's add 47 and 11:", add2(47, 11))

	//mehrere Ergebnisse als Rückgabe
	a, b := swap("Hello", "World!")
	fmt.Println("Hello World!", a, b)

	//Benannte Rückgaben, Naked Return
	fmt.Print("So naked: ")
	fmt.Println(split(17))
}
