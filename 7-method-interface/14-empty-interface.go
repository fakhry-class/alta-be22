package main

import "fmt"

//	func describe(i interface{}) {
//		fmt.Printf("(%v, %T)\n", i, i)
//	}
func describe(i any) {
	fmt.Printf("(%v, %T)\n", i, i)
}
func cetak(data int) {
	fmt.Println("datanya", data)
}

func main() {
	// var data interface{}
	var data any
	describe(data)

	data = 42
	describe(data)

	data = "hello"
	describe(data)

	data = true
	describe(data)

	data = 3.14
	describe(data)

	data = []int{1, 2, 3}
	describe(data)

	// var hasil = "halo"
	// cetak(hasil)
}
