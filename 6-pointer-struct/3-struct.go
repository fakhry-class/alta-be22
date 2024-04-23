package main

import "fmt"

type student struct {
	Id      string
	Name    string
	Height  float32
	Phone   []string
	Address address
}

type address struct {
	Street string
	City   string
	Number int
}

func main() {
	var data1 student
	data1.Id = "0001"
	data1.Name = "Fulan"
	data1.Height = 180
	data1.Phone = append(data1.Phone, "0812345")
	data1.Phone = append(data1.Phone, "0812346")
	// var addrs1 address = address{
	// 	Street: "jl lurus terus",
	// 	City:   "surabaya",
	// 	Number: 1,
	// }

	// data1.Address = addrs1
	data1.Address = address{
		Street: "jl lurus terus",
		City:   "surabaya",
		Number: 1,
	}

	for _, v := range data1.Phone {
		fmt.Println("phone", v)
	}
}
