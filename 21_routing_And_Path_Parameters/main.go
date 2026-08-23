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

func getProducts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(products)
}

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

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /products", getProducts)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("GET /products/{id}", getProduct)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", mux)
}


//************ Understand the Router ***************

                   HTTP Request
                        │
                        ↓
                     Router
                        │
          ┌─────────────┼─────────────┐
          ↓             ↓             ↓
    GET /products  POST /products  GET /products/{id}
          ↓             ↓             ↓
    getProducts    createProduct   getProduct


// The router s job is basically:
// "Which handler should receive this request?"

// This becomes increasingly important as your application grows.


The main concepts covered:


1. HTTP routing
2. http.NewServeMux()
3. GET /products
4. POST /products
5. GET /products/{id}
6. Path parameters
7. r.PathValue()
8. strconv.Atoi()
9. Finding a product by ID
10. 400 Bad Request
11. 404 Not Found
12. 201 Created
13. JSON response

## Complete `main.go`


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Product represents one product.
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

// Temporary in-memory data.
// Later we will replace this with PostgreSQL.
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
// Returns all products.
func getProducts(w http.ResponseWriter, r *http.Request) {

	// Tell the client that we are returning JSON.
	w.Header().Set("Content-Type", "application/json")

	// Convert products slice into JSON
	// and send it as the response.
	json.NewEncoder(w).Encode(products)
}

// POST /products
// Creates a new product.
func createProduct(w http.ResponseWriter, r *http.Request) {

	var product Product

	// Read JSON from request body
	// and convert it into Product struct.
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate a temporary ID.
	// We will use a database for proper IDs later.
	product.ID = len(products) + 1

	// Add the new product to the slice.
	products = append(products, product)

	// Response will be JSON.
	w.Header().Set("Content-Type", "application/json")

	// 201 means resource was successfully created.
	w.WriteHeader(http.StatusCreated)

	// Send created product as JSON.
	json.NewEncoder(w).Encode(product)
}

// GET /products/{id}
// Returns one product by ID.
func getProduct(w http.ResponseWriter, r *http.Request) {

	// Get the {id} value from the URL.
	//
	// Example:
	// /products/2
	//
	// r.PathValue("id") → "2"
	idString := r.PathValue("id")

	// Convert string ID into integer.
	//
	// "2" → 2
	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Search for the product.
	for _, product := range products {

		if product.ID == id {

			// Response is JSON.
			w.Header().Set("Content-Type", "application/json")

			// Send the product.
			json.NewEncoder(w).Encode(product)

			// Stop the function because
			// we already found the product.
			return
		}
	}

	// If the loop finishes, product was not found.
	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {

	// Create a new HTTP router.
	mux := http.NewServeMux()

	// Register routes.

	// GET /products
	mux.HandleFunc("GET /products", getProducts)

	// POST /products
	mux.HandleFunc("POST /products", createProduct)

	// GET /products/{id}
	mux.HandleFunc("GET /products/{id}", getProduct)

	fmt.Println("Server running on http://localhost:8080")

	// Start the server.
	http.ListenAndServe(":8080", mux)
}


## Understand the important part

Your router is:

mux := http.NewServeMux()

mux.HandleFunc("GET /products", getProducts)
mux.HandleFunc("POST /products", createProduct)
mux.HandleFunc("GET /products/{id}", getProduct)


Think of it as a traffic controller:


                    Request
                       ↓
                     Router
                       |
          ┌────────────┼────────────┐
          ↓            ↓            ↓
    GET /products  POST /products  GET /products/{id}
          ↓            ↓            ↓
    getProducts   createProduct   getProduct


## Path Parameter

For:

GET /products/10

this:

r.PathValue("id")

returns:

"10"

Its a string.

So we convert it:

id, err := strconv.Atoi(idString)

Now:

"10" → 10

Then we search:

for _, product := range products {

	if product.ID == id {
		// found
	}
}

## Status codes used

200 OK
    Successful GET

201 Created
    Product successfully created

400 Bad Request
    Invalid input / invalid ID / invalid JSON

404 Not Found
    Product doesn t exist


## Endpoints you now have


GET    /products
       → Get all products

POST   /products
       → Create product

GET    /products/1
       → Get product 1

GET    /products/2
       → Get product 2

GET    /products/999
       → 404 Product not found

### Postman testing

For creating a product:


POST http://localhost:8080/products

Body → raw → JSON:

```json
{
  "name": "Monitor",
  "price": 15000,
  "stock": 8
}
```

For getting all:


GET http://localhost:8080/products


For getting one:

GET http://localhost:8080/products/2

---

### The 3 lines you should remember

r.PathValue("id")

Gets the dynamic value from the URL.

strconv.Atoi(idString)

Converts the URL string into an integer.

mux.HandleFunc("GET /products/{id}", getProduct)

Creates a route with a dynamic path parameter.


------------------------------------------------------


"GET /products/{id}"
doesnt literally mean only:

/products/{id}

It means:
/products/1
/products/2
/products/50
/products/abc

The actual value can be retrieved with:
r.PathValue("id")

For:
/products/50

you get:
"50"

For:
/products/abc

you get:
"abc"

Then strconv.Atoi() determines whether it s a valid integer.