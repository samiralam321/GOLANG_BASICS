/*


What is Backend ??

suppsoe you open Instagram 

you see : username, password, post,s like, comments, followers

where does all this data come from ?
your phone / browser communicate with a backend server

Your Browser
     ↓
   Internet
     ↓
Instagram Backend
     ↓
Database

this backedn is responsible for things like : 

Login
Signup
Get posts
Create posts
Like posts
Store users
Fetch data
Validate requests


so our goal is to build this type of backedn using Go



//****************** What is Client ********************

A client is something that sends a request to your backend

CLients : asks for something

For examples : 

Browser -> "I want user 10"


******************* what is Server ***********************

A server is a program that waits for requests and responds to them 

Client
   ↓
HTTP Request
   ↓
Go Server
   ↓
Process Request
   ↓
HTTP Response
   ↓
Client

So

Client -> request
Server -> Response

this is first thing you should remember


********************* What i HTTP ? ************************

// HyperText Transfer Protocol

HTTPS is a set of rules that allows clinet and server to communicate


Browser → HTTP Request → Server
Browser ← HTTP Response ← Server

When you visit : https://google.com

your browser communicate with Google's servers using HTTP realted protocol


**************** what is HTTP Request ******************

a ewquest is basicallt 

Hey server i want Something 

for examples : GET /users/10

it means : give me user 10 

******************* What is HTTP Response ? *********************

the server responds 

for examples : 200 OK

{
    "name": "Samir",
    "age": 20
}



**************** The Most Important Flow ******************

CLIENT
  |
  | HTTP REQUEST
  ↓
SERVER
  |
  | process request
  ↓
DATABASE / LOGIC
  |
  ↓
SERVER
  |
  | HTTP RESPONSE
  ↓
CLIENT

this is the foundation of backend development


******************* What is URL ? **************************

https://example.com/users/10


https://example.com/users/10
  ↑          ↑          ↑
protocol    host       path


https
Protocol.

example.com
Server/domain.

/users/10
Path.



********************* what is localhost ********************

suppose you write a Go server on your computer : 
you can access it using localhost

localhost basically means :

This computer 

so https://localhost:8000

means: connect to an HTTP server running on this computer using port 8000



*********************** what are the HTTP Methods *********************

HTTP request usually have a method 

The important ones are : 

GET
POST
PUT
PATCH
DELETE


GET : Give me data : GET /users
POST : Create something : POST /users
PUT : Repalce/update something : PUT /users/10
PATCH : Partially update something : PATCH /users/10
DELETE : Delete something : DELETE /users/10




******************* what is Endpoint ? **********************

An endpoint is basically a specific API location that clients can interact with

GET /users   could means : Get all users 
and GET /users/10 could mean Get user 10 

Another : 

POST /users could mean Create a user.

you will build these later


************************* Handler *****************************

A handler is the code that handles an incoming HTTP request

Request
   ↓
Handler
   ↓
Process
   ↓
Response

for examples : 


GET /users
     ↓
UserHandler
     ↓
Get users
     ↓
Return JSON


******************* First GO HTTP Server ******************

Go has a standard library package called : 

net/http

this package provides tools for buildinf HTTP clients and servers 

*/


package main

import (
	"fmt" // for printing 
	"net/http" // for HTTP functionality
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		fmt.Fprintln(w, "Hello from Go Backend!")

	})

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}


// Breakdown of the above code : 

http.HandleFunc
      ↓
Register a function to handle requests

// we are saying : When someone requests /, run this function 

// What is / ?   is the path. 

func(w http.ResponseWriter, r *http.Request) {   => Handler Function 

	it receives two important things : 

	w and r

r = request  : contains the inforamtion about the incoming HTTP request
so r basically represnt : what did the clinet send to my server ? 


// w = ResponseWriter 

w http.ResponseWriter

w is what we use to send a reponse back to the client ? 

so 

r -> request coming IN
w -> Response going OUT



r = request
w = response 


// Sending a Response 

fmt.Fprintln(w, "Hello from Go Backend!!")



----------------------------------------------------

http.ListenAndServe(":8080", nil)

// it start the server and wait for thr requests 


// this starts the HTTP Server 

:8080 -> Port 8080

so Go listen for request on : localhost:8080

//

Start server
    ↓
Wait...
    ↓
Request arrives
    ↓
Handle request
    ↓
Wait...
    ↓
Another request
    ↓
Handle request

this is what a server does 



//// Examples : 


package main
import ("fmt" "net/http")

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to my go backend !!")
}


func main(){
	http.HandleFunc("/", homeHandler)
	fmt.Println("server running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}





 ////////////////////////////

package main
import (
	"fmt"
	"nrt/http"
)

func homeHandler(w https.http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Go Backend!!")
}

func main(){
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/users", userHandler)

	fmt.Println("Server running on http://localhost:8080")

	http.ListenAndServer(":8080", nil)
}

// now visit : http://localhost:8080/
// Reponse : Welcome to my Go Backend !!

and 

http://localhost:8080/users

Reponse : Users endpoint


---------------------------------------------------


                  Go Server
                     |
          ┌──────────┴──────────┐
          ↓                     ↓
         "/"                 "/users"
          ↓                     ↓
    homeHandler()        usersHandler()
          ↓                     ↓
       Response              Response


--------------------------------------------------



// Recap 

// most important things :

r -> incoming request
w -> outgoing reponse

and 

http.HandleFunc("/users", userHandler)

it means : when a request comes to /users , call usersHandler

