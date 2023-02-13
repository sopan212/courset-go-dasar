package main

import "fmt"

func main() {
	var no int16 = 4
	var x int8 = 4

	fmt.Println("value nomor", no)
	fmt.Println("pointer nomor", &no)

	fmt.Println("value x", x)
	fmt.Println("pointer x", &x)

	a := x
	fmt.Println("value a", a)
	fmt.Println("pointer a", &a)

	var y *int16 = &no

	fmt.Println("value y", y)
	fmt.Println("pointer y", &y)

}
