package main

import "fmt"

func linierSearch(elements []int, x int) int {
	var langkah = 0
	for i := 0; i < len(elements); i++ {
		langkah++
		if elements[i] == x {
			fmt.Println("langkah:", langkah)
			return i
		}
	}
	return -1
}

func binarySearch(elements []int, x int) int {
	var kiri = 0
	var kanan = len(elements) - 1
	var langkah = 0

	for kiri <= kanan {
		langkah++
		// mencari index tengah
		var tengah = (kiri + kanan) / 2
		fmt.Println("tengah:", tengah)
		// membandingkan nilai yg dicari dengan value di index tengah
		if x < elements[tengah] {
			// jika lebih kecil, maka geser kanan ke kiri
			kanan = tengah - 1
		} else if x > elements[tengah] {
			// jika nilai yg dicari lebih besar, maka geser kiri ke kanan
			kiri = tengah + 1
		} else if x == elements[tengah] {
			fmt.Println("langkah:", langkah)
			// jika value di index tengah == nilai yg dicari, maka return index tengah
			return tengah
		}
	}
	return -1
}

func main() {
	var data = []int{12, 15, 15, 19, 24, 31, 53, 59, 60}
	// index := linierSearch(data, 12)
	index := binarySearch(data, 59)
	fmt.Println("index:", index)
}
