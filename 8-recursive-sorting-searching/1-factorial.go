package main

import "fmt"

/*
factorial:
6! = 6*5*4*3*2*1 =720
6! = 6*5!

5! = 5*4*3*2*1 = 120
5! = 5*4!

4! = 4*3*2*1 = 24
3! = 3*2*1

	3*2!

2! 2*1!
1! = 1
*/

func hitung() {

}

// loop biasa
func factorial1(n int) int {
	var result = 1
	for i := n; i >= 1; i-- {
		result = result * i
	}
	return result
}

// recursive
func factorial2(n int) int {
	if n == 1 { // base case
		return 1
	} else {
		// recurrent relation
		var result = factorial2(n - 1)
		return n * result
		/*
			n = 6
			6*5! --> 6*120 = 720
				5*4! --> 5*24 = 120
					4*3! --> 4*6 = 24
						3*2!	--> 3*2 = 6
							2*1! --> 2*1 =2
								1! = return 1

			iterasi 1: result = f(1)
			return 2*f(1)
			it 2: 1
			2*1
		*/
	}
}

func main() {
	fmt.Println("loop", factorial1(5))
	fmt.Println("recursive", factorial2(5))
}
