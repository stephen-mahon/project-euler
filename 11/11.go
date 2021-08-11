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
	squareArray := squares(grid)
	var maxProduct int
	for i := range squareArray {
		if maxAdjacentProduct(squareArray[i]) > maxProduct {
			maxProduct = maxAdjacentProduct(squareArray[i])
		}
	}
	fmt.Println(maxProduct)

}

func maxAdjacentProduct(square []int) int {
	maxProduct := square[0] * square[5] * square[10] * square[15]

	if square[3]*square[6]*square[9]*square[12] > maxProduct {
		maxProduct = square[3] * square[6] * square[9] * square[12]
	}

	for i := range square {
		if i%4 == 0 {
			a := i
			b := i + 1
			c := i + 2
			d := i + 3

			if square[a]*square[b]*square[c]*square[d] > maxProduct {
				maxProduct = square[a] * square[b] * square[c] * square[d]
			}
		}
	}

	for i := 0; i < 4; i++ {
		a := i
		b := i + 4
		c := i + 8
		d := i + 12
		if square[a]*square[b]*square[c]*square[d] > maxProduct {
			maxProduct = square[a] * square[b] * square[c] * square[d]
		}

	}

	return maxProduct
}

func squares(grid [][]int) [][]int {
	squareArray := make([][]int, 17*17)
	var index int
	for i := 0; i < len(grid)-3; i++ {
		for j := 0; j < len(grid[i])-3; j++ {
			for k := 0; k < 4; k++ {
				for l := 0; l < 4; l++ {
					squareArray[index] = append(squareArray[index], grid[i+k][j+l])
				}
			}
			index++
		}
	}
	return squareArray
}
