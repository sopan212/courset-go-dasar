package main

import "fmt"

func main() {
	fmt.Println("for loop")
	// i := 0

	// for i < 5 {
	// 	fmt.Println("Perulangan i ke : ", i)
	// 	i++
	// }

	// for {
	// 	fmt.Println("perulangan i ke :", i)
	// 	i++
	// 	if i == 5 {
	// 		fmt.Println("perulangan di hentikan di break ke:", i)
	// 		break
	// 	}
	// }
	for i := 0; i <= 10; i++ {
		fmt.Println("Perulangan i ke: ", i)
	}

	for j := 0; j <= 5; j++ {

		if j == 1 {
			fmt.Println("value j ke :", j)
			// continue
		}
		fmt.Println("nilai j ke :", j)
	}
}
