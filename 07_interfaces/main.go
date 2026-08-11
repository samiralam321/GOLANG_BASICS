//++++++++++++++++++++++ Interface +++++++++++++++++++++++

// creating an interface

type Animal interface {
	Sound()
}
// any type that has a Sound() method is an Animal
// interface contains only method signatrues, not implementations

// struct

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}


// complete Program 

package main
import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Name string
}

func (d Dog) Sound(){
	fmt.Println("Woof Woof")
}

func (c Cat) Sound(){
	fmt.Println("Mewon")
}

func main(){

	dog := Dog{
		Name : "Tommy"
	}

	cat := Cat {
		Name : "Kitty"
	}

	dog.Sound()
	cat.Sound()
}

// real power of interface 

func MakeSoud(a Animal){
	a.Sound()
}

dog := Dog {
	Name : "Tommy"
}

cat := Cat {
	Name : "Kitty"
}

MakeSound(dog) // woof woof
MakeSound(cat) // mewo mewo

//the same function works for both -> Polymorphism 


//+++++++++++ Real Backend Examples 

/*
Suppose you r building a payment application
userscan pay using 
credit card, upi, paypal

wothout interface 

paybyCard()
paybyUPI()
paybyPaypal()

with interface

type PaymentMethod interface {
    Pay(amount float64)
}

*/

// CARD

tpye CreditCard struct {}

func (c CreditCard) Pay(amount float64){
	fmt.Println("Paid", amount, "using Credit Card")
}

// UPI

type UPI struct{}

func (u UPI) Pay(amount float64) {
	fmt.Println("Paid", amount, "using UPI")
}

// Function 

func Checkout (p PaymentMethod){
	p.Pay(999)
}

CheckOut(CreditCard{})
Checkout(UPI{})

// Output : 
Paid 999 using Credit Card
Paid 999 using UPI


//++++++++++++++++++++++ Interface VS Strcut +++++++++++++++++

// Struct -> it ans what does this object have ?
type User struct {
	Name string
	Email string
}

// Interface -> it ans what can this object do ? 

type PaymentMethod interface {
	Pay(amount float64)
}

// strcut : Data (what it is)
// interface : Behaviour (what it can do)


//+++++++++++ Homework +++++++++++

// create an interface and find area  

package main
import ("fmt" "math")

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length float64
	Width float64
}

type Circle struct {
	Radius float64
}


func (r Rectangle) Area() float64 {
	return r.Length & r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main(){
	rect := Rectangle {
		Length : 10,
		Width : 5,
	}

	circle := Circle {
		Radius : 7,
	}

	fmt.Println("Rectangle Area :", rect.Area())
	fmt.Pritnln("Circle Area :", circle.Area())
}


///// Vechile Interface 

package main
import "fmt"

type Vehicle interface {
	Start()
}

type Car struct{}
type Bike struct{}

func (c Car) Start(){
	fmt.Println("Car engine Started")
}

func (b Bike) Start(){
	fmt.Println("Bike Engine Started")
}

func StartVehicle (v Vehicle){
	v.Start()
}

func main(){
	car := Car{}
	bike := Bike{}

	StartVehicle(car)
	StartVehicle(bike)
}


//////////////// Real Backend Examples 

/*
let say you r building a website and when user sign up you want to send 
welcome message , some time by email, sometime by SMS
so instead of writing two different functions , 
we use interface 

*/

package main
import "fmt"

type Notifier interface {
	Send(message string)
}

type Email struct{}
type SMS struct{}

func (e Email) Send(message string){
	fmt.Println("Email Sent : ", message)
}

func (s SMS) Send(message string){
	fmt.Println("SMS Sent : ", message)
}

func NotifyUser(n Notifier){
	n.Send("Welcome to our HAWELI LOL XD")
}

func main(){
	email := Email{}
	sms := SMS{}

	NotifyUser(email)
	NotifyUser(sms)
}