package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func main() {
	fmt.Println("Starting service!")
	data := `{"lang":{"key":"ABC","value":4000},"program":{"key":"hjwey","value":7000}}`
	var mp map[string]interface{}
	json.Unmarshal([]byte(data), &mp)
	fmt.Println(mp)
	langs := mp["lang"].(map[string]interface{})
	for k, v := range langs {
		fmt.Println(k, ":", v)
	}
}
