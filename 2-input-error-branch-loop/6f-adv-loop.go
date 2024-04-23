package main

import "fmt"

func main() {
	N := 5
	for baris := 0; baris < N; baris++ {
		for kolom := 0; kolom <= baris; kolom++ {
			fmt.Print("*-")
		}
		fmt.Println()
	}

	// var base int = 10
	// var pangkat int = 2
	// var hasil int
	// //code
	// for i := 0; i < count; i++ {

	// }
	// fmt.Println(hasil)
}
