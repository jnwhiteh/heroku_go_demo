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
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello, world!")
	//pretty.Fprintf(w, "%# v", struct{ X int }{3})
}
