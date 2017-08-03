package main

import (
	"fmt"
	"net/http"
	"os"
)

var counter int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintln(w, counter)
	})
	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "Crash!\n")
		os.Exit(1)
	})

	http.ListenAndServe(":8080", nil)
}
