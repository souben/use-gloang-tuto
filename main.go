package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, R *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to myawesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3600", nil)
}
