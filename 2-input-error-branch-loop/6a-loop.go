package main

import "fmt"

func main() {
	// fmt.Println("halo alta")
	// fmt.Println("halo alta")
	// fmt.Println("halo alta")
	// fmt.Println("halo alta")
	// fmt.Println("halo alta")
	for i := 0; i < 5; i++ {
		// code
		fmt.Println("halo alta", i)
		if i == 2 {
			fmt.Println("horray")
			//for
		}
	}

	/*
		i++ --> i=i+1
	*/

	for j := 1; j <= 10; j = j + 2 {
		// code
		fmt.Println("nilai j:", j)
	}
}
