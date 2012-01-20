package main

import (
	"fmt"
	"net/http"
	//	"github.com/kr/pretty.go"
	"log"
	"os"
)

func main() {
	http.Handle("/", http.HandlerFunc(hello))
	port := os.Getenv("PORT")
	log.Printf("Starting server on :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
	//pretty.Fprintf(w, "%# v", struct{ X int }{3})
}
