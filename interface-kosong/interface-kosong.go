package main

import "fmt"

func main() {
	var secret interface{}
	fmt.Println("interface kosong")
	secret = "ethan hant"

	fmt.Println(secret)

	secret = [3]string{"nanas", "jambu", "apel"}

	fmt.Println(secret)

	secret = 201

	fmt.Println(secret)

	fmt.Println("interface kosong")
	var barang map[string]interface{}

	barang = map[string]interface{}{
		"name":   "palu",
		"ukuran": 23,
		"warna":  [3]string{"merah", "biru", "kuning"},
	}

	fmt.Println(barang)

	for k, i := range barang {
		fmt.Println(k, "--->", i)
	}
	// var data map[string]any

	// data = map[string]any{
	// 	"nama": "sopan",
	// 	"umur": 28,
	// 	"hobi": [...]string{"beranang", "berlari", "naik gunung"},
	// }

	// for i, k := range data {
	// 	fmt.Println("data ", i, " nilainya --> ", k)
	// }
}
