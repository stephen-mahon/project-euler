package main

import "fmt"

func main() {
	evenFibSum := fibSequence(5000000)
	fmt.Println(evenFibSum)
}

func fibSequence(maxVal int) int {
	a := []int{1, 2}
	var next, sum int
	for next < maxVal {
		next = a[0] + a[1]
		if a[0]%2 == 0 {
			sum += a[0]
		}
	}
	return sum + a[0]

}
