package main

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, cha chan float64) {
	var sum = 0
	for _, e := range numbers {
		fmt.Println("looop e ", e)
		sum = sum + e
	}
	cha <- float64(sum) / float64(len(numbers))
	fmt.Println("isi chanel", cha)
}

func getMax(numbers []int, cha chan int) {
	var max = numbers[0]
	for _, e := range numbers {
		if max < e {
			max = e
		}
	}
	cha <- max
}

func main() {
	runtime.GOMAXPROCS(2)
	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 6, 3}
	fmt.Println("numbers", numbers)
	var ch1 = make(chan float64)

	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("avg \t: %.2f \n", avg)

		case max := <-ch2:
			fmt.Printf("max \t: %d \n", max)

		}
	}

}
