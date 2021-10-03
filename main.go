package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type": "text/html")
	fmt.Fprint(w, "<h1>Welcome to bad site!</h1>")
	if r.URL.Path == "/"{
		fmt.Fprint(w, "<h1>Welcome to my awesome sitel</h1>")
	}else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "get in touch ...")
	}else{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Hello my name is ...</h1>")
	}
}



func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3600", nil)
	
}
