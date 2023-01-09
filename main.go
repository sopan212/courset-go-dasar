package main

import "fmt"

func main() {
	var address string = "anaskdn129u1uu94rup9u94u90u290"
	var boolean bool = true
	var umur uint8 = 19
	var suhu int8 = -10

	fmt.Println("alamat saya adalah", address)
	fmt.Println("umur saya adalah", umur)
	fmt.Println("saya sudah menikah", boolean)
	fmt.Println("suhu saya adalah", suhu)
	fmt.Println("Hello World")
	var nama string
	nama = "sopan mukti"

	var pertama, kedua, ketiga string

	pertama = "satu"
	kedua = "dua"
	ketiga = "tiga"

	fmt.Printf("nama saya adalah %s \n", nama)
	fmt.Println("urutan ke", pertama)
	fmt.Println("urutan ke", kedua)
	fmt.Println("urutan ke", ketiga)

	var angkaPertama uint8 = 20
	var angkaKedua uint8 = 50

	var jumlah uint8 = angkaPertama + angkaKedua
	var kurang uint8 = angkaPertama - angkaKedua
	var kali uint8 = angkaKedua * angkaPertama
	var bagi uint8 = angkaPertama / angkaKedua

	fmt.Println("hasil penambahan ", jumlah)
	fmt.Println("hasil pengurangan", kurang)
	fmt.Println("hasil perkalianm ", kali)
	fmt.Println("hasil pembagian ", bagi)

}
