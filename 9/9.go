package main

import "fmt"

func main() {

	fmt.Printf("a\tb\tc\ta+b+c\ta*b*c\n")
	for i := 1; i < 998; i++ {
		for j := 1; j < 998; j++ {
			for k := 1; k < 998; k++ {
				if pythagoreanTriplet(i, j, k) {
					fmt.Printf("%v\t%v\t%v\t%v\t%v\n", i, j, k, i+k+j, i*j*k)
				}
			}
		}
	}
}

func pythagoreanTriplet(a, b, c int) bool {

	if a+b+c == 1000 && (a*a)+(b*b) == (c*c) {
		return true
	}
	return false
}
