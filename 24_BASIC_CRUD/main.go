package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products = []Product{
	{
		ID:    1,
		Name:  "Laptop",
		Price: 50000,
		Stock: 10,
	},
	{
		ID:    2,
		Name:  "Mouse",
		Price: 1000,
		Stock: 25,
	},
	{
		ID:    3,
		Name:  "Keyboard",
		Price: 2000,
		Stock: 15,
	},
}

// GET /products
func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

// POST /products
func createProduct(w http.ResponseWriter, r *http.Request) {

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	product.ID = len(products) + 1

	products = append(products, product)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(product)
}

// GET /products/{id}
func getProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {

		if product.ID == id {

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(product)

			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// PUT /products/{id}
func updateProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product

	err = json.NewDecoder(r.Body).Decode(&updatedProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range products {

		if products[i].ID == id {

			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Stock = updatedProduct.Stock

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(products[i])

			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// DELETE /products/{id}
func deleteProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i := range products {

		if products[i].ID == id {

			// Remove the product
			products = append(products[:i], products[i+1:]...)

			w.Header().Set("Content-Type", "application/json")

			response := map[string]string{
				"message": "Product deleted successfully",
			}

			json.NewEncoder(w).Encode(response)

			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("GET /products/{id}", getProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)
	mux.HandleFunc("DELETE /products/{id}", deleteProduct)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", mux)
}
