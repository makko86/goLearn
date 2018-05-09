package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walk(t *tree.Tree, ch chan int) {
	//left
	if t.Left != nil {
		walk(t.Left, ch)
	}
	//self
	ch <- t.Value

	//right
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

func same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go walk(t1, ch1)
	go walk(t2, ch2)
	//only 10 values in t1 and t2...
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	//Test walk...
	ch := make(chan int)
	go walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println()

	fmt.Println(same(tree.New(1), tree.New(2)))
	fmt.Println(same(tree.New(1), tree.New(1)))
	fmt.Println(same(tree.New(2), tree.New(1)))
	fmt.Println(same(tree.New(2), tree.New(2)))
}
