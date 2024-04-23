package main

import (
	"fmt"
	"reflect"
)

func main() {
	var countries [2]string
	countries[0] = "Indonesia"
	countries[1] = "malaysia"
	countries[1] = "singapura"
	fmt.Println(countries)
	fmt.Println(countries[1])

	odd_numbers := [5]int{1, 3, 5, 7, 9}      // Intialized with values
	var even_numbers [5]int = [5]int{1, 2, 4} // Partial assignment

	fmt.Println(odd_numbers)
	fmt.Println(even_numbers)

	primes := [...]int{2, 3, 3, 6}

	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(len(primes))

	evennum := [5]int{1: 2, 2: 4, 3: 10}
	fmt.Println(evennum)
}
