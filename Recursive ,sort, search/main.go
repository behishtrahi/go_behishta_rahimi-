// P1 Q1

package main

import "fmt"

func main() {
	fmt.Println(sumfibX(5))  // 12
	fmt.Println(sumfibX(10)) //143
}

func sumfibX(n int) int {

	if n < 0 {
		return 0
	}
	fib := []int{0, 1}
	sum := 1
	for i := 2; i <= n; i++ {

		nextfib := fib[i-1] + fib[i-2]
		fib = append(fib, nextfib)
		sum += nextfib
	}
	return sum

}
