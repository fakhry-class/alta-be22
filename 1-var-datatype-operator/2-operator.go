package main

import "fmt"

func main() {
	bilA := 10
	var bilB int = 20
	var hasil1 = bilA + bilB
	fmt.Println("hasil tambah", hasil1) //30

	hasil2 := bilA - bilB
	fmt.Println("hasil kurang", hasil2) //-10

	num1 := 9
	num2 := 2
	hasil3 := num1 % num2
	hasil4 := num1 / num2
	fmt.Println("hasil modulo", hasil3)
	fmt.Println("hasil bagi", hasil4)

	helloworld := "hello" + "---" + "world"
	fmt.Println(helloworld)
	num3 := "10"
	fmt.Println(num3)
}
