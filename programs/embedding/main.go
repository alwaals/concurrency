package main

import (
	"embed"
	"fmt"
)

//go:embed Files/*
var folder embed.FS
func main(){
	file1,_:=folder.ReadFile("Files/fruits.txt")
	fmt.Println(string(file1))
	file2,_:=folder.ReadFile("Files/Vegetables.txt")
	fmt.Println(string(file2))
}