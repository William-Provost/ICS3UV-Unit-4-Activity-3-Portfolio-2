// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-30
// Fileoverview: This program stores 10 decimal numbers in an array and finds the largest.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var numbers [10]float64
	var largest float64
	scanner := bufio.NewScanner(os.Stdin)

	// get 10 numbers from user
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		scanner.Scan()
		numStr := strings.TrimSpace(scanner.Text())
		num, _ := strconv.ParseFloat(numStr, 64)
		numbers[i] = num
	}

	// find largest number
	largest = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > largest {
			largest = numbers[i]
		}
	}

	// display largest number
	fmt.Printf("\nThe largest value in your list is: %.1f\n", largest)

	// display all numbers
	fmt.Print("Here is the list of numbers: ")
	for i := 0; i < len(numbers); i++ {
		if i != len(numbers)-1 {
			fmt.Printf("%.1f, ", numbers[i])
		} else {
			fmt.Printf("%.1f", numbers[i])
		}
	}
	fmt.Println("\nDone.")
}
