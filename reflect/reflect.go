package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("reflect")

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("reflect value type", reflectValue.Type())
	fmt.Println("reflect value", reflectValue)

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable", reflectValue.Int())
	}
}
