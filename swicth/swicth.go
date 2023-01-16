package main

import "fmt"

func main() {

	fmt.Println("swicth case")
	warna := "Jingga"

	switch warna {
	case "biru":
		fmt.Println("Case Pertama")
	case "merah":
		fmt.Println("Case Kedua")
	case "Jingga":
		fmt.Println("Case Ketiga")
	default:
		fmt.Println("Case Default")
	}

	fmt.Println("switch with falltrough")

	auth := "admin"

	switch auth {
	case "admin":
		fmt.Println("halaman admin")
		fallthrough
	case "guest":
		fmt.Println("halaman tamu")
	default:
		fmt.Println("halaman register")
	}
}
