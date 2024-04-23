package main

import (
	"fmt"
	"sort"
)

/*
pecahan uang= 10,25,5,1 --> di sort dari besar ke kecil 25,10,5,1
uang = 42

42 = 1,1,1,1,1,1,1,1 .... (42x) = 42 keping
42 = 5,5,5,5,5,5,5,5,1,1  = 10 keping
42 = 10,10,10,10,1,1 = 6 keping
42 = 10,5,5,5,5,5,5,1,1 = 9 keping
42 = 25,10,5,1,1 =  5 keping

100 = 25,25,25,25 = 4 keping
56 = 25,25,5,1 = 4 keping
16 = 10,5,1 = 3 keping

*/

func coinChange(money int, coinValue []int) []int {
	// sort besar ke kecil / DESCENDING/ DESC
	sort.Slice(coinValue, func(i, j int) bool {
		return coinValue[i] > coinValue[j]
	})
	// fmt.Println(coinValue)

	// buat var penampung
	var hasil []int
	// digunakan untuk membaca per coin nya
	for _, coin := range coinValue {
		// selama uang nya masih lebih besar dari coin yg sedang dproses, maka kurang nilai uang dengan coin tsb
		// dan simpan value coin nya ke hasil
		for money >= coin {
			// simpan coin
			hasil = append(hasil, coin)
			// proses mengurangi uang.
			money = money - coin
		}
	}
	return hasil

}

func main() {
	var coinValue = []int{10, 25, 5, 1}
	var uang = 16
	fmt.Println(coinChange(uang, coinValue))
}
