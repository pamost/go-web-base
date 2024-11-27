package main

import (
	"fmt"
	"net/http"
)


func handlerFunc (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
	
}

func main(){
	http.HandleFunc("/", handlerFunc)
	fmt.Println("starting server on port 8080...")
	http.ListenAndServe(":8080", nil)
}