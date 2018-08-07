package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	cur, _ := os.Getwd()
	fmt.Println("Current directory: ", cur)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello %s", request.URL.Path)
}
