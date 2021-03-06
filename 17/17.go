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
	count := flag.Int("count", 50, "number count")
	total := 0
	for i := 1; i <= *count; i++ {
		if i <= 20 {
			num := toTwenty(i)
			fmt.Println(num)
			total += wordCount(num)
		}
	}
	fmt.Println(total)
}

func wordCount(s string) int {
	return len(s)
}

func toTwenty(i int) string {
	return numWords[i]
}
