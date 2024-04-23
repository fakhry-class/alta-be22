package main

import "fmt"

func selectionSort(elements []int) []int {
	n := len(elements)
	for k := 0; k < n; k++ {
		minimal := k
		for j := k + 1; j < n; j++ {
			if elements[j] < elements[minimal] {
				minimal = j
			}
		}
		// swap
		elements[k], elements[minimal] = elements[minimal], elements[k]
		/*
			element k = gelas A
			elemen minimal = gelas B
			temp = gelas C
		*/
		// var temp = elements[k]
		// elements[k] = elements[minimal]
		// elements[minimal] = temp
	}
	return elements
}

func main() {
	var data = []int{10, 50, 30, 40, 20, 60, 4, 7, 6, 8, 11, 2, 3, 5, 78, 9, 45, 1}
	fmt.Println("sebelum", data)
	result := selectionSort(data)
	fmt.Println("sesudah", result)
}
