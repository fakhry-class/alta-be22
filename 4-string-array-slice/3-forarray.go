package main

import "fmt"

func main() {
	// var ujian1 int = 100
	// var ujian2 int = 80
	// var ujian3 int = 90
	// var ujian4 int = 60
	// var ujian5 int = 95

	// var rata2 = (ujian1 + ujian2 + ujian3 + ujian4 + ujian5) / 5

	var ujianFakhry [5]int = [5]int{100, 80, 90, 60, 95}
	// fmt.Println(ujianFakhry[0])
	// fmt.Println(ujianFakhry[1])
	// fmt.Println(ujianFakhry[2])
	// fmt.Println(ujianFakhry[3])
	// fmt.Println(ujianFakhry[4])
	var jumlah int
	// for i := 0; i < len(ujianFakhry); i++ {
	// 	jumlah = jumlah + ujianFakhry[i]
	// }
	for _, element := range ujianFakhry {
		jumlah = jumlah + element
	}
	fmt.Println("jumlah", jumlah)
	var rata2 = jumlah / len(ujianFakhry)
	fmt.Println("rata2", rata2)

	// technique 1
	for index := 0; index < len(ujianFakhry); index++ {
		fmt.Println(ujianFakhry[index])
	}
	// technique 2
	for index, element := range ujianFakhry {
		fmt.Println(index, "=>", element)
	}
	for _, value := range ujianFakhry {
		fmt.Println(value)
	}
	// technique 3
	index := 0
	for range ujianFakhry {
		fmt.Println(ujianFakhry[index])
		index++
	}

}
