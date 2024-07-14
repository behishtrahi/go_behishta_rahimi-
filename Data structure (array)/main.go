// priority 1 q1

package main

import "fmt"

func main() {
	res1 := merge([][]int{
		{1, 2, 5, 4, 3},
		{1, 2, 7, 8},
	})
	fmt.Println(res1) // [1,2,5,4,3,7,8]

	res2 := merge([][]int{
		{1, 2, 5, 4, 5},
		{7, 9, 10, 5},
	})
	fmt.Println(res2) // [1, 2, 5, 4, 7, 9, 10]

	res3 := merge([][]int{
		{1, 4, 5},
		{7},
	})
	fmt.Println(res3) // [1, 4, 5, 7]
}

func merge(data [][]int) []int {
	elementMap := make(map[int]bool)
	var result []int

	for _, row := range data {
		for _, elem := range row {
			if !elementMap[elem] {
				elementMap[elem] = true
				result = append(result, elem)
			}
		}
	}

	return []int{}
}
