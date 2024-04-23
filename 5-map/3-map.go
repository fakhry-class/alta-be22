package main

import "fmt"

func main() {
	var datemonth = map[string]int{}
	datemonth["januari"] = 30
	datemonth["februari"] = 29
	datemonth["maret"] = 31
	datemonth["april"] = 30

	value, isExist := datemonth["juni"]
	fmt.Println("val", value, "exist", isExist)
	if isExist == false {
		datemonth["juni"] = 31
		fmt.Println("datanya tidak ada")
	} else {
		fmt.Println("datanya ada cuy")
	}

	var datemonth1 = map[string]bool{}
	datemonth1["januari"] = true
	datemonth1["februari"] = false
	datemonth1["maret"] = true
	datemonth1["april"] = false
	value1, isExist1 := datemonth1["juni"]
	fmt.Println("val", value1, "exist", isExist1)

	for _, v := range datemonth1 {
		fmt.Println(v)
	}
}
