package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var numb1 uint16 = 20
	var numb2 uint8 = 30
	jumlhUint := numb1 + uint16(numb2)
	fmt.Println(jumlhUint)

	var num3 int8 = 24
	jumlInt := num3 + int8(numb1)
	fmt.Println(jumlInt)

	//int to string
	var Snumb1 = strconv.Itoa(int(numb1))
	fmt.Println(Snumb1)
	fmt.Println(reflect.ValueOf(Snumb1).Type())

	//string to int

	var inToStr = "123"

	str1, _ := strconv.Atoi(inToStr)
	fmt.Println(reflect.ValueOf(str1))
	fmt.Println(reflect.ValueOf(str1).Type())
	// fmt.Println(reflect.ValueOf(err))
	// fmt.Println(reflect.ValueOf(err).Type())

	//string to float32

	str2 := "10,2"

	fl, _ := strconv.ParseFloat(str2, 32)
	fmt.Println(reflect.ValueOf(fl))
	fmt.Println(reflect.ValueOf(fl).Type())

}
