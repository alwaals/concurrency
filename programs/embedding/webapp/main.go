package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

//go:embed webpages
var web embed.FS

func main() {
	fmt.Println("Starting service!")
	mux := http.NewServeMux()
	mux.Handle("/", homeHandler())
	http.ListenAndServe(":8000", mux)
}
func homeHandler() http.Handler {
	file := fs.FS(web)
	html, err := fs.Sub(file, "webpages")
	if err != nil {
		log.Fatal("Page not found:", err)
	}
	return http.FileServer(http.FS(html))
}
