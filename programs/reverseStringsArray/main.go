package main

import (
	"fmt"
	"strings"
)


func main(){
	s := " the sky is blue "
	strs:=strings.Fields(s)
	for i,j:=0,len(strs)-1;i<len(strs)/2;i,j=i+1,j-1{
		strs[i],strs[j]=strs[j],strs[i]
	}
	fmt.Println(strs)
}