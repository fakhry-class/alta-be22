package main

import "fmt"

func main() {
	var v1 int = 300
	var v2 int = 700

	if v1 == 400 {
		if v2 == 700 {
			fmt.Printf("Value of v1 is 400 and v2 is 700\n")
		} else {
			fmt.Println("v1 = 400 tapi v2 nya != 700")
		}
	} else {
		fmt.Println("ini else v1")
	}
}
