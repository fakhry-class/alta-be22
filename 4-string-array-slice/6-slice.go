package main

import "fmt"

func main() {
	var colors = []string{"red", "green", "yellow"}
	fmt.Println(colors)
	fmt.Println("len", len(colors))
	colors = append(colors, "purple")
	colors = append(colors, "blue")
	fmt.Println(colors)
	fmt.Println("len", len(colors))

	// copied_colors := make([]string, 10)

	// copy(copied_colors, colors) // copy from colors to copied_colors
	// fmt.Println(copied_colors)

	var data1 = [...]int{1, 2, 3}
	fmt.Println(data1) // 0 0 0 0 0
	var data2 []int
	fmt.Println(data2)
	data2 = append(data2, 5)
	fmt.Println(data2)
}
