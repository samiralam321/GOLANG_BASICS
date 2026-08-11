package main

import "fmt"

/*
===========================================================
                    MODULE 5: POINTERS
===========================================================

A pointer is a variable that stores the MEMORY ADDRESS
of another variable.

Example:

    age := 20
    p := &age

Here:

    age  -> stores 20
    &age -> gives the address of age
    p    -> stores that address
    *p   -> gives the value stored at that address

IMPORTANT OPERATORS:

    &x  -> address of x
    *p  -> value stored at address p

===========================================================
1. BASIC POINTER
===========================================================
*/

func basicPointer() {

	age := 20

	fmt.Println("Value:", age)
	fmt.Println("Address:", &age)

	// p stores the address of age
	p := &age

	fmt.Println("Pointer:", p)

	// *p gives the value stored at that address
	fmt.Println("Value using pointer:", *p)
}

/*
===========================================================
2. MODIFYING VALUE USING POINTER
===========================================================

Because p points to age, changing *p changes age.
*/

func modifyUsingPointer() {

	age := 20

	p := &age

	fmt.Println("Before:", age)

	*p = 50

	fmt.Println("After:", age)
}

/*
===========================================================
3. EXPLICIT POINTER DECLARATION
===========================================================
*/

func explicitPointer() {

	num := 100

	var p *int

	p = &num

	fmt.Println("Number:", num)
	fmt.Println("Address:", p)
	fmt.Println("Value:", *p)
}

/*
===========================================================
4. NIL POINTER
===========================================================

The zero value of a pointer is nil.

A nil pointer does not point to any valid value.

IMPORTANT:

Never do:

    *p

when p is nil.

That will cause a runtime panic.
*/

func nilPointer() {

	var p *int

	fmt.Println("Pointer:", p)

	if p == nil {
		fmt.Println("Pointer is nil")
	}
}

/*
===========================================================
5. POINTER AS FUNCTION PARAMETER
===========================================================

Go is PASS-BY-VALUE.

If we pass an int:

    change(num)

the function receives a copy.

If we pass a pointer:

    change(&num)

the function receives a copy of the pointer,
but that pointer points to the original value.

Therefore, we can modify the original value.
*/

func changeValue(num *int) {

	*num = 100
}

func pointerFunction() {

	num := 10

	fmt.Println("Before:", num)

	changeValue(&num)

	fmt.Println("After:", num)
}

/*
===========================================================
6. VALUE PARAMETER VS POINTER PARAMETER
===========================================================
*/

// This receives a copy.
func changeWithoutPointer(num int) {

	num = 100
}

// This receives a pointer.
func changeWithPointer(num *int) {

	*num = 100
}

func compareFunctionParameters() {

	num1 := 10
	num2 := 10

	changeWithoutPointer(num1)
	changeWithPointer(&num2)

	fmt.Println("Without pointer:", num1)
	fmt.Println("With pointer:", num2)
}

/*
===========================================================
7. SWAP USING POINTERS
===========================================================
*/

func swap(a *int, b *int) {

	*a, *b = *b, *a
}

func pointerSwap() {

	a := 10
	b := 20

	fmt.Println("Before:", a, b)

	swap(&a, &b)

	fmt.Println("After:", a, b)
}

/*
===========================================================
8. POINTER WITH STRUCT
===========================================================
*/

type User struct {
	Name  string
	Age   int
	Email string
}

func pointerWithStruct() {

	user := User{
		Name:  "Samir",
		Age:   20,
		Email: "samir@gmail.com",
	}

	// Pointer to the struct
	p := &user

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	fmt.Println("Email:", p.Email)

	/*
		Go automatically dereferences p when accessing
		struct fields.

		p.Name

		is effectively:

		(*p).Name
	*/
}

/*
===========================================================
9. MODIFY STRUCT USING POINTER
===========================================================
*/

func modifyStruct() {

	user := User{
		Name:  "Samir",
		Age:   20,
		Email: "samir@gmail.com",
	}

	p := &user

	p.Name = "Rahul"
	p.Age = 21

	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)
}

