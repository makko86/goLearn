/*
* go f(x, y, z)
* f, x, y, z are evaluated in current goroutine; execution happens ins new goroutine
* access to shared memory must be synchronized
 */
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world!")
	say("Hello")
}
