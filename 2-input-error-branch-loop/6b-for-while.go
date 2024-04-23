package main

import "fmt"

func main() {
	// i := 0
	// for i < 5 {
	// 	// code
	// 	fmt.Println("halo alta", i)
	// 	i++
	// }

	sum := 1
	for sum < 10 {
		sum += sum
		// sum = sum + sum
		fmt.Println("nilai sum", sum)
	}
	fmt.Println("hasil akhir", sum)
}
