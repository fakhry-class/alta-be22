package main

import "fmt"

func main() {

	// long declaration with value
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a, len(salary_a))

	salary_a["nabilah"] = 2000 // assign value
	fmt.Println(salary_a)

	delete(salary_a, "iswanul") // remove value by key
	fmt.Println(salary_a)

	fmt.Println(salary_a["budi"])
	fmt.Println(salary_a["nabilah"])
	value, isExist := salary_a["budi"] // check key is exist
	fmt.Println("isinya:", value, "isExist:", isExist)

	for key, value := range salary_a { // loop over maps
		fmt.Println("->", key, value)
		if value == 2000 {
			fmt.Println("ketemu 2000:", key)
		}
	}
}
