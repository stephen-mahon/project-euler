package main

import "fmt"

func main() {
	var vals []int
	var val int
	i := 3000
	for len(vals) < 500 {
		val = sumNatNumber(i)
		vals = factor(val)
		i++
	}
	fmt.Printf("[%v] %v:\t%v, %v\n", i, val, vals, len(vals))

}

func factor(n int) []int {
	var s []int
	if isPrime(n, 2) {
		s = append(s, 1)
		s = append(s, n)
		return s
	}
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			s = append(s, i)
		}
	}
	s = append(s, n)
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

func sumNatNumber(n int) int {
	var total int
	for i := 0; i <= n; i++ {
		total += i
	}

	return total
}
