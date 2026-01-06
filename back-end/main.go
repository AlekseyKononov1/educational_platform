package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		randomNumber := rand.Intn(11) // 0–10
		fmt.Fprintf(w, "Hello World %d", randomNumber)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
