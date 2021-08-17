package main

import "fmt"

func main() {
	fmt.Println(palindroneProduct())

}

func palindroneProduct() (int, int, int) {
	var val1, val2, largestProduct int
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			val := fmt.Sprint(i * j)
			if val == reverse(val) && (i*j > largestProduct) {
				val1 = i
				val2 = j
				largestProduct = i * j
			}
		}
	}
	return val1, val2, largestProduct
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
