package main

import (
	"fmt"
)

var title = "Project Euler\nProblem 14: Longest Collatz sequence"

func main() {
	fmt.Println(title)
	var ans [2]int
	for i := 1; i <= 1000000; i++ {
		count := collatzSeq(i)
		if count > ans[1] {
			ans[0], ans[1] = i, count
		}
	}
	fmt.Println("Starting number", ans[0], "has a chain of", ans[1], "numbers")
}

func isEven(n int) bool {
	if n%2 != 0 {
		return false
	}
	return true
}

func collatzSeq(n int) int {
	var count int
	for n != 1 {
		count++
		if isEven(n) {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return count
}
