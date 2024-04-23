package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

type Mobil struct {
	Type string
}

func fullNameFunc(firstName string, lastName string) (fullName string) {
	fullName = firstName + " " + lastName
	return
}

func (em Employee) fullName(gelar string) string {
	result := em.FirstName + " " + em.LastName + ". " + gelar
	return result
}

func main() {
	emp1 := Employee{
		FirstName: "Ross",
		LastName:  "Geller",
	}

	emp2 := Employee{
		FirstName: "bambang",
		LastName:  "pamungkas",
	}

	// mobil1 := Mobil{
	// 	Type: "sedan",
	// }
	// mobil1.fullName("skom")

	fmt.Println(fullNameFunc(emp1.FirstName, emp1.LastName))

	fmt.Println(emp1.fullName("s.kom"))
	fmt.Println(emp2.fullName("m.eng")) // bambang pamungkas
}
