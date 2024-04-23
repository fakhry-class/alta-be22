package main

import "fmt"

func main() {
	var firstName string
	fmt.Println("Please enter your first name: ")
	fmt.Scanln(&firstName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s!\n", firstName)

	var angka3 int
	fmt.Println("masukkan angka:")
	fmt.Scanln(&angka3)
	fmt.Println("angkanya:", angka3)

}
