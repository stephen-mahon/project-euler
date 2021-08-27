package main

import (
	"fmt"
)

func main() {
	numberOfSundays()

}

func numberOfSundays() {

	var daysOfWeek = map[int]string{
		0: "Monday",
		1: "Tuesday",
		2: "Wednesday",
		3: "Thursday",
		4: "Friday",
		5: "Saturday",
		6: "Sunday",
	}

	var numYears, extraDays int
	for i := 1901; i <= 2000; i++ {
		numYears++
		if i%4 == 0 || (i%1000 == 0 && i%400 == 0) {
			extraDays++
		}
	}

	numDays := (numYears * 365) + extraDays

	fmt.Println("Jan 1, 2001 was a", daysOfWeek[(numDays+1)%7])
}
