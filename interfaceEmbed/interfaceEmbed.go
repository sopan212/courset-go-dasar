package main

import (
	"fmt"
	"math"
)

type Hitung2d interface {
	luas() float64
	keliling() float64
}

type Hitung3d interface {
	volume() float64
}

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Kubus struct {
	sisi float64
}

func (k *Kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *Kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *Kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {

	var bangunRuang Hitung = &Kubus{4}

	fmt.Println("Keliling kubus :", bangunRuang.keliling())
	fmt.Println("Luas kubus :", bangunRuang.luas())
	fmt.Println("Volume kubus :", bangunRuang.volume())
}
