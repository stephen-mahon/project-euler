package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("grid.txt")
	if err != nil {
		fmt.Printf("file of 20x20 grid not found! Please see README.\n")
		log.Fatal(err)
	}
	defer file.Close()

	var strLn []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strLn = append(strLn, scanner.Text())
	}

	grid := make([][]int, 20)
	index := -1
	for _, v := range strLn {
		strSplit := strings.Split(v, " ")
		for i, u := range strSplit {
			val, _ := strconv.Atoi(string(u))
			if i%20 == 0 {
				index++
			}
			grid[index] = append(grid[index], val)

		}
	}
	var a, b, c, d int
	var greatestProduct int

	// This could be much better - lots of redundency (abcd = dcba)
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			// Right
			if j < 17 {
				a = grid[i][j]
				b = grid[i][j+1]
				c = grid[i][j+2]
				d = grid[i][j+3]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
			// Left
			if j > 3 {
				a = grid[i][j-3]
				b = grid[i][j-2]
				c = grid[i][j-1]
				d = grid[i][j]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
			// Down
			if i < 17 {
				a = grid[i][j]
				b = grid[i+1][j]
				c = grid[i+2][j]
				d = grid[i+3][j]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
			// Up
			if i > 3 {
				a = grid[i-3][j]
				b = grid[i-2][j]
				c = grid[i-1][j]
				d = grid[i][j]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
			// Diagonal Right
			if i < 17 && j < 17 {
				a = grid[i][j]
				b = grid[i+1][j+1]
				c = grid[i+2][j+2]
				d = grid[i+3][j+3]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
			// Diagonal Left
			if i > 3 && j > 3 {
				a = grid[i-3][j-3]
				b = grid[i-2][j-2]
				c = grid[i-1][j-1]
				d = grid[i][j]
				if a*b*c*d > greatestProduct {
					greatestProduct = a * b * c * d
					fmt.Printf("%v %v %v %v %v\n", a, b, c, d, a*b*c*d)
				}
			}
		}
	}
}
