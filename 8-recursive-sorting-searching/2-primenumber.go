package main

import (
	"fmt"
	"math"
)

func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i <= sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func checkPrime1(number int) bool {
	if number < 2 {
		return false
	}
	var counter int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			counter++
		}
	}
	if counter > 2 {
		return false
	}
	return true
}

func main() {
	// 	fmt.Println(checkPrime1(1))
	// 	fmt.Println(checkPrime1(11))
	// 	fmt.Println(checkPrime1(15))
	fmt.Println(checkPrime(200000000001))
}
