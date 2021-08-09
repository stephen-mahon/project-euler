package main

import "fmt"

func main() {
	i, n := 1, 2
	lastPrime := 10001
	var primes []int
	for len(primes) < lastPrime {
		if isPrime(i, n) {
			primes = append(primes, i)
		}
		i++
	}
	fmt.Println(primes[lastPrime-1])
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
