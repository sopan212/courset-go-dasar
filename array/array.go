package main

import "fmt"

func main() {
	fmt.Println("array ")

	var a [3]string
	a[0] = "buah"
	a[1] = "rumah"
	a[2] = "mobil"
	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])

	angka := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(angka[0])
	fmt.Println(angka[2])
	fmt.Println(angka[3])
	fmt.Println(angka[4])
	fmt.Println(angka[5])
	fmt.Println(angka[6])
	for i := 0; i < len(angka); i++ {
		fmt.Println("angka ke:", i, "nilainya", angka[i])
	}

	buah := [3]string{"jambu", "apel", "tomat"}
	for i, val := range buah {
		fmt.Println("index", i, "adalah", val)
	}

	//aray tampa declarasi panjang/index

	arr := [...]int{
		1,
		4,
		7,
		8,
	}
	fmt.Println(arr)

	//array multidimensi

	fmt.Println("array multidimensi")
	buah2 := [2][3]int{
		[3]int{1, 2, 4},
		[3]int{0, 3, 5},
	}
	fmt.Println(buah2)

	//versi simple array multidimensi

	nama := [2][3]string{{"dani", "joko", "rani"}, {"siha", "raja", "nangka"}}
	fmt.Println(nama)

}
