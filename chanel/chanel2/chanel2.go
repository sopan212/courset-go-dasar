package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloFor = func(name string) {
		var item = fmt.Sprintln("hallo", name)

		messages <- item
	}

	go sayHelloFor("mukti")
	var res = <-messages
	fmt.Println(res)

	go sayHelloFor("sopan")
	var res1 = <-messages
	fmt.Println(res1)

	go sayHelloFor("sopan")
	var res2 = <-messages
	fmt.Println(res2)
}
