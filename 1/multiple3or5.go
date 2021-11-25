package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 && args[0] == "/help" {
		fmt.Println("Given a number returns the sum of all the natural numbers below that number that are multiples of 3 or 5.")
		fmt.Println("Usage: multiple3or 5 <number>")
		fmt.Print("Example:\t% ./multiple3or5 10\n\t\t% 23\n")
	} else {
		if len(args) != 1 {
			fmt.Println("You must enter only one argument! Type /help for help.")
		} else {
			val, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Printf("input must be a nummber!\n")
				log.Fatal(err)
			}
			fmt.Printf("%v\n", multiple3or5(val))
		}
	}

}

func multiple3or5(num int) int {
	var sum int
	for i := 1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
