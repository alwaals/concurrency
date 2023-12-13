package main

import (
	"fmt"
	"time"
)

func main() {
	mpI := make(map[string]int)
	mpI["first"]=2
	mpI["second"]=6
	mpI["third"]=7

	mpF := make(map[string]float64)
	mpF["first"]=4.8
	mpF["second"]=6.3
	mpF["third"]=7.8

	res1:=sumOfFloatsInt[string,int](mpI)
	res2:=sumOfFloatsInt[string,float64](mpF)

	fmt.Println(res1,res2)

	// The date that is to be parse
	myDateString := "2018-01-20 04:35"
	fmt.Println("My Starting Date:\t", myDateString)

	// Parse the date string into Go's time object
	// The 1st parameter specifies the format whereas 2nd is the date string
	// myDate, err := time.Parse("2006-01-02 15:04", time.Now().String())
	// if err != nil {
	// 	panic(err)
	// }
	// //fmt.Println(tFormat.String())
	// fmt.Println("Just The Date:\t\t", myDate.Format("2006-01-02"))
	myTime, err := time.Parse("2 Jan 2006 03:04PM", "10 Nov 2016 11:00PM")

	if err != nil {
		panic(err)
	}
	fmt.Println(myTime.String())
	now := time.Now()
	
	fmt.Println(now.Before(myTime))
	
	fmt.Println(now.After(myTime))

}
func sumOfFloatsInt[K comparable, V int | int64 | float64](inp map[K]V) V {
	var sum V
	for _, val := range inp {
		sum += val
	}
	return sum
}
