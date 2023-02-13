package main

import (
	"fmt"
	"runtime"
	"time"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}
func print1(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func print2(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func print3(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func print4(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

	go print(5, "hallo")
	go print1(5, "haii")
	go print2(5, "hei")
	go print3(5, "ha")
	go print4(5, "hehehhe")

	print(5, "apa kabar")

	time.Sleep(2 * time.Second)

}
