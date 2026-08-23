//***************** what is JSON *************************

/*

JSON => JavaScript Object Notation 

it is a common format of exchanging data between applications 


{
  "name": "Samir",
  "age": 20,
  "email": "samir@gmail.com"
}

A frontend can send this to your Go backend.
A Go backend can send it to a frontend.
A mobile application can send it to your Go backend.
Another backend can send it to your Go backend.
That's why JSON is so important.


JSON and Go

Frontend
   ↓
JSON
   ↓
Go Backend
   ↓
Go Struct

and when sending a response

Go Struct
   ↓
  JSON
   ↓
Frontend

*/



// 4 Important JSON functions 

json.Marshall()
json.Unmarshall()
json.NewEncoder()
json.NewDecoder()

-----------------------------------------

Marshal
Go → JSON

Unmarshal
JSON → Go

Encoder
Go → JSON stream

Decoder
JSON stream → Go


//********************* Json.Marshal() *****************

// GO -> Marshal

type User struct {
   Name string
   Email string
   Age int
}

//create : 

user := User {
   Name : "Samir",
   Email : "samir@gmail.com",
   Age : 20,
}


data, err := json.Marshal(user) // converts the Go struct into JSON bytes


Go struct
   ↓
Marshal
   ↓
JSON

// Examples : 

package main
import (
   "encodinf/json"
   "fmt"
)


type User struct {
   Name string
   Email string
   Age int
}


func main(){
   user := User {
      Name : "Samir",
      Email : "samir@gmail.com",
      Age : 20,
   }

   data, err := json.Marshal(user)

   if err != nil {
      fmt.Println("Error : ", err)
      return
   }
   fmt.Println(string(data))
}


// Output : {"Name":"Samir","Email":"samir@gmail.com","Age":20}

//Note that : the JSON field names are : 

Name
Email
Age

// But usually API wants : 

name
email
age

// that is where JSON tag comes in 

type User struct {
   Name string `json:"name"`
   Email string `json"email"`
   Age int `json:"age"`
}

data, err := json.Marshal(user)


// so now it produces : 

{
  "name": "Samir",
  "email": "samir@gmail.com",
  "age": 20
}

// this is much more common in APIs

//****************** what does Marshal() reutrn ?? *************************
data, err := json.Marshal(user)

// return []bytes

// not a normal string
// so we use : 

fmt.Println(string(data))

// to converts the bytes into a string for printing 

[]byte
 ↓
string()
 ↓
string


//************************* json.Unmarshal() **********************

// now suppose we have JSON

{
  "name": "Samir",
  "email": "samir@gmail.com",
  "age": 20
}

// and we want to convert it into User 

// so use json.Unmarshal()

JSON
 ↓
Unmarshal
 ↓
Go struct


// Simple Unmarshal Examples 

package main
import (
   "encoding/json"
   "fmt"
)

type User struct {
   Name string `json:"name`
   Email string `json:"email"`
   Age int `json:"age"`
}

func main(){
   data := []byte(`{
      "name" : "samir",
      "email" : "sa@gmail.com",
      "age" : 20   
   }`)

   var user User 

   err := json.Unmarshal(data, &user)
   if err != nil {
      fmt.Println("Error : ", err)
      return 
   }
   fmt.Println("Name : ", user.Name)
   fmt.Println("email : ", user.email)
   fmt.Println("Age : ", user.Age)
}

// Output : 

Name: Samir
Email: samir@gmail.com
Age: 20


// why &user Again ?? 

/// json.Unmarshal(data, &user)

// casue Unmarshal needs to modify the user variable 

so : 

&user
 ↓
address of user
 ↓
Unmarshal can fill the struct

//******************* Marshal Vs Unmarshal *****************
// Marshal : Go -> JSON
//UNmarshal : JSON -> GO



//****************** Encoder Vs Decoder ******************************

json.NewDecoder(r.Body).Decode(&product)

json.NewEncoder(r.Body).Encode(product)


Marshal / Unmarshal
→ Work with data in memory

Encoder / Decoder
→ Work with streams such as HTTP request/response bodies


// For HTTP APIs, you will commonly use : 

json.NewDecoder(r.Body) for incoming JSON

json.NewDecoder(w) for outgoing JSON


//************** The Complete JSON Flow in an API ******************

// imagine : POST /users

Client Sends : 

{
   "name" : "samir",
   "age" : 20,
}

// GO receives : 

r.Body
   ↓
Decoder
   ↓
User struct

// then your backend process the user : 

// finally : 

User struct
   ↓
Encoder
   ↓
JSON
   ↓
HTTP Response


//********************** omitempty ************************

type User struct {
   Name string `json:"name"`
   Email string `json:"email"`
   Age int `json:"age"`
}


user := User {
   Name : "Samir",
}

// Marshal it , you will get ; json

{
   "name" : "samir",
   "email" : "",
   "age" : 0
}

// sometime we do not want the empty values in the JSON, so we can use : 

omitempty

type User struct {
   Name string `json:"name"`
   Email string `json:"email,omitempty"`
   Age int      `json:"age,omitempty"`
}


// now you will get 
// Go
user := User {
   Name : "Samir",
}


//JSON

{
   name : "Samir"
}

// casue 
Email = ""
Age = 0

these filed are omitted 

//******************** Why is omitempty Useful **************************

// suppose your API returns : 
// JSON

