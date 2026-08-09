//+++++++++++++++++++++ Struct and Methods ++++++++++++++++++

// A struct groups related data together

// Creating a struct
package main
import "fmt"

type Student struct {
	Name string
	Age int
	Branch string
	CGPA float64
}

func main(){
	var s Student

	s.Name = "Samir"
	s.Age = 20
	s.Branch = "ECE"
	s.CGPA = 8.7

	fmt.Println(s)

	// accessing fields
	fmt.Println(s.Name)
	fmt.Println(s.Age)
	fmt.Println(s.Branch)
}


// Initializing Struct (Recommended)

student := Student {
	Name : "Samir",
	Age : 20,
	Branch : "CSE",
	CGPA : 7.8,
}
fmt.Println(student)


//  Anonymous Initialization

student := Student {
	"Samir",
	20,
	"CSE",
	9.9,
}
// avoid this 


/////////// Mutiples Objects 

student1 := Student{
	Name: "Samir",
	Age: 20,
	Branch: "CSE",
	CGPA: 8.7,
}

student2 := Student{
	Name: "Rahul",
	Age: 21,
	Branch: "IT",
	CGPA: 9.1,
}

fmt.Println(student1)
fmt.Println(student2)


//++++++++++++++++++++++ Slice of Structs ++++++++++++++++++++

students := [] Student {
	{
		Name : "Samir",
		Age : 20,
		Branch : "CSE",
		CGPA : 7.7
	},

	{
		Name : "X",
		Age : 7,
		Branch : "ECE",
	},
	{
		Name: "Aman",
		Age: 19,
		Branch: "ECE",
		CGPA: 8.3,

	},
}


//////// Loop through items 

for _, student := range students {
	fmt.Println(student.Name)
	fmt.Println(student.Branch)
}



/////// Practice 

users := []User{

	{
		ID:1,
		Name:"Samir",
		Email:"samir@gmail.com",
	},

	{
		ID:2,
		Name:"Rahul",
		Email:"rahul@gmail.com",
	},

}

for _, user := range users {
	if user.ID == 2 {
		fmt.Println(user.Name)
	}
}


//+++++++++++++++++++++ Methods ++++++++++++++++++++++++

// a method belongs to a struct

type Student struct {
	Name string
	Age int
}

// method

func (s Student) Introduce(){
	fmt.Println("Hi, I am,", s.Name)

}

student := Student {
	Name : "Samir",
	Age : 20,
}

student.Introduce()

// Complete Program 

package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Branch string
	CGPA   float64
}

func main() {

	// var s Student

	// s.Name = "Samir"
	// s.Age = 20
	// s.Branch = "CSE"
	// s.CGPA = 8.8

	// fmt.Println(s)

	s := Student {
		Name : "Samir",
		Age : 20,
		Branch : "CSE",
		CGPA : 7.7,
	}

	// this is called struct literal
}



/// Examples 

package main

import "fmt"

type Student struct {
	ID        int
	Name      string
	Branch    string
	Semester  int
	CGPA      float64
	Email     string
	IsPlaced  bool
}

func main() {

	student1 := Student{
		ID:       101,
		Name:     "Samir Alam",
		Branch:   "CSE",
		Semester: 5,
		CGPA:     8.74,
		Email:    "samir@gmail.com",
		IsPlaced: false,
	}

	fmt.Println(student1)
}



// Real life Example 

package main
import "fmt"

type Order struct {
	OrderID int
	CustomerName string
	FoodName string
	Price float64
	Status string
	Address string
}

func main() {
	order := Order{
		OrderID:1001,
		CustomerName:"Samir",
		FoodName:"Chicken Biryani",
		Price:249,
		Status:"Preparing",
		Address:"Patna",
	}
	fmt.Println(order)
}


//++++++++ Slice of Structs ++++++++++++ (VERY IMP)

// instead of writing : 
user1 := User{}
user2 := User{}
user3 := User{}

// we create : 

users := []User{
	{
		ID:1,
		Name:"Samir",
		Email:"samir@gmail.com",
	},

	{
		ID:2,
		Name:"Rahul",
		Email:"rahul@gmail.com",
	},

	{
		ID:3,
		Name:"Aman",
		Email:"aman@gmail.com",
	},
}

// printing all the users 

for_, user := range users {
	fmt.Println("--------------------")
	fmt.Println("ID :", user.ID)
	fmt.Println("Name :", user.Name)
	fmt.Println("Email :", user.Email)
}


//++++++++++++++++++++ Methods - Part - 2 ++++++++++++++++++++++

// a method belongs to a struct

// function 

func greet(name string) {
	fmt.Println("Hello", name)
}
greet("Samir")



// Method

func (s Student) Greet() {
	fmt.Println("Hello", s.Name)
}

// call
student.Greet()


/////////////// First Method //////////////// 

package main
import "fmt"

type Student struct{
	Name string
	CGPA float64
}

func(s Student)Display(){  // this method belongs to Student; s is called receiver
	fmt.Println("Name :", s.Name)
	fmt.Println("CGPA :", s.CGPA)
}

func main(){
	student := Student{
		Name: "Samir",
		CGPA: 8.9,
	}
	student.Display()

}


////////// Method Returning a Value

type Rectange struct {
	Length float64
	WIdth float64
}

func (r Rectangle) Area() float {
	return r.Length * r.Width
}

func main(){
	rect := Rectangle {
		Length : 10,
		Width : 5,
	}

	fmt.Println(rect.Area()) // 50
}


////////// Multiple Methods 

type Employee struct {
	Name   string
	Salary float64
}

// method 1 
func (e Employee) Display() {
	fmt.Println(e.Name)
	fmt.Println(e.Salary)
}

// method 2 
func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

// call
employee := Employee{
	Name: "Samir",
	Salary: 50000,
}
employee.Display()
fmt.Println(employee.AnnualSalary())




////////// Real Backend Examples 

package main
import "fmt"

type BankAccount struct{
	Name    string
	Balance float64
}

func (b BankAccount) CheckBalance(){
	fmt.Println("Account Holder:", b.Name)
	fmt.Println("Balance:", b.Balance)
}

func main() {
	account := BankAccount{
		Name: "Samir",
		Balance: 45000,
	}
	account.CheckBalance()

}