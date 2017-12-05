package main

import (
	"fmt"
)

// can only send to chanel
func ping(c chan<- string) {
	for {
		var str string
		fmt.Scanln(&str)
		c <- str
	}
}

// can only receive from channel
func printer(c <-chan string) {
	for {
		msg := <- c
		fmt.Println("received", msg)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go printer(c)

	for {}
}
