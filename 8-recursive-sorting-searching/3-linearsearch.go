package main

import "fmt"

/*
data = 10, 50, 30, 40, 20,60,4,7,6,8,11,2,3,5,78,9,45,1
x = 11
*/
func linierSearch(elements []int, x int) int {
	// proses membaca data 1 per 1
	for i := 0; i < len(elements); i++ {
		// membandingkan data yg lgi dibaca apakah sama dengan data yg dicari
		fmt.Println(elements[i])
		if elements[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var data = []int{10, 50, 30, 40, 20, 60, 4, 7, 6, 8, 11, 2, 3, 5, 78, 9, 45, 1}
	var cari = 11
	index := linierSearch(data, cari)
	fmt.Println(index)
	if index < 0 {
		fmt.Println("data gak ada cuy")
	} else {
		fmt.Println("data ditemukan di ", index)
	}
}
