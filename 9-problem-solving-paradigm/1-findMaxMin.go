package main

import "fmt"

func findMaxMin(source []int) (max int, min int) {
	// asumsi awal,
	max = source[0]
	min = source[0]

	// digunakan untuk membaca 1 per 1 data
	for i := 0; i < len(source); i++ {
		// mencari nilai yang lebih besar dari max
		// membandingkan data yg dibaca dg value max
		if source[i] > max {
			// jika lebih besar, maka nilai max diganti
			max = source[i]
		}

		// mencari nilaiyg lebih kecil dari min
		if source[i] < min {
			// jika ada data yg lebih kecil dari min, maka nilai min diganti
			min = source[i]
		}
	}
	return max, min
}

func main() {
	var data = []int{10, 7, 3, 5, 8, 2, 9}
	hasil1, hasil2 := findMaxMin(data)
	fmt.Println("max:", hasil1, "min:", hasil2)
}
