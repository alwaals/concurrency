package main

import (
	"log"
	"os"
)

func main(){
	err:=os.Mkdir("mkdir",0750)
	if err!=nil{
		if !os.IsExist(err){
			log.Fatal("directory already exists")
		}
	}
	_=os.Chdir("mkdir")
	file,fErr:=os.Create("text")
	if fErr!=nil{
		log.Fatal(fErr.Error())
	}
	file.Write([]byte("This is Golang program!"))
}