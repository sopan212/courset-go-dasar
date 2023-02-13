package main

import "fmt"

type Makanan struct {
	sayur Sayur
	Buah
	minuman string
}

type Sayur struct {
	jenis string
	warna string
	harga uint16
}
type Buah struct {
	jenis string
	rasa  string
	harga uint16
}

func main() {
	fmt.Println("struct")

	var manggaKueni Buah
	manggaKueni.jenis = "mangga kueni"
	manggaKueni.harga = 25000
	manggaKueni.rasa = "manis"

	fmt.Println("mangga kueni", manggaKueni)

	jerukBali := Buah{
		jenis: "jeruk bali",
		rasa:  "manis",
		harga: 25000,
	}
	fmt.Println("jeruk bali", jerukBali)

	var makanSiang Makanan
	makanSiang.sayur.jenis = "bayam"
	makanSiang.minuman = "es jeruk"
	makanSiang.Buah = Buah{
		jenis: "apel",
	}
	fmt.Println(makanSiang)

}
