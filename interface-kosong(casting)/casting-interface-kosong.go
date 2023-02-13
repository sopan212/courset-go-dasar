package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 15

	var subt = secret.(int) * 10

	fmt.Println(subt)
	var subt2 = secret.(int) * 4

	fmt.Println(subt2)

	secret = []string{"apel", "nanas", "mangga"}
	var joinStr = strings.Join(secret.([]string), ":")

	fmt.Println(joinStr)

	//casting variable interface ke pointer

	type Person struct {
		name string
		age  uint8
	}

	var secret2 interface{} = &Person{"mukti", 28}

	var name = secret2.(*Person).name
	var age = secret2.(*Person).age

	fmt.Println(name, age)

	//kombinasi map,string,interface

	var person []map[string]interface{}

	person = []map[string]interface{}{
		{"name": "sopan mukti", "age": 28},
		{"name": "sopan", "age": 23},
		{"name": "mukti", "age": 18},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"nama": "strobery", "total": 20},
		[]string{"manggo", "pinaple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
