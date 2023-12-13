package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "jainey1rohan2amisha3jain"
	vowelsCount, numbersCount := 0, 0
	for _, v := range str {
		switch rune(v) {
		case 'a', 'e', 'i', 'o', 'u':
			vowelsCount++
		case '0','1','2','3','4','5','6','7','8','9':
			numbersCount++
		}
	}
	fmt.Println(vowelsCount,numbersCount)
	matched,_:=regexp.MatchString("[A-Z]",str)
	fmt.Println(matched)
}
