package main

import "fmt"

func main() {
	//long declaration

	var nama string
	fmt.Println("namanya", nama)

	var country string = "singapore"
	fmt.Println("negaranya", country)

	var age uint = 10
	fmt.Println(age)

	var bilangan1 int8 = 127
	fmt.Println(bilangan1)

	var bilangan2 float32 = 3.14
	fmt.Println(bilangan2)

	var cek bool
	fmt.Println("cek", cek)

	// var bil1, bil2 int
	var bil1, bil2 int = 10, 20
	fmt.Println(bil1)
	fmt.Println(bil2)

	// short declaration
	angka1 := 5
	fmt.Println(angka1)
	nama1 := "alta"
	fmt.Println(nama1)

	var berat = 10
	fmt.Println("berat a: ", berat) //10
	//proses memasukkan/memberikan value/nilai 20 e variabel berat
	//assign ulang value
	berat = 20
	berat = 100
	fmt.Println("berat b: ", berat) //20

	//constanta
	const phi float64 = 3.14
	fmt.Println("phi", phi)
	// phi = 100
	fmt.Println(phi)
}
