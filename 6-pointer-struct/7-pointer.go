package main

import "fmt"

func fillCupValue(cup string) string {
	// fmt.Println("value awal", cup)
	fmt.Println("address", &cup)
	cup = "kopi"
	// fmt.Println("value akhir", cup)
	return cup
}

func fillCupReference(cup *string) string {
	fmt.Println("address", cup)
	*cup = "kopi"
	return *cup
}

func main() {
	var cup string = "kosong"
	fmt.Println("addrs", &cup)
	// result := fillCupValue(cup)
	result := fillCupReference(&cup)

	fmt.Println("value result", result)
	fmt.Println("value cup", cup)
}
