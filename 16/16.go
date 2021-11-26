package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

var title = "Project Euler\nProblem 16: Power digit sum"

func main() {
	n := flag.Int("2^n", 1000, "n is the base 2 exponent")
	fmt.Println(title)
	ans := fmt.Sprintf("%.f", math.Pow(2, float64(*n)))
	fmt.Println("2^", *n, "=", ans)
	total := 0
	for i := range ans {
		val, err := strconv.Atoi(string(ans[i]))
		if err != nil {
			log.Fatalf("error in string convert %v: %v", string(ans[i]), err)
		}
		total += val
	}
	fmt.Println("Power digit sum =", total)
}
