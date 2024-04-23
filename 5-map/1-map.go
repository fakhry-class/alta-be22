package main

import "fmt"

func main() {
	// long declaration
	var salary = map[string]int{}
	fmt.Println(salary)
	// assign value ke map
	salary["andi"] = 10000
	salary["budi"] = 20000
	salary["10"] = 30000
	fmt.Println(salary)
	fmt.Println(salary["budi"])

	var data1 = map[int]string{}
	data1[10] = "sepuluh"
	data1[5] = "lima"
	data1[10] = "satu"
	fmt.Println(data1)

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	// short declaration
	salary_b := map[string]int{}
	fmt.Println(salary_b)

	// using make
	var salary_c = make(map[string]int)
	salary_c["doe"] = 7000 // assign value
	fmt.Println(salary_c)

	var data2 = map[bool]int{}
	data2[true] = 10
	data2[false] = 20
	data2[true] = 100
	fmt.Println(data2)
	fmt.Println(len(data2))
}
