package main

import "fmt"

func main() {
	food := [...]string{"egg", "rice", "noodle", "beef", "chiken"}

	newfood := food[1:2]
	fmt.Println(newfood)

	newfood1 := food[:4]
	fmt.Println(newfood1)

	newfood2 := food[2:]
	fmt.Println(newfood2)

	newfood4 := food[3:4]
	fmt.Println(newfood4)
	fmt.Println("capacity", cap(newfood4))
	fmt.Println("length", len(newfood4))

	newfood4 = append(newfood4, "sambal", "petai")
	fmt.Println(newfood4)
	fmt.Println("capacity", cap(newfood4))
	fmt.Println("length", len(newfood4))

	fruit := []string{"guava", "apple", "penalpe"}

	newfood4 = append(newfood4, fruit...)
	fmt.Println(newfood4)
	fmt.Println("capacity", cap(newfood4))
	fmt.Println("length", len(newfood4))

}
