////*************** ERROR ////////////////

/*

Suppose a user want to access the name of a user from the database
GET /users/100

But user 100 does not exist

so your program needs a way to say : 

"Something went wrong."  That's an error

In GO, errors are actually returned as values */

result, err := someFunction()

if err != nil {
	fmt.Println(err)
}

// First Program 

package main
import "fmt"

func divide(a int, b int) (int, err) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil    // nil means no ERROR
}

func main(){
	result, err := divide(10,2)

	if err != nil {
		fmt.Println("Error" , err)
		return
	}
	fmt.Println("Result : ", result)
}


// template :

result, err := function()

if err != nil {
	// handle the error
}

//****************** errors.New() ********************

package main
import (
	"fmt"
	"errors"
)

func login(password string) error {
	if password != "1234" {
		return errors.New("Invalid Password")
	}
}

func main(){
	err := login("1111")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Login Succesful")
}

// Error : Invalid Password


//****************** fmt.Errorf() ***********************

// foramatted errors 

fmt.Errorf("user %s not found", name)

// Exmaples : 

pacakge main
import "fmt"

func findUser(name string) error {
	return fmt.Errorf("user %s not found", name)
}

func main(){
	err := findUser("Samir")

	if err != nil {
		fmt.Println(err)
	}
}

// user Samir not found

//******************* errors.New() Vs fmt.Errorf() **********************

// use errors.New() when you have a fixed message
// use fmt.Errorf when you need a dynamic information

errors.New("Something went wrong or INvalid Password")
fmt.Errorf("users %s not found", name)


//********** Returnig Erros From Functions 

func registerUser(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}



//*********** Mutiple Return values + Error


package main
import ("fmt" "errors")

type User struct {
	Name string
	Age int
}

func getUser(id int) (User, error) {
	if id < 0 {
		return User{}, errors.New("invalid user ID")
	}

	user := User {
		Name : "Samir",
		Age : 20,
	}
	return user, nil
}

func main(){
	user, err := getUser(1)

	if err != nil {
		fmt.Println("Errors:", err)
		return
	}

	fmt.Println("User:", user.Name)
	fmt.Println("Age:", user.Age)
}


//*************** Real Backend Example  : User Login ///////////

package main
import (
	"fmt"
	"errors"
)

func login(email string, password string) error {
	if email == "" {
		return errors.New("email is required")
	}

	if password == "" {
		return errros.New("password is required")
	}

	if email != "samir@gmail.com" {
		return errors.New("user not found")
	}

	if password != "123456" {
		return errors.New("incorrect password")
	}

	return nil
}


func main(){
	err := login("samir@gmai.com", "123456")

	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}
	fmt.Println("Login Successful")
}


//************** Error Propagation 

func getUser() error {
	err := databaseCall()

	if err != nil {
		return err
	}
	return nil
}

// this is called propagating the error



//********************************* panic ******************************

package main
import "fmt"

func main(){
	fmt.Println("Start")

	panic("something went terribly wrong")

	fmt.Println("END")
}


// output : 

Start
panic : something want terribly wrong
...

The program stops

//**************** panic VS errors ******************

Erros : 

// Expected or manageable failure:
User not found
Invalid input
Database unavailable
File does not exist

usually return err


Panic : 

A severe situation where normal execution cannot continue, 
or an unrecoverable programming/runtime problem

panic("....")

// for normal business failures, prefer erros instead of panic



// NOTE : Don't Use Panic for Normal User Errors



//********************** recover *********************

// recover(), which can catch a panic when used inside a deferred function

pacakge main
import "fmt"

func main(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("something went wrong")

	fmt.Println("This won't execute")
}

// Output : Recovered : Something went wrong

// recover is mainly for handling panics at appropriate boundaries, 
// such as protecting a server from crashing because one request panicked.


//***************** defer ******************

defer fmt.Pritnln("Goodbye")

// means : Run this when the surrounding function is about to return 


package main
import "fmt"

func main(){
	defer fmt.Println("Goodbye")

	fmt.Println("Hello")
}

// Output : 

// Hello
// Goodbye


////////// Multiple Defer 

pacakge main
import "fmt"

func main(){
	defer fmt.Println("First")
	defer fmt.Println("Second")
	defer fmt.Println("Third")

	fmt.Println("main")
}

/*
Main
Third
Second
First
*/

// defer follows : LIFO



//    Why is defer is useful in Backedn Development 

// it is very useful for cleanup

file, err := os.Open("data.txt")

if err != nil {
	return err
}

defer file.Close()

// meaning : Once this function finishes, close the file


//***************** Real Backend-Style Erros Flow **********************

HTTP Request
     ↓
Handler
     ↓
Service
     ↓
Database

Database return : user not found

Repository :

return User{}, errors.New("user not found")

Service : 

user, err := getUser(id)
if err != nil {
	return User{}, err
}


Handler : 

user, err := service.GetUser(id)

if err != nil {
	// convert error into HTTP response
	return
}

// this is how error travel through backend layer

