package main

import "fmt"

func main() {
	var namaUser string = "john"
	fmt.Println("value namauser:", namaUser)
	// digunakan untuk mendapatkan alamat memory dari suatu variabel
	fmt.Println("mem address namauser:", &namaUser)

	// declare pointer
	var pointNama *string = &namaUser
	fmt.Println("value pointnama:", pointNama) //0x0001
	// membaca value yg tersimpan di alamat memory
	fmt.Println("value dari alamat yg tersimpan di pointnama:", *pointNama)

	var temp string = namaUser
	fmt.Println("temp:", temp) // john
	temp = "doe"
	fmt.Println("temp:", temp)               // doe
	fmt.Println("value namauser:", namaUser) //john

	// assign value
	*pointNama = "fulan"
	fmt.Println("value dari alamat yg tersimpan di pointnama:", *pointNama) // fulan
	fmt.Println("value namauser:", namaUser)                                //fulan

}
