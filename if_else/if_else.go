package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var bolean bool = true
	if bolean {
		fmt.Println("kondisi benar")
	} else {
		fmt.Println("kondisi salah")
	}

	//operator pembanding
	var umur uint16 = 32

	if umur > 32 {
		fmt.Println("kondisi benar")
	} else {
		fmt.Println("kondisi salah")
	}

	//else if
	fmt.Println("else if")
	var condition uint16 = 13000
	if condition >= 13000 {
		fmt.Println("kondisi benar pertama")
	} else if condition >= 10000 {
		fmt.Println("kondisi benar kedua")
	} else if condition >= 5000 {
		fmt.Println("kondisi benar ketiga")
	} else {
		fmt.Println("kondisi salah")
	}

	fmt.Print("else if short hand \n")

	if yen := condition / 1000; yen >= 15 {
		fmt.Println("uang anda lebih dari 15 yen", yen, "YEN")
	} else if yen > 10 {
		fmt.Println("uang anda lebih dari 10 yen", yen, "YEN")
	} else if yen > 5 {
		fmt.Println("uang anda lebih dari 5 yen", yen, "YEN")
	} else {
		fmt.Println("uang anda", yen, "YEN")
	}
}
