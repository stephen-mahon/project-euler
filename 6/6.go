package main

import "fmt"

func main() {
	val := 100
	vals := createArray(val)
	sumSquares, squareSums := squareofSums(vals), sumofSquares(vals)
	fmt.Printf("First %v natural numbers\nSquare of Sums: %v\nSum of Squares: %v\nDifference = %v", val, sumSquares, squareSums, sumSquares-squareSums)
}

func createArray(n int) []int {
	var ns []int
	for i := 1; i <= n; i++ {
		ns = append(ns, i)
	}
	return ns
}

func sumofSquares(ns []int) int {
	var sumN int
	for _, n := range ns {
		sumN += n * n
	}
	return sumN
}

func squareofSums(ns []int) int {
	var sumN int
	for _, n := range ns {
		sumN += n
	}
	return sumN * sumN
}
