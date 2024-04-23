package main

import (
	"fmt"
)

func main() {
	sentence := "Hello alta"

	fmt.Println("length:", len(sentence))
	// i = 0,1,2,3,4
	// sentence[0]
	// sentence[1]
	// sentence[2]
	// sentence[3]
	// sentence[4]
	for i := 0; i < len(sentence); i++ {
		fmt.Println("karakter:", string(sentence[i]))
	}

	fmt.Println("-----")

	for position, value := range sentence {
		fmt.Printf("character %c starts at byte position %d\n", value, position)
	}
}
