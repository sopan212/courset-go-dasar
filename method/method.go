package main

import "fmt"

func tanya(a bool) (resp string) {
	if a {
		resp = "benar"
	} else {
		resp = "salah"
	}
	return resp
}

func sayHello(name string) string {
	text := "Hallo " + name
	return text
}

func kurangDanBagi(angka1, angka2 float32) (kurang, bagi float32) {
	kurang = angka1 - angka2
	bagi = angka1 / angka2
	return bagi, kurang
}
func hitung(p uint8, l uint8) (h uint8) {
	hasil := p * l
	return hasil
}
func main() {

	name := "sopan"

	runSayHello := sayHello(name)
	fmt.Println("result: ", runSayHello)

	var p uint8 = 12
	var l uint8 = 14

	runHitung := hitung(p, l)
	fmt.Println("hasil dari hitung", runHitung)

	var angka1 float32 = 23
	var angka2 float32 = 16

	runBagi, runKurang := kurangDanBagi(angka1, angka2)

	fmt.Println("hasil pembagian", runBagi)
	fmt.Println("hasil pengurangan", runKurang)
	var a bool = false
	fmt.Println("apakah benar?", tanya(a))

	//closure
	var nilai = func(n string) (nilai uint16) {
		nilaiSiswa := map[string]uint16{
			"sopan": 80,
			"mukti": 89,
			"fani":  23,
			"bagus": 90,
		}
		nilai = nilaiSiswa["sopan"]
		return nilai
	}
	fmt.Println("nilai sopan", nilai("sopan"))
}
