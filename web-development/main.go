package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey there")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is response")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/home", homeHandler)
	http.ListenAndServe(":8000", nil)
}
