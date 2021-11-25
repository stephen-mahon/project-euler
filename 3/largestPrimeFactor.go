package main

import "fmt"

func main() {
	primeFactors := factorPrime(13195)
	maxPrime := primeFactors[len(primeFactors)-1]
	fmt.Print(maxPrime)
}

func factorPrime(n int) []int {
	var s []int
	if isPrime(n, 2) {
		s = append(s, 1)
		s = append(s, n)
		return s
	}
	for i := 1; i <= n/2; i++ {
		if n%i == 0 && isPrime(i, 2) {
			s = append(s, i)
		}
	}
	return s
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