/*
===========================================================
10. POINTER RECEIVER
===========================================================

A pointer receiver allows a method to modify the
original struct.

Syntax:

    func (u *User) ChangeName(...)
*/

func (u *User) ChangeName(name string) {

	u.Name = name
}

func pointerReceiver() {

	user := User{
		Name:  "Samir",
		Age:   20,
		Email: "samir@gmail.com",
	}

	user.ChangeName("Rahul")

	fmt.Println("Updated name:", user.Name)
}

/*
===========================================================
11. VALUE RECEIVER
===========================================================

A value receiver receives a COPY of the struct.

Use it when the method only needs to read the data.
*/

func (u User) Display() {

	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
	fmt.Println("Email:", u.Email)
}

func valueReceiver() {

	user := User{
		Name:  "Samir",
		Age:   20,
		Email: "samir@gmail.com",
	}

	user.Display()
}

/*
===========================================================
12. VALUE RECEIVER CANNOT MODIFY ORIGINAL
===========================================================
*/

func (u User) ChangeNameWrong(name string) {

	u.Name = name
}

func valueReceiverExample() {

	user := User{
		Name: "Samir",
		Age:  20,
	}

	user.ChangeNameWrong("Rahul")

	// Original is still Samir.
	fmt.Println("Name:", user.Name)
}

/*
===========================================================
13. POINTER RECEIVER CAN MODIFY ORIGINAL
===========================================================
*/

func (u *User) ChangeNameCorrect(name string) {

	u.Name = name
}

func pointerReceiverExample() {

	user := User{
		Name: "Samir",
		Age:  20,
	}

	user.ChangeNameCorrect("Rahul")

	fmt.Println("Name:", user.Name)
}

/*
===========================================================
14. new() FUNCTION
===========================================================

new(int):

    - creates memory for an int
    - initializes it to zero
    - returns *int

Example:

    p := new(int)
*/

func newFunction() {

	p := new(int)

	fmt.Println("Initial value:", *p)

	*p = 500

	fmt.Println("Updated value:", *p)
}

/*
===========================================================
15. POINTER TO POINTER
===========================================================

A pointer can point to another pointer.

    num
      ↑
      |
      p
      ↑
      |
      pp

*/

func pointerToPointer() {

	num := 10

	p := &num

	pp := &p

	fmt.Println("num:", num)
	fmt.Println("*p:", *p)
	fmt.Println("**pp:", **pp)
}

/*
===========================================================
16. POINTER COMPARISON
===========================================================
*/

func pointerComparison() {

	num := 10

	p1 := &num
	p2 := &num

	if p1 == p2 {
		fmt.Println("Both pointers point to the same variable")
	}
}

/*
===========================================================
17. POINTER WITH SLICE
===========================================================

Slices already refer to an underlying array.

Therefore, you normally DON'T need:

    *[]int

just to modify elements.

*/

func pointerWithSlice() {

	nums := []int{10, 20, 30}

	nums[0] = 100

	fmt.Println("Slice:", nums)
}

/*
===========================================================
18. POINTER WITH MAP
===========================================================

Maps are also reference-like.

Normally you don't need:

    *map[string]int

to modify map values.
*/

func pointerWithMap() {

	ages := map[string]int{
		"Samir": 20,
		"Rahul": 21,
	}

	ages["Samir"] = 25

	fmt.Println("Map:", ages)
}

/*
===========================================================
19. POINTER TYPE
===========================================================

Examples:

    *int
    *string
    *float64
    *User

These mean:

    pointer to int
    pointer to string
    pointer to float64
    pointer to User
*/

func pointerTypes() {

	var a int = 10
	var b string = "Hello"
	var c float64 = 10.5

	var p1 *int = &a
	var p2 *string = &b
	var p3 *float64 = &c

	fmt.Println(*p1)
	fmt.Println(*p2)
	fmt.Println(*p3)
}

