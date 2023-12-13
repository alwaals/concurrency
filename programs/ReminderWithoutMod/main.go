package main

import "fmt"

func main() {
	fmt.Println(findPositiveReminder(28,4))
	fmt.Println(findNegativeReminder(-30,7))
}
func findNegativeReminder(n int, mod int) int {
	rem:=-n
	for rem > mod {
		rem =rem - mod
		
	}
	fmt.Println(rem,n)
	return n-rem
}
func findPositiveReminder(n int, mod int) int {
	for n >= mod {
		n = n - mod
	}
	return n
}
