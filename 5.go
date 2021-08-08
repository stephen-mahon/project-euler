package main

import "fmt"

func main() {
	fmt.Println(smallestMultiple(20))
}

func smallestMultiple(n int) int {
	ans := n * 2
	modSum := -1
	for modSum != 0 {
		ans += 1
		modSum = 0
		for i := n; i > 0; i-- {
			modSum += ans % i
		}
	}
	return ans
}
