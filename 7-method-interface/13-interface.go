package main

import "fmt"

type hewan interface {
	makan() string
	bergerak(gaya string) string
}

type hidup interface {
	bernapas() string
}

type kucing struct {
	Jenis   string
	Makanan string
	Habitat string
}

type burung struct {
	Jenis   string
	Makanan string
	Habitat string
}

// bergerak implements hewan.
func (b burung) bergerak(gaya string) string {
	return b.Jenis + " bergerak dengan cara " + gaya
}

// makan implements hewan.
func (b burung) makan() string {
	return b.Jenis + " makanannya adalah " + b.Makanan
}

// bergerak implements hewan.
func (k kucing) bergerak(gaya string) string {
	return k.Jenis + " bergerak dengan cara " + gaya
}

// makan implements hewan.
func (k kucing) makan() string {
	return k.Jenis + "makanannya adalah " + k.Makanan
}

// bernapas implements hidup.
func (k kucing) bernapas() string {
	if k.Habitat == "darat" || k.Habitat == "udara" {
		return k.Jenis + " bernapas dengan paru-paru"
	} else if k.Habitat == "laut" {
		return k.Jenis + " bernapas dengan insang"
	} else {
		return "tidak diketahui"
	}
}

func main() {
	var anggora = kucing{
		Jenis:   "anggora abu",
		Makanan: "wiskas",
		Habitat: "darat",
	}

	var IKucing1 hewan
	IKucing1 = anggora
	fmt.Println(IKucing1.makan())
	fmt.Println(IKucing1.bergerak("jalan"))

	var IKucing2 hidup
	IKucing2 = anggora
	fmt.Println(IKucing2.bernapas())

	var merpati = burung{
		Jenis:   "merpati dara",
		Makanan: "biji-bijian",
		Habitat: "pohon",
	}
	var IBurung hewan
	IBurung = merpati
	fmt.Println(IBurung.makan())
	fmt.Println(IBurung.bergerak("terbang"))
}
