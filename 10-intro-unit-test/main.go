package main

import (
	"fakhry/calculator/calculate"
	"fakhry/calculator/hitung"
	"fmt"
)

func main() {
	resultKali := calculate.Perkalian(-5, 10)
	fmt.Println(resultKali)

	resultBonus := hitung.HitungBonusTHR(10000000, "senior")
	fmt.Printf("%0.2f\n", resultBonus)
}
