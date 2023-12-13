package main

import "fmt"

func main() {
	inp := []int{1, 2, 1, 9, 8, 9, 3}
	var minIndex, maxIndex, max, min int
	if len(inp) > 2 {
		for i := 0; i < len(inp)-1; i++ {
			tempMax, indMax := findMax(inp[i], inp[i+1], i)
			tempMin, indMin := findMin(inp[i], inp[i+1], i)
			if tempMax > max { 
				max = tempMax     
				maxIndex = indMax 
			}
			if tempMin < min {
				min = tempMin
				minIndex = indMin
			}
			//fmt.Println("Finally:", minIndex, maxIndex)
		}

	}
	fmt.Println("Smallest subsequence with min and Max values is:",inp[minIndex:maxIndex+1])
}
func findMin(a, b, index int) (int, int) {
	if a < b {
		return a, index
	}
	return b, index + 1
}
func findMax(a, b, index int) (int, int) {
	if a > b {
		return a, index
	}
	return b, index + 1
}
