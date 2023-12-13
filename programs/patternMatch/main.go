package main

import (
	"fmt"
	"os"
	"strings"
)
 
func main() {
	s := []string{"This", "is", "Scaler", "Academy."}
 
    v := strings.Join(s, " Swetha \n")
     
    fmt.Printf("%v",v)  // printing the output.
    // var name string = "Scaler"
    // var wordToSearch string = "ler"
    // res := checkIfExists(name, wordToSearch)
    // if res {
    //     fmt.Println(wordToSearch, "is present inside the string", name)
    // } else {
    //     fmt.Println(wordToSearch, "is not present inside the string", name)
    // }
	bytes,err:=fmt.Fprintf(os.Stdout, "%s",v)
	fmt.Println(bytes,err)
	s1 := fmt.Sprintf("%d", 33)
	fmt.Println(s1)
	
}