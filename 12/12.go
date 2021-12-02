package main

import (
	"flag"
	"fmt"

	"modernc.org/mathutil"
)

var title = "Highly divisible triangular number"

type FactorTerm struct {
	Prime uint32 // The divisor
	Power uint32 // Term == Prime^Power
}

func main() {
	fmt.Println(title)
	div := flag.Int("d", 10, "Largest number of divisors")
	flag.Parse()
	//var vals []FactorTerm
	var triNum, i int
	for i < *div {
		i++
		triNum += i
		fmt.Println(i, mathutil.FactorInt(uint32(triNum)))
	}
	//fmt.Printf("%v: %v\t%v\n", i, triNum, vals)
}
