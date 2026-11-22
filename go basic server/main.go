package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(
		"Access-Control-Allow-Origin",
		"http://localhost:3000",
	)

	w.Header().Set(
		"Access-Control-Allow-Headers",
		"Content-Type, name",
	)

	w.Header().Set(
		"Access-Control-Allow-Methods",
		"GET, OPTIONS",
	)
	
	if r.Method == "OPTIONS"{
		w.WriteHeader(200)
		return
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	value, _ := r.Cookie("name")
	response := fmt.Sprintf("This is home handler. And the value of 'name' cookie is : %s.", value)

	fmt.Fprintf(w, "%s", response)
}

func main() {

	testRouter := http.NewServeMux()

	testRouter.HandleFunc("/", homeHandler)

	serverObj := &http.Server{
		Addr:    ":3000",
		Handler: testRouter,
	}

	fmt.Println("Server is running at http://localhost:3000 ...✅")
	err := serverObj.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Find error to run the server", err)
	}

}
