GET    /products        → Get all products
POST   /products        → Create product
GET    /products/{id}   → Get one product



Client
  ↓
PUT /products/2
  ↓
Get ID from URL
  ↓
Read JSON body
  ↓
Find product
  ↓
Update product
  ↓
Return updated product


func updateProduct(w http.ResponseWriter, r *http.Request)


// Inside this function : 

Step 1 → Get ID from URL
Step 2 → Convert string ID to int
Step 3 → Read JSON body
Step 4 → Find product
Step 5 → Update product
Step 6 → Return updated product


// Complete  updateFunction Function 


func updateProduct(w http.ResponseWriter, r *http.Request) {

	// Get ID from URL
	idString := r.PathValue("id")

	// Convert string to integer
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Store new data
	var updatedProduct Product

	// Decode JSON request body
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Find the product
	for i := range products {

		if products[i].ID == id {

			// Update fields
			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Stock = updatedProduct.Stock

			// Send JSON response
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(products[i])

			return
		}
	}

	// Product not found
	http.Error(w, "Product not found", http.StatusNotFound)
}



// complete main.go

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

	// Get ID from URL
	idString := r.PathValue("id")

	// Convert string to int
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Store new product data
	var updatedProduct Product

	// Read JSON body
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Find and update product
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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("GET /products/{id}", getProduct)
	mux.HandleFunc("PUT /products/{id}", updateProduct)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", mux)
}


// Compelte Flow of the Program 


PUT /products/2
        │
        ↓
r.PathValue("id")
        │
        ↓
"2"
        │
        ↓
strconv.Atoi()
        │
        ↓
2
        │
        ↓
Read JSON body
        │
        ↓
updatedProduct
        │
        ↓
Find products[i].ID == 2
        │
        ↓
Update products[i]
        │
        ↓
Return updated product


//**************** IMP Note ****************

// for Searching, Reading and Checking Usee : 

for _, product := range products {

}


// for Updating products use : 
// Use for updating, Changing and Deleting

for i : range products {

}











