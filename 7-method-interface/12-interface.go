package main

import (
	"fmt"
	"math"
)

// declare interface
type calculate interface {
	luas() int
	ExtendLuas(nominal int) int
}

type persegi struct {
	sisi int
}

type lingkaran struct {
	jari int
}

func (s persegi) luas() int {
	return s.sisi * s.sisi
}

func (s persegi) keliling() int {
	return 4 * s.sisi
}
func (s persegi) ExtendLuas(bil int) int {
	return (s.sisi * s.sisi) * bil
}

// luas implements calculate.
func (l lingkaran) luas() int {
	result := math.Pi * float64(l.jari) * float64(l.jari)
	return int(result)
}

// ExtendLuas implements calculate.
func (l lingkaran) ExtendLuas(nominal int) int {
	result := math.Pi * float64(l.jari) * float64(l.jari)
	hasil := result * float64(nominal)
	return int(hasil)
}

func main() {
	p1 := persegi{
		sisi: 50,
	}
	var dimResult calculate
	dimResult = p1
	fmt.Println("large square :", dimResult.luas())
	fmt.Println("extend luas persegi :", dimResult.ExtendLuas(2))

	l1 := lingkaran{
		jari: 5,
	}

	var ILingkaran calculate
	ILingkaran = l1
	fmt.Println("large lingkaran :", ILingkaran.luas())
}
