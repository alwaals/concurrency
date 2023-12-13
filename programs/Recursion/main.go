package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println(callSum(4))
	// fmt.Println(callFactorial(3))
	 fmt.Println(findAthFibonacci(5))
	// fmt.Println(palindromeRecursion([]int{1, 2, 3, 2, 2, 1}, 0, 6))
	fmt.Println(searchFileDirectories("/Users/swetha/concurrency/deepdive/trace/","main.go"))
}
func searchFileDirectories(dir string,fileName string) bool {
	info, err := os.Stat(dir)
	fmt.Println(info.IsDir())
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
func palindromeRecursion(n []int, startIndex, endIndex int) bool {
	if startIndex == len(n)/2 {
		return true
	}
	if n[startIndex] != n[endIndex-1] {
		return false
	}
	startIndex++
	endIndex--
	return palindromeRecursion(n, startIndex, endIndex)
}
func findAthFibonacci(A int) int {
	if A == 1 {
        return 1
    }
    if A == 0 {
        return 0
    }
    return findAthFibonacci(A-1) + findAthFibonacci(A-2)
}
func callSum(n int) int {
	if n == 0 {
		return n
	}
	return n + callSum(n-1)
}
func callFactorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * callFactorial(n-1)
}
