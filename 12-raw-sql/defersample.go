package main

import "fmt"

func main() {
	function1()
	defer fmt.Println("hello 1")
	defer fmt.Println("hello 2")
	fmt.Println("hello 3")
	fmt.Println("hello 4")
}

func function1() {
	fmt.Println("func 1")
	defer fmt.Println("func 2")
	fmt.Println("func 3")
}
