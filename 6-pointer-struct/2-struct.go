package main

import "fmt"

type mahasiswa struct {
	Nim      string
	Nama     string
	Fakultas string
	DOB      string
	IPK      float32
}

func main() {
	var dataMHS []mahasiswa
	fmt.Println("awal:", dataMHS)
	var mhs1 = mahasiswa{
		Nim:      "MHS0001",
		Nama:     "Andi",
		Fakultas: "FILKOM",
		DOB:      "01-01-2000",
		IPK:      4.00,
	}

	var mhs2 = mahasiswa{
		Nim:      "MHS0002",
		Nama:     "Budi",
		Fakultas: "MIPA",
		DOB:      "02-01-2000",
		IPK:      3.9,
	}

	dataMHS = append(dataMHS, mhs1)
	dataMHS = append(dataMHS, mhs2)
	fmt.Println("hasil:", dataMHS)
	fmt.Println("hasil:", dataMHS[1])
}
