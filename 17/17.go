package main

import (
	"flag"
	"fmt"
)

var title = "Project Euler\nProblem 17: Number letter counts"

var numWords = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",
	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	18: "eighteen",
	19: "nineteen",
	20: "twenty",
	30: "thirty",
	40: "forty",
	50: "fifty",
	60: "sixty",
	70: "seventy",
	80: "eighty",
	90: "ninety",
}

func main() {
	fmt.Println(title)
	count := flag.Int("count", 20, "number count")
	total := 0
	for i := 1; i <= *count; i++ {
		val := 0
		switch i {
		case i <= 20:
			val = toTwenty(i)
		default:
			val = 0
		}
		total += val
	}
	fmt.Println(total)
}

func toTwenty(i int) int {
	val := 0
	for j := range numWords[i] {
		_ = j
		val += 1
	}
	return val
}
