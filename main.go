//Luka Brown
//Collatz length function
package main

import (
	"fmt"
)

//brings num to 1 using collatz conjecture
func collatzLength(n int) int {
	length := 1
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		length++
	}
	return length
}

//main
func main() {
	maxLength := 0
	maxNumber := 0

  //tests nums 1-10000 to find longest length
	for i := 1; i <= 10000; i++ {
		length := collatzLength(i)
		if length > maxLength {
			maxLength = length
			maxNumber = i
		}
	}

	fmt.Printf("The number %d has the longest Collatz sequence of length %d.\n", maxNumber, maxLength)
}
