package main

import "fmt"

func main() {
	var myAge = 30

	if myAge == 5 {
		fmt.Println("You are too young")
	}
	if myAge == 17 {
		fmt.Println("So Sweet")
	}
	if myAge >= 17 && myAge < 30 {
		fmt.Println("My Age is between 17 and 30")
	}
	if dadAge := 9; dadAge < myAge {
		fmt.Println(dadAge)
	}

}
