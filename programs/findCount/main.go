package main

import (
	"fmt"
)

func main() {
	// inp:=[]int{2,2,4,4,4,7,7,7,8}
	// //count:=0
	// found :=false
	// for i:=0;i<len(inp)-1;i++{
	// 	if inp[i]==inp[i+1]{
	// 		continue
	// 	}
	// 	if inp[i]==len(inp)-i-1{
	// 		found=true
	// 	}
	// }
	// fmt.Println(found)

	//find pair which gives min diff
	// inp := []float64{3, 11, 15, 21, 28, 32}
	// min := math.Abs(inp[0] - inp[1])
	// if len(inp) > 2 {
	// 	for i := 1; i < len(inp)-1; i++ {
	// 		if math.Abs(inp[i]-inp[i+1]) < math.Abs(inp[i]-inp[i-1]) {
	// 			if min > math.Abs(inp[i]-inp[i+1]) {
	// 				min = math.Abs(inp[i] - inp[i+1])
	// 			}
	// 		} else if min > math.Abs(inp[i]-inp[i-1]) {
	// 			min = math.Abs(inp[i] - inp[i-1])
	// 		}
	// 	}
	// 	fmt.Println(min)
	// }

	///Given an array, check if we can form an Arithmetic Progression using values of array
	// inp := []int{3, 7, 11, 15, 19, 23, 28}
	// fmt.Println(findAP(inp))

	/// Colors prpgram: Given the colors[0,1,2], sort any array in the colors order
	// inp2 := []int{2, 1, 0, 2, 2, 1, 1, 0, 0, 1, 2}
	// fmt.Println(findColorsOrder(inp2))

	// Search a[i]+a[j]=target, if a given pair is found or not
	// inp2 := []int{10,7,8,9,4,1}
	// fmt.Println(searchTargetInGivenArray(inp2,10))

	//Search a[i]+a[j]=target, if a given pair is found or not
	// inp2 := []int{10,7,8,9,4,1}
	// fmt.Println(searchTargetInGivenArray(inp2,10))

	//Find no.of distinct elements in an array
	// inp:=[]int{6,3,7,3,8,6,9}
	// fmt.Println(findDistinct(inp))

	//Given N array, find how many pairs have ar[i]-ar[j]=k such that i!=j
	// inp := []int{3, 8, 5, 10, 6, 12, 4}
	// fmt.Println(findCountOfPairs(inp, 1))

	inp:=[]int{3,-2,1,3,-2}
	fmt.Println(sumOfIndexes(inp,1,4))

}
func sumOfIndexes(inp []int,a,b int)int{
	sum:=0
	// for i:=a;i<=b;i++{
	// 	sum+=inp[i]
	// }
    arr:=[]int{}
	for i:=0;i<len(inp);i++{
		inSum:=0
		for j:=0;j<=i;j++{
	    	inSum+=inp[j]
		}
		arr=append(arr, inSum)
	}
	fmt.Println(arr)

	return sum
}
type mapResp struct {
	value int
	index int
	freq  int
}
type Response struct {
	Key   int
	Value int
}
func findCountOfPairs(inp []int,target int)int{
	mp:=insertIntoHM(inp)
	count:=0
	for i:=0;i<len(inp);i++{
		a:=inp[i]
		b:=a-target
		if value,exists:=mp[b];exists{
			if b!=value{
				count=count+value
			}
		}
	}
	return count
}
func insertIntoHM(inp []int) map[int]int {
	mp := make(map[int]int)
	for _, v := range inp {
		if _, exists := mp[v]; exists {
			mp[v] += 1
			continue
		}
		mp[v] = 1
	}
	return mp
}
func findPairs(inp []int, target int) []Response {
	fmt.Println("Input array:",inp)
	mp := make(map[int]mapResp)
	for i, v := range inp {
		if count, exists := mp[v]; exists {
			mp[v] = mapResp{index: i, freq: count.freq + 1,value: v}
			continue
		}
		mp[v] = mapResp{index: i, freq: 1,value: v}
	}
	fmt.Println(mp, "target:", target)
	indexes := []Response{}
	for _, v := range inp {
		k := v - target
		fmt.Println("value:", v, "diff:", k)
		if value, exists := mp[k]; exists {
			fmt.Println("exists:", exists, "value:", value)
			if value.value != k {
				fmt.Println("Appending",value.index,value.value)
				indexes = append(indexes, Response{Key: value.index, Value: value.value})
			} else {
				if value.freq > 1 {
					fmt.Println("Found more than once:",value.index,value.value)
					indexes = append(indexes, Response{Key: value.index, Value: value.value})
				}
			}
		}
	}
	return indexes
}

func findDistinct(inp []int) []int {
	newArr := []int{}
	mp := make(map[int]bool)
	for _, v := range inp {
		if _, exists := mp[v]; exists {
			continue
		}
		mp[v] = true
		newArr = append(newArr, v)
	}
	return newArr
}
func searchTargetInGivenArray(inp []int, target int) bool {
	mp := make(map[int]int)
	for _, v := range inp {
		if _, exists := mp[v]; exists {
			mp[v] += 1
			continue
		}
		mp[v] = 1
	}
	for i := 0; i < len(inp); i++ {
		a := inp[i]
		b := target - a
		if v, exists := mp[b]; exists {
			if a != b {
				return true
			} else {
				if v > 1 {
					return true
				}
			}
		}
	}
	fmt.Println(mp)
	return false
}
func findColorsOrder(inp []int) []int {
	mp := make(map[int]int)
	newInp := make([]int, len(inp))
	for _, v := range inp {
		if _, exists := mp[v]; exists {
			mp[v] += 1
			continue
		}
		mp[v] = 1
	}
	indexCount := 0
	for i := 0; i < len(mp); i++ {
		v := mp[i]
		for v > 0 {
			newInp[indexCount] = i
			indexCount++
			v--
		}
	}
	return newInp
}
func findAP(inp []int) bool {
	mp := make(map[int]int)
	diff := inp[1] - inp[0]
	min := findMin(inp[1], inp[0])
	var secondMin = inp[len(inp)-1]
	for i := 0; i < len(inp); i++ {
		mp[inp[i]] = 1
		if i != len(inp)-1 {
			tempMin := findMin(inp[i+1], inp[i])
			switch {
			case tempMin < min:
				secondMin = min
				min = tempMin
			case tempMin > min && tempMin < secondMin && secondMin != min:
				secondMin = tempMin
			}
		}
	}
	fmt.Println(min, secondMin, mp)
	for i := 0; i < len(inp)-1; i++ {
		value := inp[i] + diff
		fmt.Println("value:", value)
		if _, exists := mp[value]; !exists {
			return false
		}
	}
	return true
}
func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
