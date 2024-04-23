package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"cb", "ca", "ab", "ba"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting `int`s.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	fl := []float64{30.5, 12.4, 20.5}
	sort.Float64s(fl)
	fmt.Println("float", fl)

	// We can also use `sort` to check if a slice is
	// already in sorted order.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)

	// bagaimana cara melakukan pengurutan dari besar ke kecil?

}
