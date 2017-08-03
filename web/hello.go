package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var counter int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})
	http.HandleFunc("/chain", func(w http.ResponseWriter, r *http.Request) {
		response, err := http.Get("http://backend/api")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer response.Body.Close()

		var data struct {
			Response string
		}
		err = json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprintln(w, data.Response)
	})
	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		counter++
		fmt.Fprintln(w, counter)
	})
	http.HandleFunc("/crash", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stderr, "Crash!")
		os.Exit(1)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
