package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func main() {
	fmt.Println("chanel as parameter")

	runtime.GOMAXPROCS(2)

	var message = make(chan string)
	for _, each := range []string{"sopan", "mukti", "areht"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			message <- data
		}(each)
	}
	for i := 0; i < 3; i++ {
		printMessage(message)
	}

}
