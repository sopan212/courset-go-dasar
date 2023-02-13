package main

import (
	"fmt"
	"strconv"
)

type Buah struct {
	jenis string
	rasa  string
	harga uint16
}

func (b Buah) beliBuah(namaOrang string) string {

	hargaString := strconv.FormatUint(uint64(b.harga), 10)

	pembeli := namaOrang + " membeli" + b.jenis + " dengan rasa" + b.rasa + " dengan harga " + hargaString
	return pembeli
}
func (b Buah) editHargaBuah(hargaBaru uint16) Buah {
	b.harga = hargaBaru

	return b
}
func main() {
	apelMalang := Buah{
		jenis: "apel",
		rasa:  "manis",
		harga: 2000,
	}
	apelMalang = apelMalang.editHargaBuah(25000)
	fmt.Println(apelMalang.beliBuah("sopan"))
}
