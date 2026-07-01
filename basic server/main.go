package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Hello page")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is About page")
	})

	fmt.Println("Server is running at http://localhost:3000 ✅")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Find error : ", err)
	}
}
