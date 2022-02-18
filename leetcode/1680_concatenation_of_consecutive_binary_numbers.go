package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Max(1, 2))
	fmt.Println(concatenatedBinary(12))
}
func concatenatedBinary(n int) int {
	if n == 1 {
		return 1
	}

	count := bitCount(n)
	return (concatenatedBinary(n-1)<<count + n) % (int)(math.Pow(10, 9)+7)
}

func bitCount(n int) int {
	count := 0
	for n != 0 {
		n = n >> 1
		count++
	}
	return count
}
