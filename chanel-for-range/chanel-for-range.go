package main

import (
	"fmt"
	"runtime"
)

func sendMessage(cha chan<- string) {
	for i := 0; i < 20; i++ {
		cha <- fmt.Sprintf("data %d", i)
	}
	close(cha)
}

func printMessage(cha <-chan string) {
	for message := range cha {
		fmt.Println(message)
	}
}

func main() {

	runtime.GOMAXPROCS(2)

	var messgae = make(chan string)

	go sendMessage(messgae)
	printMessage(messgae)

}
