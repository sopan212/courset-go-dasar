package main

import "fmt"

func main() {
	fmt.Println("materi MAP")
	var barang map[string]string
	barang = map[string]string{}
	barang["pertama"] = "sapu"
	barang["kedua"] = "pel"

	fmt.Println(barang["pertama"])

	fmt.Println(barang["kedua"])

	code := map[int]string{
		1: "piring",
		2: "gelas",
		3: "sendok",
	}
	fmt.Println(code[1])

	fmt.Println(code[2])

	nilaiSiswa := map[string]uint8{
		"sopan": 80,
		"mukti": 89,
		"fani":  23,
		"bagus": 90,
	}
	fmt.Println(nilaiSiswa["sopan"])
	fmt.Println(nilaiSiswa["mukti"])
	fmt.Println(nilaiSiswa["fani"])

	for key, val := range nilaiSiswa {
		fmt.Println(key, "nilainy adalah = ", val)
	}

	var nilaiBagus, ifExist = nilaiSiswa["bagus"]
	if ifExist {
		fmt.Println(nilaiBagus)
	} else {
		fmt.Println("nilai tidak ada")
	}

	fmt.Println(len(nilaiSiswa))
	delete(nilaiSiswa, "sopan")
	fmt.Println(nilaiSiswa["sopan"])

	//slice of map

	var alamat = []map[string]string{
		map[string]string{
			"nama":      "sopan",
			"alamat":    "jagakarsa",
			"pekerjaan": "IT",
		},
		map[string]string{
			"nama":      "mukti",
			"alamat":    "jaksel",
			"pekerjaan": "staff",
		},
		map[string]string{
			"nama":      "dimas",
			"alamat":    "jaktim",
			"pekerjaan": "staff",
			"hobi":      "naikgunung",
		},
	}
	for _, val := range alamat {
		fmt.Println("nama :", val["nama"])
		fmt.Println("alamat :", val["alamat"])
		fmt.Println("pekerjaan :", val["pekerjaan"])
		fmt.Println("hobi :", val["hobi"])
	}
	fmt.Println(alamat)

	//map of slice

	some := map[string][]string{
		"Fruits": []string{"jambu", "apel", "jeruk"},
		"Car":    []string{"bmw", "audi", "honda"},
	}
	for key, val := range some {
		fmt.Println(key, "value of :")
		for _, v := range val {
			fmt.Println("---", v)
		}
	}
}
