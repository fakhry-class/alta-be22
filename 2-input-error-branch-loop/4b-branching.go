package main

import "fmt"

func main() {
	number := 4
	if number%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("Ganjil")
	}

	var myAge = 30
	if myAge >= 17 && myAge < 30 {
		fmt.Println("My Age is between 17 and 30")
	} else {
		fmt.Println("My Age is lower than 17 and higher than 29")
	}

	var myAge2 = 30
	//implementasi OR
	if myAge2 >= 17 || myAge2 < 30 {
		fmt.Println("dewasa")
	} else {
		fmt.Println("tidak masuk kelompok umur")
	}

}
