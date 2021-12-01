package main

import (
	"flag"
	"fmt"
	"math"
)

var i, sum int = 3, 2

func main() {
	vals := flag.Int("v", 10, "vals below")
	flag.Parse()
	fmt.Println("The sum of all the primes below", *vals)
	for i < *vals {
		if isPrime(int64(i)) {
			sum += i
		}
		i += 2

	}

	fmt.Printf("= %v\n", sum)
}

func isPrime(x int64) bool {
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}

	var i, z int64 = 3, int64(math.Sqrt(float64(x)))
	for ; i <= z; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
