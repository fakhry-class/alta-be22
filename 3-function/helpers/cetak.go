package helpers

import "fmt"

func CetakHello() {
	fmt.Println("hello from helpers")
	hitungLuasPersegi()
}

func printHello() {
	fmt.Println("hello from helpers.printHello")
}

// func with type return
func TampilHello() string {
	var hasil string = "hello from TampilHello"
	return hasil
}

func TampilBintang() string {
	fmt.Println("*")
	fmt.Println("*")
	fmt.Println("*")
	var hasil string
	hasil += "*\n"
	hasil += "*\n"
	hasil += "*\n"
	return hasil
}
