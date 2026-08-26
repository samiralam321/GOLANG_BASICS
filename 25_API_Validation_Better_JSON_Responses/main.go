// what is validation : means : checking whether the data
// sent by the client is correct before processing it 

{
  "name": "",
  "price": -500,
  "stock": -10
}

//invalid data 

Client sends JSON
       ↓
Backend receives JSON
       ↓
Validate data
       ↓
Valid?
   ↙       ↘
 No        Yes
 ↓          ↓
Error     Continue
Response  Create/Update


//*********** Where should Validation Happen *****************

// createProduct function me : 

Receive Request
      ↓
Decode JSON
      ↓
Validate Data
      ↓
Create Product
      ↓
Send Response

// validation should happen after decoding JSON


// Exmples 

var product Product 

err := json.NewDecoder(r.Body).Decode(&product)

if err != nil {
	//invalid JSON
}

// after this :

validateProduct(product)

//********* Instead of creating a validation code again and again, we create one function 

func validateProduct(prduct Product) string {
	if product.Name == "" {
		return "Product name is required"
	}

	if product.Price <= 0 {
		return "Price must be greater than 0"
	}

	if product.Stock < 0 {
		return "stock cannot be nagative"
	}

	return ""
}

//************ Creating a JSON Error Helper Function *************

func sendError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := map[string] string {
		"error" : message,
	}

	json.NewEncoder(w).Encode(response)
}


// Instead of writing : 

http.Error(w, "Invalid JSON", http.StatusBadRequest)

// we can write : 

sendError(w, "Invalid JSON", http.StatusBadRequest)



//*********** what is interface{} ******************

data interface{} => This function can accept different types of data


//************ Update createProduct **************

func createProduct (w http.ResponseWriter, r *http.Request){
	var product Product 

	err := json.NewEncoder(r.Body).Decode(&product)

	if err != nil {
		sendError ( w, "Invalid JSON", http.StatusBadRequest)
		return 
	}

	errorMessage := validateProduct(product)

	if errorMessage != "" {
		sendError(w, errorMessage, http.StatusBadRequest)
		return 
	}

	product.ID = len(products) + 1
	products = append(products, product)

	sendJSON(w, product, http.StatusCreated)
}

// FLOW 

POST Request
      ↓
Decode JSON
      ↓
Valid JSON?
   ↙      ↘
 No       Yes
 ↓         ↓
Error    Validate Product
             ↓
         Valid data?
          ↙     ↘
        No       Yes
        ↓         ↓
      Error    Create
                 ↓
               201

// Updated updateProduct 

func updateProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		sendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product

	err = json.NewDecoder(r.Body).Decode(&updatedProduct)

	if err != nil {
		sendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	errorMessage := validateProduct(updatedProduct)

	if errorMessage != "" {
		sendError(w, errorMessage, http.StatusBadRequest)
		return
	}

	for i := range products {

		if products[i].ID == id {

			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Stock = updatedProduct.Stock

			sendJSON(w, products[i], http.StatusOK)

			return
		}
	}

	sendError(w, "Product not found", http.StatusNotFound)
}


//***************  Complete Code ***************

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

// validateProduct checks whether product data is valid.
func validateProduct(product Product) string {

	if product.Name == "" {
		return "Product name is required"
	}

	if product.Price <= 0 {
		return "Price must be greater than 0"
	}

	if product.Stock < 0 {
		return "Stock cannot be negative"
	}

	return ""
}

// sendError sends an error response in JSON format.
func sendError(w http.ResponseWriter, message string, statusCode int) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := map[string]string{
		"error": message,
	}

	json.NewEncoder(w).Encode(response)
}

// sendJSON sends a successful JSON response.
func sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(data)
}

// GET /products
func getProducts(w http.ResponseWriter, r *http.Request) {

	sendJSON(w, products, http.StatusOK)
}

// POST /products
func createProduct(w http.ResponseWriter, r *http.Request) {

	var product Product

	// Decode JSON request body.
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		sendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate product data.
	errorMessage := validateProduct(product)

	if errorMessage != "" {
		sendError(w, errorMessage, http.StatusBadRequest)
		return
	}

	// Generate temporary ID.
	product.ID = len(products) + 1

	// Add product to slice.
	products = append(products, product)

	// Send created product.
	sendJSON(w, product, http.StatusCreated)
}

// GET /products/{id}
func getProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		sendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range products {

		if product.ID == id {

			sendJSON(w, product, http.StatusOK)

			return
		}
	}

	sendError(w, "Product not found", http.StatusNotFound)
}

// PUT /products/{id}
func updateProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		sendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product

	// Decode JSON.
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)

	if err != nil {
		sendError(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate data.
	errorMessage := validateProduct(updatedProduct)

	if errorMessage != "" {
		sendError(w, errorMessage, http.StatusBadRequest)
		return
	}

	// Find and update product.
	for i := range products {

		if products[i].ID == id {

			products[i].Name = updatedProduct.Name
			products[i].Price = updatedProduct.Price
			products[i].Stock = updatedProduct.Stock

			sendJSON(w, products[i], http.StatusOK)

			return
		}
	}

	sendError(w, "Product not found", http.StatusNotFound)
}

// DELETE /products/{id}
func deleteProduct(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		sendError(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for i := range products {

		if products[i].ID == id {

			products = append(products[:i], products[i+1:]...)

			response := map[string]string{
				"message": "Product deleted successfully",
			}

			sendJSON(w, response, http.StatusOK)

			return
		}
	}

	sendError(w, "Product not found", http.StatusNotFound)
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


// Now the API is much better : 

Request
   ↓
Decode JSON
   ↓
Validate
   ↓
Business Logic
   ↓
JSON Response












































