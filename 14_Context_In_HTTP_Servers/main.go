/*

Imagine you open your browser and visit:

http://example.com/users

Your browser sends a message to the server.

That message is called an HTTP request.
The server receives it and sends a message back.
That message is called an HTTP response.


Browser
   |
   | HTTP Request
   ↓
Go Server
   |
   | HTTP Response
   ↓
Browser


So in short : An HTTP request is a message sent by a client to a server

A handler's job is : Receive the HTTP request and decide what response to send



Request
   ↓
Handler
   ↓
Service
   ↓
Database

---

The handler says : "Someone requested users"
The service says : "Okay i know how to get users"
The database says : "Here are the users"

                 ┌──────────────┐
                 │   Browser    │
                 └──────┬───────┘
                        │
                   HTTP Request
                        │
                        ↓
                 ┌──────────────┐
                 │   Handler    │
                 └──────┬───────┘
                        │
                        ↓
                 ┌──────────────┐
                 │   Service    │
                 └──────┬───────┘
                        │
              ┌─────────┴─────────┐
              ↓                   ↓
        ┌──────────┐       ┌─────────────┐
        │ Database │       │ External API│
        └──────────┘       └─────────────┘

*/

/*

What Problem Context Solves ??

Imagine you send this request : 

GET /users/10

your servers starts processing it
may be : 

Browser
   ↓
Handler
   ↓
Service
   ↓
Database

but suddenyl...

the user close the browser
the reqiest is no longer useful
however your Go server might still be doing: 
database query or external API call or some expensive operations

so why should the server continus doing unncecessay work ??
it should not !!!!

we need a way to say : 

hey this reqest is no longer needed, so stop the work associated with it

this is one of the main jobs of context


*/

func getUsers(w https.ResponseWriter, r *https.Request) {
	fmt.Fprintln(w, "Users")
}

// r *https.Request - The request contains inforamtion about the HTTP request
// and r.Context() give you the context associated with that request

//******************** r.Context() ******************


func getUsers(w https.ResponseWriter, r *https.Request) {
	ctx := r.Context()
	fmt.Println(ctx)
}



//********** What is Context ***********

contect.Context


// you can think of Context as a small object that travels along with a request

HTTP Request
     |
     ↓
  Handler
     |
     ↓
  Service
     |
     ↓
 Database


 

 context travels alongside it: 

 Request
   |
   | Context
   ↓
Handler
   |
   | Context
   ↓
Service
   |
   | Context
   ↓
Database


// the most imp things Context can communicates are : 

// cancellation
// timeout
// deadlines
// request - scoped values 


//*************** Where does the Context comes from **********************

// when Go receives an HTTP request, the request already has a Context


func helloHandler(w http.ResponseWriter, r *http.Request)

// r represnt the HTTP request
// that request has r.Context()

ctx := r.Context() // gets the context associated with that request

// Example

func helloHandler(w https.ResponseWriter, r *https.Request) {
	ctx := r.Context()

	fmt.Println(ctx)
	fmt.Fprintln(w, "Hello!")
}



// usually we don't print the context
// we use it to pass cancellation, deadlines etc.


// when handling an HTTP request

ctx := r.Context()

// then pass that Context to function that perfom work

// Example

func helloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	doSomething(ctx)

	fmt.Println(w, "DONE")
}


// then

func doSomething(ctx context.Context) {
	// do some work
}

// and if that function calls another function 

func doSomething(ctx context.Context) {
	doDatabaseWork(ctx)
}

// then 

func doDatabaseWork(ctx context.Context) {
	// database work
}



// so the context travel down 

HTTP Request
     |
     ↓
 Handler
     |
    ctx
     ↓
 Service
     |
    ctx
     ↓
 Database


//************* Why pass Context everywhere *****************

// because all these operations belong to the same request

// imagine : 

User Request
     |
     ↓
Handler
     |
     ↓
Service
     |
     ↓
Database

// suppose the user cancel the request
// the context can communicate : "STOP!!" to the lower level operations 



User cancels request
        ↓
    Context
        ↓
     Handler
        ↓
     Service
        ↓
    Database

 its called cancellation propagation 



 //********************** First simple Context Example ******************

 package main
 import ("context" "fmt")

 func main(){
	ctx := context.Background() // creates a basic Context like "give me a starting context"
	fmt.Println(ctx)
 }


