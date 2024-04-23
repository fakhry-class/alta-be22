package main

import "fmt"

func main() {
	// N := 5
	// for baris := 0; baris < N; baris++ {
	// 	for kolom := 0; kolom < N; kolom++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	N := 5
	for baris := 0; baris < N; baris++ {
		for kolom := 0; kolom < N; kolom++ {
			// fmt.Print("*")
			fmt.Printf("b:%d,k:%d ", baris, kolom)
		}
		fmt.Println()
	}
}
