//P1Q1

package main

import (
	"fmt"
	"strconv"
)

// Convert integer to binary and store in a slice
func intToBinary(n int) []string {
	binarySlice := make([]string, n+1)
	for i := 0; i <= n; i++ {
		binarySlice[i] = strconv.FormatInt(int64(i), 2)
	}
	return binarySlice
}

func main() {
	n := 3
	result := intToBinary(n)
	fmt.Println(result)
}
