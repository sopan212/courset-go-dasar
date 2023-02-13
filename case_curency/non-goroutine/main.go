package main

import "fmt"

// restoran -> checking belanjaan (ikan, daging, sayur, buah, bawang)
//
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur -> kupas -> simpan di box -> simpan chiler
// buah -> cuci -> simpan chiler
// bawang -> kupas -> simpan di box -> simpan chiler
// jika bukan barang di atas -> buang sampah

func main() {

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik",
		"bambu", "buah", "sterofoam", "bawang", "ikan",
		"sayur", "buah", "buah", "buah", "daging", "gula",
		"ikan", "plastik", "bawang", "daging",
	}

	for _, barang := range belanja {
		switch barang {
		case "ikan":
			fmt.Println("menemukan item", barang)
			cuci(barang)

		case "sayur":
			fmt.Println("menemukan item", barang)
			kupas(barang)

		case "daging":
			fmt.Println("menemukan item", barang)
			potong(barang)

		case "buah":
			fmt.Println("menemukan item", barang)
			cuci(barang)

		case "bawang":
			fmt.Println("menemukan item", barang)
			kupas(barang)

		default:
			fmt.Println("menemukan item", barang)
			buangSampah(barang)
		}

	}

}

func buangSampah(item string) {
	fmt.Println("membuang", item, "kesampah")
}
func potong(item string) {
	fmt.Println("memotong", item)
	simpanDiFrozen(item)
}

func kupas(item string) {
	fmt.Println("mengupas", item)
	simpanDiBox(item)
	simpanDiChiler(item)
}

func simpanDiBox(item string) {
	fmt.Println("menyimpan", item, "di box")
}

func cuci(item string) {
	fmt.Println("mencuci", item)
	if item == "buah" {
		simpanDiChiler(item)
	} else {
		simpanDiFrozen(item)
	}
}

func simpanDiChiler(item string) {
	fmt.Println("menyimpan", item, "di chiler")

}

func simpanDiFrozen(item string) {
	fmt.Println("menyimpan", item, "di Frozen")
}
