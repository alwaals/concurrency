package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	inp := []int{2, 7, 11, 10}
	target := 12

	indexes := usingMap(inp, target)
	fmt.Println("found indexes:", indexes)
	fmt.Println("Total time taken:", time.Since(t))
}
func usingMap(inp []int, target int) []int {
	indexes := []int{}
	mp := make(map[int]int)
	for index, value := range inp {
		if key, exists := mp[target-value]; exists {
			indexes = append(indexes, key, index)
		}
		mp[value] = index
	}

	return indexes
}
func findTwoSum(inp []int, target int) []int {
	indexes := []int{}
	if len(inp) > 1 {
		for i := 0; i < len(inp); i++ {
			find := target - inp[i]
			for j := i + 1; j < len(inp); j++ {
				if find == inp[j] && i != j {
					indexes = append(indexes, i, j)
					break
				}
			}
		}
	}
	return indexes
}
