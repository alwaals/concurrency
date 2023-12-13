package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	url := "https://raw.githubusercontent.com/sd2995/GoLang/main/HelloWorldGo/main.go"

	respBody, err := http.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer respBody.Body.Close()
	resp, err := io.ReadAll(respBody.Body)
	if err != nil {
		log.Fatal("Error reading :", err)
	}
	f, err := os.Create("httpResponse")
	if err != nil {
		log.Fatal("Error Create :", err)
	}
	f.Write(resp)

	httpReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
	respR, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		log.Fatal("Error DefaultClient.Do:", err)
	}
	defer respR.Body.Close()

	var b bytes.Buffer
	io.Copy(&b, respR.Body)

	file, err := os.Create("httpNewRequest")
	if err != nil {
		log.Fatal("Error Create:", err)
	}
	file.Write(b.Bytes())
}
