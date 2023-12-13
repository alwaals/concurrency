package main

import "fmt"

func main(){
	inp:=[]string{"a","d","g","a","g","a","g","f","g"}
	gCount:=0
	agCount:=0
	for i:=len(inp)-1;i>=0;i--{
		if inp[i]=="g"{
			gCount+=1
		}
		if inp[i]=="a"{
			agCount+=gCount
		}
	}
	fmt.Println(agCount)
}