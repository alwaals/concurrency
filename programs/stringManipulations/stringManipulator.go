package main

import (
	//"bytes"
	"bytes"
	"fmt"
)

const p = 3.14*10
func main() {
	// fmt.Println("Reverse of powder:", reverse("powder"))
	// fmt.Println("Is powder palindrome:", palindrome("powder"))
	// fmt.Println("Is powder madam:", palindrome("madam"))
	// fmt.Println(`Reversing "This is a string":`, reverseStrings("This is string"))
	// fmt.Println("Is abbc is palindrome", swapToMakePalin("abbc"))
	// fmt.Println("Is nitin is palindrome", swapToMakePalin("nitin"))
	// fmt.Println("Which comes first", findLexicographic("Vikaas", "Saurav"))
	//fmt.Println("Which comes first", findLexicographicSwap("abzxdbxa"))
	fmt.Println("", twoStrsAnagrams("abcda", "dbaca"))
	fmt.Println(p)
}
func twoStrsAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	mp := make(map[rune]int)
	for _, s := range str1 {
		value, exists := mp[s]
		if exists {
			mp[s] += value
			continue
		}
		mp[s] = 1
	}
	fmt.Println(mp)
	for _, v := range str2 {
		value, exists := mp[v]
		if exists {
			mp[v]--
			value -= 1
		}
		if value == 0 {
			delete(mp, v)
		}
		fmt.Println(mp)
	}
	return len(mp) == 0
}
func findLexicographicSwap(a string) string {
	freq := make(map[rune]int)
	for i := 0; i < len(a); i++ {
		if value, exists := freq[rune(a[i])]; exists {
			freq[rune(a[i])] += value
			continue
		}
		freq[rune(a[i])] = 1
	}
	fmt.Println(freq)
	var max, least int32
	least = rune(a[0])
	for i := 0; i < len(a); i++ {
		value, _ := freq[rune(a[i])]
		if max < rune(a[i])-int32(value) {
			max = rune(a[i]) - int32(value)
		}
		if least > rune(a[i])-int32(value) {
			least = rune(a[i]) - int32(value)
		}
	}
	fmt.Println("max,least:", max, least)
	return ""
}
func findLexicographic(a, b string) string {
	if len(a) != len(b) {
		if len(a) < len(b) {
			return a
		}
		return b
	}
	if a[0] > b[0] {
		return b
	} else if a[0] < b[0] {
		return a
	} else {
		for i := 1; i < len(a); i++ {
			if a[i] < b[i] {
				return a
			} else if a[i] > b[i] {
				return b
			}
		}
	}
	return ""
}
func swapToMakePalin(str string) bool {
	c := 0
	for i := 0; i < len(str)/2; i++ {
		lastIndex := len(str) - i - 1
		if str[i] != str[lastIndex] {
			c++
		}
	}
	return c == 1
}
func reverseStrings(str string) string {
	rev := reverse(str)
	rev = rev + " "
	var res bytes.Buffer
	var s1 []rune
	for i := 0; i < len(rev); i++ {
		if rev[i] == ' ' {
			res.WriteString(reverse(string(s1)))
			res.WriteString(" ")
			s1 = []rune{}
		} else {
			s1 = append(s1, rune(rev[i]))
		}
	}
	return res.String()
}
func palindrome(str string) bool {
	r := []rune(str)
	for i := 0; i < len(str)/2; i++ {
		if r[i] != r[len(str)-i-1] {
			return false
		}
	}
	return true
}
func reverse(str string) string {
	var st []rune
	st = []rune(str)
	for i := 0; i < len(str)/2; i++ {
		lastIndex := len(str) - i - 1
		st[i], st[lastIndex] = st[lastIndex], st[i]
	}
	return string(st)
}
