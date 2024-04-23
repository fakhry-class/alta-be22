package main

import "fmt"

func main() {
	hour := 10
	if hour < 12 {
		fmt.Println("Selamat Pagi")
	} else if hour < 18 {
		fmt.Println("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}

	if hour < 18 {
		fmt.Println("masih jam kerja")
	}

	// jam := "10.30 am"
	makanan := "tahu"
	if makanan == "tempe" || makanan == "tahu" {
		fmt.Println("ini protein nabati")
	} else if makanan == "ayam" {
		fmt.Println("ini protein hewani")
	} else {
		fmt.Println("ini karbo")
	}

}
