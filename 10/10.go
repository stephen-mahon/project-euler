package main

import "fmt"

func main() {
	i, n := 1, 2
	var total int
	valsBelow := 2000000
	for i < valsBelow {
		if isPrime(i, n) {
			total += i
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
