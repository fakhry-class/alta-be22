package main

import "fmt"

func main() {
	// loop1:
	for i := 0; i < 100; i++ {
		if i == 1 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i)

		// switch {
		// case i == 1:
		// 	continue
		// 	// continue loop1
		// case i > 3:
		// 	break
		// 	// break loop1
		// default:
		// 	fmt.Println(i)
		// }
	}
}