// practice 

 // Why Background()

 /*
coz contexts form a Tree

Background Context
       |
       ├── Request Context
       |
       ├── Another Request Context
       |
       └── Another Context


context.Background() is commonly used as a root Context when you do not already have one
// but in HTTP handler, you normally dont create a new Background context

INstead : ctx := r.Context()

coz Go already gives you the request Context

 */


 //**************** Context with timout *******************

 // supppose i have a function that takes 5 seconds 

 func dowork() {
	time.Sleep(5 * time.Second)
	fmt.Println("Work Completed")
 }


// but i want this function to allow only 2 second
// so we can create a timeout Context ;

ctx, cancel := context.WithTimeout (
	context.Background(),
	2*time.Second,
)

defer cancel()


// it means create a context that will automatically be cancelled after 2 seconds 

// but how dowa out function know about cancellation ? 

// doWork(ctx)

func doWork(ctx context.Context) {
	select {
	case <-time.After(5*time.Second):
		fmt.Println("Work Completed")
	case <- ctx.Done():   // wait untill the Context is cancelled
		fmt.Println("Work Cancelled")
	}
}



func doWork(ctx context.Context){
	select {
	case <-time.After(5*time.Second):
		fmt.Println("Work Completed")
	case <- ctx.Done():
		fmt.Println("Work Cancelled")
	}
}

// there are two possibilites 

// case 1 : Work finishes first : Work Completed
// case 2 : Context times out first : Work cancelled 

// coz the context timeout happened before the work finished


// Completed Examples

package main
import ("context" "fmt" "time")

func doWork(ctx context.Context) {
	fmt.Println("Work Started")
	select {
	case <- time.After(5*time.Second):
		fmt.Println("Work Completed")
	case <- ctx.Done():
		fmt.Println("Work Cancelled")
	}
}

func main(){
	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second, // so we create context, that allow this operation for 2 seconds
	)
	defer cancel()

	doWork(ctx) // start the work 
}

// the context only allow : 2 seconds , so aftr 2 seconds 
context -> cancelled -> ctx.Done() -> select receives it -> "Work cancelled"

//************* Create an HTTPS servers ******************

package main
import ("fmt" "net/http" "time")

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("Request received")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "work completed")
	case <- ctx.Done():
		fmt.Println("Request cancelled")
	}
}

func main(){
	http.HandleFunc("/work", handler)
	http.ListenAndServ(":8080", nil)
}


// Suppose the user request : GET /profile

func profileHandler(w http.ResponseWriter, r *https.Request){
	ctx := r.Context()

	prfile, err := getProfile(ctx)

	if err != nil {
		http.Error(w, "Something went wrong", 500)
		return
	}
	fmt.Fprintln(w, profile)
}


// Service :

func getProfile(ctx context.Context) (string, error){
	//database work
	return "Samir's profile", nil
}

// we pass the Context in the getProfile function coz, the profile operation belongs to this HTTP request

//****************** The 4 Context concepts you need to learn ***********************

// context.Context

ctx context.Context

//2. context.Background()

ctx := context.Background()

//3. context.Background()

ctx := context.Background()

//4. contxt.WithCancel : Creates a context that you manually cancel

ctx, cancel := context.WithCancel(parent)
defer cancel()

//5. context.WithTimeout() : Automatically cancels after a specified duration.

ctx, cancel := context.WithTimeout(
	parent,
	5*time.Second,
)

defer cancel()

//6. context.WithDeadline() :Cancels at a specific time

// ctx.Done() : the most important part for cancellation 

<- ctx.Done()

// ctx.Err()

// tells you why the context is stopped

//*****************************************************************
func someFunction(ctx context.Context) error {
	// work

	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := someFunction(ctx)

	if err != nil {
		// handle error
	}
}


//*************** Real life Analogy **********************

/* Imagine you are ordering a food, you place an order 123, the restruatant starts preapring it
The kitchen starts :
cooking -> packing -> delivery 

but then you cancel your order
you do not want the restaurant to continue preparing it

Context is somewhat like the cancellation/ lifetime signal attached to that order

Order -> Handler -> Service -> Databas/API

If the order is cancelled -> Stop the Work 

*/












