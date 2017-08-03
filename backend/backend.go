package main

import (
	"encoding/json"
	"net/http"
)

var counter int

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		var data = struct {
			Response string
		}{
			Response: "Hello World!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}
