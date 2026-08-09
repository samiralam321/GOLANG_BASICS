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














