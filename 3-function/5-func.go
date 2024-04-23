package main

import "fmt"

func main() {
	var angka = 10
	fakhtor := faktorBilangan(angka)
	fmt.Println(fakhtor)
	// var jmlFaktor int
	// jmlFaktor = faktorBilangan(angka)
	// fmt.Println("jml faktor", jmlFaktor)
	//
	//	if jmlFaktor <= 4 {
	//		fmt.Println("bilangan keren")
	//	} else {
	//
	//		fmt.Println("bukan keren")
	//	}
}

func faktorBilangan(bil int) []int {
	// bil = bilangan yg ingin dicari jumlah faktor
	// return ini merepresentasikan jumlah faktor dari bil
	/*
		bil = 10
		faktor = 1,2,5,10 = 4
		return 4
		bil = 20
		return 6(1,2,4,5,10,20)
	*/
	var hasil []int
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			hasil = append(hasil, i)
		}
	}
	fmt.Println("jml", hasil)
	return hasil
}
