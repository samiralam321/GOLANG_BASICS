// u know how to receive request 

Client
   ↓
HTTP Request
   ↓
Go Handler

// now we want the client to actually send data to our backend

// for examples - a fronted wants to create a user : 

// POST /users 

and send 

{
	"name" : "Samir",
	"email" : "sa8103339@gmail.com",
	"age" : 20
}

so our backend needs to 


Receive JSON
     ↓
Read Body
     ↓
Convert JSON → Go struct
     ↓
Process data
     ↓
Send response



GET /users    : Server, give me the users 
POST /users   : Servers, here is some data, Create a user with it


id := r.URL.Query().Get("id")


POST Usually Sends Data in the Body 

POST /users 

body : 

{
  "name": "Samir",
  "email": "samir@gmail.com",
  "age": 20
}

// noe the data is not in the URL it is in the body r.Body

r.URL    → URL information
r.Header → Header information
r.Body   → Body data


package main
import (
	"fmt"
	"io"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintln(w, "Could not read body")
		return
	}

	fmt.Fprintln(w, string(body))
}
func main(){
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServer(":8080", nil)
}


// body, err := io.ReadAll(r.Body)

// which gives u raw bytes 

// we do not want to manually work with JSON

// we want : 

JSON
 ↓
Go struct   that is where encoding/json comes in 

// struct 

type User struct {
	Name string
	Email string
	Age int
}



//********************** JSON -> GO Struct ********************

// use => json.Decoder

decoder := json.NewDecoder(r.Body)
err := decoder.Decode(&user)


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string
	Email string
	Age int
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return 
	}
	fmt.Fprintln(w, "Name:", user.Name)
	fmt.Fprintln(w, "Email:", user.Email)
	fmt.Fprintln(w, "Age:", user.Age)
}


func main(){
	http.HandleFunc("/users", userHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


// Send 

{
  "name": "Samir",
  "email": "samir@gmail.com",
  "age": 20
}

// Response

Name: Samir
Email: samir@gmail.com
Age: 20


// so we have converted JSON into Go struct 



//***************************** JSON Tags **************************

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}


// now you can explicitly saying : 

JSON "name"  → Go Name
JSON "email" → Go Email
JSON "age"   → Go Age

// this is the standard approach u will see in GO APIs



package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Age int      `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Name:", user.Name)
	fmt.Fprintln(w, "Email:", user.Email)
	fmt.Fprintln(w, "Age:", user.Age)
}

func main(){
	http.HandleFunc("/users", userHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServer(":8080", nil)
}



//*************************** Add Method Checking ******************

//our /users endpoint should not accept every method 

// we want POST /users

if r.Method != http.MethodPost {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}


//


func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Name:", user.Name)
	fmt.Fprintln(w, "Email:", user.Email)
	fmt.Fprintln(w, "Age:", user.Age)
}



/////********************* 

Currently we are  returning:
Name: Samir
Email: ...
Age: 20


But APIs usually return JSON.

// we want :

{
  "message": "User created successfully",
  "user": {
    "name": "Samir",
    "email": "samir@gmail.com",
    "age": 20
  }
}


// so we need to convert : Go struct -> JSON 

// This is called Encoding / marshalling 


//// Complete POST API ****************


package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func main() {

	http.HandleFunc("/users", userHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}


//********************** Decoder vs Encoder ***************

// Decoder

json.NewDecoder(r.Body).Decode(&user)

// Direction : JSON -> GO

// Encoder

json.NewEncoder(w).Encode(user)

// Direction : GO -> JSON


Incoming:

JSON
 ↓
Decoder
 ↓
Go struct


Outgoing:

Go struct
 ↓
Encoder
 ↓
JSON


//************************* Marshal Vs Encoder *********************

json.Marshal(user)

this converts : 

Go -> []byte JSON


while json.NewEncoder(w).Encodr(user)

writes JSON directly to the response

For HTTP handlers, this is very convnient

json.NewEncoder(w).Encode(user)



//******************* IMP Backend Pattern ******************

if r.Method != http.MethodPost {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	return
}

var user User

err := json.NewDecoder(r.Body).Decode(&user)

if err != nil {
	http.Error(w, "Invalid JSON", http.StatusBadRequest)
	return
}

it is basically 

1. Check request
2. Read request
3. Validate request
4. Process request
5. Send response




