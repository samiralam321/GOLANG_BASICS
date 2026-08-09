// Array => Array is a collection of elements of the same type

var nums[5]int
// this creates [0 0 0 0 0]


package main
import "fmt"

func main(){
	nums := [5]int{10,203,304,50}
	fmt.Println(nums)

	fmt.Println(nums[0]) // Accesssing Elements
	fmt.Println(nums[2])

	nums[1] = 12 // Updating Elements

	// Looping Through Arrays 
	for i:=0; i<len(nums); i++ {
		fmt.Println(nums[i])
	}

	// Range Loop
	for index, values := range nums {
		fmt.Println(index, value)
	}

	// Ignoring index
	for _, value := range nums {
		fmt.Println(value)
	}
}

// Sum of Array 

func arraySum(nums[5]int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

//+++++++++ Array Have Fixed Size +++++++++

var nums[5] int
// Exactly 5 elements
// cannot become 6,7,8 etc element that is why we use "SLICE"


//+++++++++++++++++++++ SLICES +++++++++++++++++++++++++++

// A Slice is a dynamic Array 
// Array + Flexibilility

// Creating a Slice 

nums := []int{10,20,40} // notice [] -> no size specified
fmt.Println(len(nums)) // 3

//+++++++++++ Append ++++++++++++

nums := []int{10,20,30}
nums = append(nums,40)  // adding single value
nums = append(nums, 50,60,70) // added multiple values 

//+++++++++ Real Backend Examples +++++++++++
//Users joining app :

users := []string{}

users = append(users,"Samir")
users = append(users, "Rahul")
users = append(users, "Diya")

// Result = [Samir Rahul Aman]


//+++++++++++++++++++ make() ++++++++++++++++

// Professional Go code uses this frequently

nums := make([]int, 5)

//creates [0 0 0 0 0] of length 5

// updating
nums[0] = 100
nums[1] = 200

// Result [100 200 0 0 0]


//+++++++++++++++++ SLICING ++++++++++++++++

// Create a slice from another slice

nums := [] int {10,20,30,40,50}

// First 3 Elements 
nums[:3]
// Output : [10 20 30]

// Last Elements
nums[2:]
// Output : [30 40 50]

// Middle Elements
nums[1:4]
//Output : [20 30 40]


//++++++++++++++++ MAPS ++++++++++++++++++++

// syntax  =>  map[keyType]valueType

// so maps are key value pairs

ages := map[string] int {
	"Samir" : 20,
	"Rahul" : 21,
}

// Access Value
fmt.Println(ages["Samir"])

// Add Entry
ages["Aman"] = 22

// delete entry
delete(ages, "Rahul")


// checky key exists 

age, exists := ages["Samir"]

fmt.Println(age) // 20 
fmt.Println(exists) // true


// Loop through Maps

for key, value := range ages {
	fmt.Println(key,value)
}

// Output : 
// Samir 20
// rahul 21

// note that order is not guranteed


///// Real Backend Examples 

users := map[int] string {
	1 : "Samir",
	2 : "Rahul",
	3 : "Aman",
}

// find users
fmt.Println(users[2])   // Rahul


/*


| Feature      | Array   | Slice  | Map  |
| ------------ | -----   |----    | ------|
| Fixed Size   | ✅     | ❌     | ❌   |
| Dynamic      | ❌     | ✅     | ✅   |
| Index Access | ✅     | ✅     | ❌   |
| Key-Value    | ❌     | ❌     | ✅   |


*/


package main
import "fmt"

func main(){
	ages := map[string] int {
		"Samir" : 20,
		"Rahul" : 21,
		"Aman" : 19,
	}
	fmt.Println(ages)
}

// output : map[Aman:19 Rahul:21 Samir:20]
// order is not guranteed cuz go maps are unordered


// All Map Operations 

package main
import "fmt"

func main(){
	// empty mapes
	ages := make(map[string] int )

	//add values 
	ages["Samir"] = 20
	ages["Sami"] = 22
	ages["Sam"] = 21

	fmt.Println("Intitial Map : ")
	fmt.Println(ages)

	//access values
	fmt.Println("\n Age of Samir : ")
	fmt.Println(ages["Samir"])

	// update values 
	ages["Samir"] = 99

	fmt.Println("\n After Update : ")
	fmt.Println(ages)

	// check if key exists
	age, exits := ages["Rahul"]

	if exists {
		fmt.Println("\n Rahul Found")
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Rahul not Found!")
	}

	// delete keys
	delete(ages,"Aman")

	// Length
	fmt.Println("\n Length:")
	fmt.Println(len(ages))

	// loop
	for key , value := range ages {
		fmt.Println(key, "->", value)
	}
}

///// Nested Maps (Advaced) /////////////////

users := map[string]map[string] string {
	"Samir" : {
		"city" : "Patna",
		"Role" : "SWE",
	},

	"Rahul" : {
		"City" : "Delhi",
		"Role" : "SWE",
	},
}

// fmt.Println(users["Samir"]["city"])


////////// Real Backend  Examples 

users := map[string] string {
	"samir@gmail.com" : "123445",
	"rahul@gmail.com" : "abcdef",
}

email := "samir@gmail.com"

password, exits := users[email]

if exists {
	fmt.Println("Password : ", password)
} else {
	fmt.Println("User not found!")
}



// Real use case of map
package main

import "fmt"

func main() {

	// Fake Database
	users := map[string]string{
		"samir@gmail.com": "123456",
		"rahul@gmail.com": "abcdef",
		"aman@gmail.com":  "hello123",
	}

	var email string
	var password string

	fmt.Print("Enter Email: ")
	fmt.Scan(&email)

	fmt.Print("Enter Password: ")
	fmt.Scan(&password)

	// Check if user exists
	storedPassword, exists := users[email]

	if !exists {
		fmt.Println("❌ User Not Found")
		return
	}

	// Check password
	if storedPassword == password {
		fmt.Println("✅ Login Successful")
	} else {
		fmt.Println("❌ Incorrect Password")
	}
}