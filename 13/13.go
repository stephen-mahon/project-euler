package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var title = "Project Euler: 13 Large sum"

func main() {
	fmt.Printf("%s\n\n", title)

	fileName := flag.String("file name", "largeSum.txt", "the file name for the largest sum")

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("file not found: %s", err)
	}
	defer file.Close()

	var total float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if val, err := strconv.ParseFloat(scanner.Text(), 64); err == nil {
			total += val
		}
	}

	val := total / math.Pow10(int(math.Log10(total)))

	strVal := strconv.Itoa(int(val))
	for len(strVal) < 10 {
		val *= 10
		strVal = strconv.Itoa(int(val))
	}

	fmt.Println("ans:", strVal)
}
