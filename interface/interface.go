package main

import (
	"fmt"
	"math"
)

type Bilangan struct {
	a int8
	b int8
}

func (b Bilangan) tambah() int8 {
	return b.a + b.b
}

type Calculate interface {
	luas() float64
	keliling() float64
}

type Persegi struct {
	sisi float64
}

type Lingkaran struct {
	diameter float64
}

func (p Persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (l Lingkaran) luas() float64 {
	jariJari := l.diameter / 2
	return math.Pi * math.Pow(jariJari, 2)
}

func (p Persegi) keliling() float64 {
	return 4 * p.sisi
}
func (l Lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}
func main() {

	bilangan := Bilangan{
		a: 8,
		b: 7,
	}

	fmt.Println(bilangan.tambah())

	var p Calculate
	p = Persegi{30}
	fmt.Println("luas persegi", p.luas())
	fmt.Println("keliling persegi ", p.keliling())

	var l Calculate
	l = Lingkaran{15}
	fmt.Println("luas Lingkaran :", l.luas())
	fmt.Println("Keliling Lingkaran :", l.keliling())
}
