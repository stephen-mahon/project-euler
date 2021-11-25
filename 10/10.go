package main

import "fmt"

var i, n int = 1, 2

func main() {
	var total int
	valsBelow := 2000000
	for i < valsBelow {
		if isPrime(i, n) {
			total += i
			//fmt.Printf("+%v = %v\n", i, total)
		}
		i++

	}

	fmt.Printf("%v\n", total)
}

func isPrime(p, n int) bool {

	if p == 0 || p == 1 {
		return false
	}

	if p == n {
		return true
	}

	if p%n == 0 {
		return false
	}
	n++
	return isPrime(p, n)
}
