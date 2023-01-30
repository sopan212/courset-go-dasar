package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	var messages = make(chan int, 3)

	go func() {
		for {
			i := <-messages
			fmt.Println("recive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
