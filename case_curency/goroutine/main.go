package main

import (
	"fmt"
	"runtime"
)

// restoran -> checking belanjaan (ikan, daging, sayur, buah, bawang)
//
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur -> kupas -> simpan di box -> simpan chiler
// buah -> cuci -> simpan chiler
// bawang -> kupas -> simpan di box -> simpan chiler
// jika bukan barang di atas -> buang sampah
var cuciChanel = make(chan string)
var potongChanel = make(chan string)
var kupasChanel = make(chan string)
var boxChanel = make(chan string)
var frozenChanel = make(chan string)
var chilerChanel = make(chan string)
var trashChanel = make(chan string)

func main() {
	runtime.GOMAXPROCS(2)

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik",
		"bambu", "buah", "sterofoam", "bawang", "ikan",
		"sayur", "buah", "buah", "buah", "daging", "gula",
		"ikan", "plastik", "bawang", "daging",
	}

	go checkBarang(belanja)
	go cuci()
	go potong()
	go kupas()
	go simpanBox()
	go simpanDiChiler()
	go simpanDiFrozen()
	go trash()

	var input string
	fmt.Scanln(&input)

	buah := []string{"nanas", "apel", "jeruk", "semangka"}
	belanjaan := append(belanjaan, buah...)
	fmt.Println("test", belanjaan)
}

func checkBarang(belanja []string) {
	for _, barang := range belanja {
		switch barang {
		case "ikan":
			fmt.Println("menemukan item", barang)
			cuciChanel <- barang

		case "sayur":
			fmt.Println("menemukan item", barang)
			kupasChanel <- barang

		case "daging":
			fmt.Println("menemukan item", barang)
			potongChanel <- barang

		case "buah":
			fmt.Println("menemukan item", barang)
			cuciChanel <- barang

		case "bawang":
			fmt.Println("menemukan item", barang)
			kupasChanel <- barang

		default:
			fmt.Println("menemukan item", barang)
			trashChanel <- barang
		}
	}
}

func cuci() {
	for cuciItem := range cuciChanel {
		fmt.Println("mencuci", cuciItem)
		if cuciItem == "buah" {
			chilerChanel <- cuciItem
		} else {
			frozenChanel <- cuciItem
		}
	}
}
func potong() {
	for itemPotong := range potongChanel {
		fmt.Println("memotong", itemPotong)
		frozenChanel <- itemPotong
	}
}
func kupas() {
	for kupasItem := range kupasChanel {
		fmt.Println("mengupas", kupasItem)
		boxChanel <- kupasItem
	}
}
func simpanBox() {
	for itemBox := range boxChanel {
		fmt.Println("menyimpan", itemBox)
		chilerChanel <- itemBox
	}
}

func simpanDiChiler() {
	for itemChiler := range chilerChanel {
		fmt.Println("menyimpan", itemChiler, "dichiler")
	}

}

func simpanDiFrozen() {
	for itemFrozen := range frozenChanel {
		fmt.Println("menyimpan", itemFrozen, "di frozen")
	}
}
func trash() {
	for itemTrash := range trashChanel {
		fmt.Println("membuang", itemTrash)
	}
}
