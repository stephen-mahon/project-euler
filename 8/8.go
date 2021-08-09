package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("1000digits.txt")
	if err != nil {
		fmt.Printf("file of 1000 digits not found! Please see README.\n")
		log.Fatal(err)
	}
	defer file.Close()

	var numStr string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numStr = scanner.Text()
	}

	var digits []int
	for _, v := range numStr {
		n, _ := strconv.Atoi(string(v))
		digits = append(digits, n)
	}

	var largestProduct, temp int
	adjNum := 13
	for i := 0; i < len(digits)-adjNum; i++ {
		temp = 1
		for j := 0; j < adjNum; j++ {
			temp *= digits[i+j]
		}
		if temp > largestProduct {
			largestProduct = temp
		}
	}

	fmt.Println(largestProduct)
}