{
  "name": "Samir",
  "email": "",
  "phone": "",
  "age": 0,
  "address": ""
}

// this is often unncecessary
// so by using json : "email,omitempty"
// we can produce more cleaner repsonse



//******************* Nested JSON *********************

// Real APIs don't always return simple objects 
// JSON
{
   "name" : "Samir",
   "address" : {
      "city" : "Kapurthala",
      "country" : "INDIA"
   }
}

//GO

type Addresss struct {
   City string `json:"city"`
   Country string `json:"country"`
}

type User struct {
   Name string `json:"name`
   Email string `json:"email"`
   Address Address `json:"address"`
}

// now 

//GO

user := User {
   Name : "Samir",
   Email : "sa@gmail.com",

   Address : Address {
      City : "Kapurthala",
      Country : "India",
   },
}

/// Mrshal it : Go -> JSON

data, _ := json.Marshal(user)
fmt.Println(string(data))

//JSON
{
   "name" : "Samir",
   "email" : "samir@gmail.com",
   "address" : {
      "city" : "Kapurthala",
      "country" : "India"
   }
}


//*********************** Slices in JSON *****************************

type Product struct {
   Name string `json:"name"`
   Price int `json:"price"`
}


//Create : 

//Go
products := []Product {
   {
      Name : "Laptop",
      Price : 500000,
   },

   {
      Name : "Mouse",
      Price: 10000,
   },
}

// Now marshal it 

data, _ := json.Marshal(products)
fmt.Println(string(data))

//JSON

[
   {
      "name" : "Laptop",
      "price" : 5000
   },

   {
      "name" : "Mouse",
      "price" : 10000
   }
]


// this is how APIs return lists of data 



//******************* Real APIs Response ****************************

GET /products

// your backend might returns : 
// JSON

[
   {
      "name" : "Laptop",
      "price" : 50000
   },

   {
      "name" : "Mouse",
      "price" : 1000
   }
]


//***************** Response Object ****************

// often real APIs wrap the data : 

// so instead of : 

[
   {
      "name" : "Laptop",
      "price" : 50000
   }
]


// you might return : 

{
   "message" : "Products fetched succesfully",

   "products" : [
      {
         "name" : "Laptop",
         "price" : 50000
      }
   ]
}

// we can create : 
//Go

type ProductReponse struct {
   Message string `json:"message"`
   Products []Product `json:"products"`
}


// then
// Go


response := ProductResponse {
   Message : "Products fetched Successfully",
   Products : products,
}

// and ecode it 

package main
import (
   "encoding/json"
   "fmt"
)

type Product struct {
   Name string `json:"name"`
   Price int `json:"price`
}

type ProductResponse struct {
   Message string `json:"message"`
   Products []Product `json:"products"`

}

func main(){
   products := []Product {
      {
         Name : "laptop",
         Price : 50000,
      },

      {
         Name : "Mouse",
         Price : 10000,
      },
   }

   response := ProductResponse{
      Message : "Products fectched Successfully",
      Products : products,
   }

   data,err := json.Marshal(response)

   if err != nil {
      fmt.Println("Errors : ", err)
      return 
   }

   fmt.Println(string(data))
}


//**************** Now conect this to your HTTP Server ********************

// make a real ednpoint

GET /products 

package main
import (
   "encoding/json"
   "fmt"
   "net/http"
)

type Product struct {
   Name string `json:"name"`
   Price int `json:"price"`
}

func produchHandler(w http.ResponseWriter, r *http.Request) {
   products := []Product {
      {
         Name : "Laptop",
         Price : 500000,
      },

      {
         Name : "Mouse",
         Price : 10000,
      },
   }

   w.Header().Set("Content-Type", "application/json")
   json.NewEncoder(w).Encode(products)
}

func main(){
   http.HandleFunc("/products", productHandler)

   fmt.Println("Server is running on http://locahost:8080")

   http.ListenAndServe(":8080", nil)
}

// ok so now visitt : http://localhost:8080/products

// you will receive : 

//JSON

[
  {
    "name": "Laptop",
    "price": 50000
  },
  {
    "name": "Mouse",
    "price": 1000
  }
]


//************** Ok now let's Understand the Above Code in Detailed  (Recap) ******************************

// your handler is : 

func productsHandler(w http.ResponseWriter, r *http.Request)

// so there are two imp things are here : 

r -> Request coming FROM the client
w -> Reponse going BACK TO the client

// think like your broswer requesting : 

GET /products

// the flow is like : 

Browser
   |
   | HTTP Request
   ↓
Your Go Server
   |
   | HTTP Response
   ↓
Browser

// Go gives your handler two objects to deal with this : 

r = what did the client aks me ?
w = what should i send back to the client ? 


////////////////////// What ir r ? 
// so it contain the information about the incoming request : 

like : 

r.Method might be GET 

and 

r.URL.Path might be /products 

you can also get headers : 

r.Header

// these are headers that the CLIENT sent to your server 

// so : r.Header is about the incoming request

//*********** What is w ? *************

w is http.ReponseWriter

// It represents the HTTP response that your server is going to send back.

// for examples , your server might send : 

status 200 OK

Content-Type: application/json

[
   {
      "name" : "Laptop",
      "price" : 50000
   }
]

// we use w to build this reponse got it ???

r → Request → coming IN

w → Response → going OUT

// ok so now let's understand this line :

w.Header().Set("Content-Type", "application/json")


w : means i m working with the response
w.Header : 

















