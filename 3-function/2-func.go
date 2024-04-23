package main

import "fmt"

/*
menghitung luas persegi
luas = sisi * sisi *10
*/
func main() {
	var sisiP1 int = 10
	// var luasP1 = sisiP1 * sisiP1 * 10
	var luasP1 = hitungLuasPersegi(sisiP1) // sisi = 10
	fmt.Println(luasP1)

	var sisiP2 int = 5
	// var luasP2 = sisiP2 * sisiP2 * 10
	var luasP2 = hitungLuasPersegi(sisiP2) // sisi = 5
	fmt.Println(luasP2)

	var sisiP3 int = 8
	var luasP3 = hitungLuasPersegi(sisiP3)
	if luasP3 > 10 {
		fmt.Println("persegi nya luas banget")
	}
	fmt.Println(luasP3)
}

func hitungLuasPersegi(sisi int) int {

	var luas = sisi * sisi
	return luas
}
