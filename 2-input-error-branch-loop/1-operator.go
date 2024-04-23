package main

import "fmt"

func main() {
	var angka1 = 15
	var angka2 = 2

	// var hasil1 float64 = float64(angka1 / angka2)
	var hasil1 = float64(angka1) / float64(angka2)
	fmt.Printf("hasil: %.2f \n", hasil1)
}