/*
===========================================================
20. POINTERS AND MEMORY
===========================================================

Conceptually:

    age := 20

    age
     |
     v
    +------+
    |  20  |
    +------+
      ^
      |
    address

    p := &age

    p
    |
    v
  address
    |
    v
  +------+
  |  20  |
  +------+

Therefore:

    &age -> address
    p    -> address
    *p   -> 20
*/

func memoryConcept() {

	age := 20

	p := &age

	fmt.Println("Address stored in p:", p)
	fmt.Println("Value at that address:", *p)
}

/*
===========================================================
21. POINTER ARITHMETIC
===========================================================

Go does NOT support normal pointer arithmetic.

You cannot do:

    p++
    p = p + 1

This is different from C/C++.
*/

/*
===========================================================
22. IMPORTANT: GO IS PASS-BY-VALUE
===========================================================

This is an important interview concept.

Go always passes values.

Even when we pass a pointer:

    change(&num)

the pointer itself is copied.

But the copied pointer points to the same original memory.

Therefore:

    *p = 100

can modify the original value.
*/

/*
===========================================================
23. REAL BACKEND EXAMPLE
===========================================================

Imagine a bank account.

We want Deposit() to modify the original balance.

Therefore we use a pointer receiver.
*/

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {

	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {

	if amount <= b.Balance {
		b.Balance -= amount
	} else {
		fmt.Println("Insufficient balance")
	}
}

func bankExample() {

	account := BankAccount{
		Owner:   "Samir",
		Balance: 10000,
	}

	fmt.Println("Initial balance:", account.Balance)

	account.Deposit(5000)

	fmt.Println("After deposit:", account.Balance)

	account.Withdraw(2000)

	fmt.Println("After withdrawal:", account.Balance)
}

/*
===========================================================
24. POINTER CHEAT SHEET
===========================================================

    &x

    Gives the ADDRESS of x.


    p := &x

    p stores the ADDRESS of x.


    *p

    Gives the VALUE stored at the address.


    *p = 100

    Changes the original value.


    *int

    Pointer to an int.


    *User

    Pointer to a User.


    nil

    Pointer points to nothing.


    new(int)

    Creates an int and returns *int.


    func (u *User)

    Pointer receiver.

===========================================================
25. MOST IMPORTANT RULES
===========================================================

1. & -> address

2. * -> value at address

3. Pointer stores an address.

4. Dereferencing means accessing the value
   through a pointer.

5. A nil pointer points to nothing.

6. Never dereference a nil pointer.

7. Use pointer parameters when a function
   needs to modify the original value.

8. Use pointer receivers when a method needs
   to modify the original struct.

9. Value receivers work with a copy.

10. Go is pass-by-value.

11. Go does not support normal pointer arithmetic.

12. Slices and maps usually don't need pointers
    for normal modifications.

13. new() returns a pointer.

===========================================================
*/

// =========================================================
// MAIN FUNCTION
// =========================================================

func main() {

	fmt.Println("========== 1. BASIC POINTER ==========")
	basicPointer()

	fmt.Println()

	fmt.Println("========== 2. MODIFY USING POINTER ==========")
	modifyUsingPointer()

	fmt.Println()

	fmt.Println("========== 3. EXPLICIT POINTER ==========")
	explicitPointer()

	fmt.Println()

	fmt.Println("========== 4. NIL POINTER ==========")
	nilPointer()

	fmt.Println()

	fmt.Println("========== 5. POINTER FUNCTION ==========")
	pointerFunction()

	fmt.Println()

	fmt.Println("========== 6. VALUE VS POINTER ==========")
	compareFunctionParameters()

	fmt.Println()

	fmt.Println("========== 7. SWAP ==========")
	pointerSwap()

	fmt.Println()

	fmt.Println("========== 8. POINTER WITH STRUCT ==========")
	pointerWithStruct()

	fmt.Println()

	fmt.Println("========== 9. MODIFY STRUCT ==========")
	modifyStruct()

	fmt.Println()

	fmt.Println("========== 10. POINTER RECEIVER ==========")
	pointerReceiver()

	fmt.Println()

	fmt.Println("========== 11. VALUE RECEIVER ==========")
	valueReceiver()

	fmt.Println()

	fmt.Println("========== 12. VALUE RECEIVER EXAMPLE ==========")
	valueReceiverExample()

	fmt.Println()

	fmt.Println("========== 13. POINTER RECEIVER EXAMPLE ==========")
	pointerReceiverExample()

	fmt.Println()

	fmt.Println("========== 14. new() FUNCTION ==========")
	newFunction()

	fmt.Println()

	fmt.Println("========== 15. POINTER TO POINTER ==========")
	pointerToPointer()

	fmt.Println()

	fmt.Println("========== 16. POINTER COMPARISON ==========")
	pointerComparison()

	fmt.Println()

	fmt.Println("========== 17. POINTER WITH SLICE ==========")
	pointerWithSlice()

	fmt.Println()

	fmt.Println("========== 18. POINTER WITH MAP ==========")
	pointerWithMap()

	fmt.Println()

	fmt.Println("========== 19. POINTER TYPES ==========")
	pointerTypes()

	fmt.Println()

	fmt.Println("========== 20. MEMORY CONCEPT ==========")
	memoryConcept()

	fmt.Println()

	fmt.Println("========== 23. REAL BACKEND EXAMPLE ==========")
	bankExample()
}



//////////////////////////////////////////


package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func changeValue(num *int) {
	*num = 100
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func (u *User) changeName(name string) {
	u.Name = name
}

func (u User) display() {
	fmt.Println("Name:", u.Name)
	fmt.Println("Age:", u.Age)
}

func main() {

	// 1. Basic Pointer

	age := 20

	fmt.Println("Value of age:", age)
	fmt.Println("Address of age:", &age)

	p := &age

	fmt.Println("Pointer p:", p)
	fmt.Println("Value using *p:", *p)

	// 2. Modify value using pointer

	*p = 30

	fmt.Println("Age after changing *p:", age)

	// 3. Explicit Pointer Type

	num := 50

	var ptr *int
	ptr = &num

	fmt.Println("Number:", num)
	fmt.Println("Pointer:", ptr)
	fmt.Println("Value using pointer:", *ptr)

	// 4. Nil Pointer

	var nilPointer *int

	fmt.Println("Nil pointer:", nilPointer)

	if nilPointer == nil {
		fmt.Println("Pointer is nil")
	}

	// 5. Pointer as Function Parameter

	number := 10

	fmt.Println("Before changeValue:", number)

	changeValue(&number)

	fmt.Println("After changeValue:", number)

	// 6. Swap Using Pointers

	a := 10
	b := 20

	fmt.Println("Before swap:", a, b)

	swap(&a, &b)

	fmt.Println("After swap:", a, b)

	// 7. Pointer with Struct

	user := User{
		Name: "Samir",
		Age:  20,
	}

	userPointer := &user

	fmt.Println("User name:", userPointer.Name)
	fmt.Println("User age:", userPointer.Age)

	// 8. Modify Struct Using Pointer

	userPointer.Name = "Rahul"

	fmt.Println("Updated user name:", user.Name)

	// 9. Pointer Receiver

	user.changeName("Aman")

	fmt.Println("Name after pointer receiver:", user.Name)

	// 10. Value Receiver

	user.display()

	// 11. new() Function

	pointerNumber := new(int)

	fmt.Println("Value using new:", *pointerNumber)

	*pointerNumber = 500

	fmt.Println("Updated value:", *pointerNumber)

	// 12. Pointer to Pointer

	x := 10

	p1 := &x
	p2 := &p1

	fmt.Println("Original value:", x)
	fmt.Println("Using *p1:", *p1)
	fmt.Println("Using **p2:", **p2)

	// 13. Pointer Comparison

	if p1 == *p2 {
		fmt.Println("Both pointers point to the same location")
	}

	// 14. Pointer with Slice

	nums := []int{10, 20, 30}

	nums[0] = 100

	fmt.Println("Slice:", nums)

	// 15. Pointer with Map

	ages := map[string]int{
		"Samir": 20,
	}

	ages["Samir"] = 21

	fmt.Println("Map:", ages)
}