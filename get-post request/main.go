package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// declaration part
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

// all controllers
func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is HOME page")
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Send only GET requests", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

func createProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Please send POST request", 400)
		return
	}

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Send a valid json", 400)
		return
	}

	newProduct.Id = len(productList) + 1
	w.WriteHeader(201)

	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)

}

func main() {

	// routers
	mux := http.NewServeMux()

	// req-res cycle
	mux.HandleFunc("/", getHome)
	mux.HandleFunc("/products", getProducts)
	mux.HandleFunc("/create-product", createProducts)

	fmt.Println("Server is running at http://localhost:3000")

	// start the server
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Find error to start server : ", err)
	}
}

func init() {
	product1 := Product{
		Id:          1,
		Name:        "Laptop",
		Price:       75000,
		Description: "High performance gaming laptop",
		ImgURL:      "https://example.com/images/laptop.jpg",
	}

	product2 := Product{
		Id:          2,
		Name:        "Smartphone",
		Price:       28000,
		Description: "Android smartphone with display",
		ImgURL:      "https://example.com/images/phone.jpg",
	}

	product3 := Product{
		Id:          3,
		Name:        "Wireless Mouse",
		Price:       1200,
		Description: "Ergonomic wireless mouse",
		ImgURL:      "https://example.com/images/mouse.jpg",
	}

	product4 := Product{
		Id:          4,
		Name:        "Mechanical Keyboard",
		Price:       4500,
		Description: "RGB mechanical keyboard",
		ImgURL:      "https://example.com/images/keyboard.jpg",
	}

	product5 := Product{
		Id:          5,
		Name:        "Monitor",
		Price:       18000,
		Description: "24-inch Full HD IPS monitor",
		ImgURL:      "https://example.com/images/monitor.jpg",
	}

	product6 := Product{
		Id:          6,
		Name:        "Headphone",
		Price:       3500,
		Description: "Noise cancelling wireless headphone",
		ImgURL:      "https://example.com/images/headphone.jpg",
	}

	productList = append(productList, product1)
	productList = append(productList, product2)
	productList = append(productList, product3)
	productList = append(productList, product4)
	productList = append(productList, product5)
	productList = append(productList, product6)

}
