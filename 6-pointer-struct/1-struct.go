package main

import "fmt"

type person struct {
	FirstName string
	LastName  string
	Age       uint
	Address   string
	Gender    string
	Weight    float32
	isActive  bool
}

func main() {
	// long declaration
	var orang1 person
	orang1.FirstName = "andi"
	orang1.LastName = "abdul"
	orang1.Age = 20
	orang1.Address = "jakarta"
	orang1.Gender = "laki-laki"
	orang1.Weight = 60.5
	fmt.Println(orang1)
	fmt.Println("nama", orang1.FirstName)

	var data2 person
	data2.FirstName = "fulan"
	data2.Age = 25
	fmt.Println("nama:", data2.FirstName)
	fmt.Println("weight:", data2.Weight)
	fmt.Println("aktif:", data2.isActive)

	// long declaration with assigned value
	var orang2 = person{"Rizky", "Kurniawan", 26, "surabaya", "laki-laki", 70, true}
	fmt.Println(orang2)

	// long declaration with assigned value each name fields
	var orang3 = person{
		LastName:  "Umam",
		FirstName: "Iswanul",
		Age:       25,
	}
	fmt.Println(orang3)

	//short declaration
	orang4 := person{"Rizky", "Kurniawan", 26, "surabaya", "laki-laki", 70, true}
	fmt.Println(orang4)

	// short declaration with new keyword
	orang5 := new(person)
	orang5.FirstName = "Muhammad"
	orang5.LastName = "Ismail"
	orang5.Age = 30
	fmt.Println(*orang5)
}
