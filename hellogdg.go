package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello GDG called...")
	msg := os.Getenv("SIMPLE_MSG")
	if msg == "" {
		msg = ":( SIMPLE_MSG variable not defined"
	}
	fmt.Fprintf(w, "%s", msg)
}

func main() {
	flag.Parse()
	log.Print("Simple app server started...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
