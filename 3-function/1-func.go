package main

import (
	"fakhry/be22/helpers"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	sayHelloWorld()
	var hasil = helpers.CetakHello()
	fmt.Println(helpers.TampilHello())
	var hasilBintang = helpers.TampilBintang()
	fmt.Println("hasilnya", hasilBintang)
	var hasilPenambahan = helpers.HitungPenambahan()
	fmt.Println("hasilpenambahan:", hasilPenambahan)
	var hasilPenambahanParam = helpers.HitungPenambahanWithParam(50, 60)
	fmt.Println("hasilpenambahanparam:", hasilPenambahanParam)
	fmt.Println("proses lagi", hasilPenambahanParam+10)
	if hasilPenambahanParam > 10 {
		// ...
	}

	// helpers.printHello()

}

func sayHelloWorld() {
	fmt.Println("hello world from function")
	fmt.Println("hello world from function 2")
}
