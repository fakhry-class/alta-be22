package main

import "fmt"

func main() {
	var today int = 5
	// if today == 1 {
	// 	fmt.Printf("Today is Monday")
	// } else if today == 2 {
	// 	fmt.Printf("Today is Tuesday")
	// } else {
	// 	fmt.Printf("Invalid Date")
	// }

	switch today {
	case 1:
		fmt.Printf("Today is Monday\n")
	case 2:
		fmt.Printf("Today is Tuesday\n")
	default:
		fmt.Printf("Invalid Date\n")
	}

	// --------------
	switch {
	case today == 1:
		fmt.Printf("Today is Monday\n")
	case today == 2:
		fmt.Printf("Today is Tuesday\n")
	default:
		fmt.Printf("Invalid Date\n")
	}

	// ---------------

	var nama string = "budi"
	switch nama {
	case "andi":
		fmt.Println("halo andi")
	case "budi":
		fmt.Println("halo budi")
	case "cindy":
		fmt.Println("halo cindy")
	default:
		fmt.Println("invalid")
	}
}
