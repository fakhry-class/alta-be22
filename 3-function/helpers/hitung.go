package helpers

import "fmt"

func hitungLuasPersegi() {
	fmt.Println("ini func hitung luas persegi")
}

func HitungPenambahan() int {
	var bil1 = 20
	var bil2 = 30
	hasil := bil1 + bil2
	// fmt.Println(hasil)
	return hasil
}
func HitungPenambahanWithParam(bil1 int, bil2 int) int {

	hasil := bil1 + bil2
	// fmt.Println(hasil)
	return hasil
}
