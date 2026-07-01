package main

import (
	"fmt"
	"net/http"
)

func getHome(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This is Home Page")
}

func getAbout(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "This is About Page")
}

func main() {

	// router objects
	mux := http.NewServeMux()

	//req-res handle
	mux.HandleFunc("/", getHome)
	mux.HandleFunc("/about", getAbout)

	fmt.Println("Server is running @ http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("Find error : ", err)
	}

}
