package main

import (
	"fmt"
	"math/rand"
	"time"
)

// goroutine is a function that runs concurrently with other functions
func ping(x int) {
	for i := 0; i < 100; i++ {
		fmt.Println(x, "ping", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(250)))
	}
}

func main() {
	for i := 0; i < 10; i++ {
		// this runs concurrently
		go ping(i)
	}
	var input string
	fmt.Scanln(&input)
}
