//++++++++++++++++++++ Error Handling ++++++++++++++++

// Need of this stuff ???

/*

so let say if u building a backedn and a user tries to log in 
the possible outcomes may be 
- user found
- database connection failed
- user does not exist
- wrong password

so how does Go tell us that something was wrong ? 
the answer is error type 

=> An error is simply a Value

*/


=> in Go : error is a built in interface 

=> nil : means , everything is OK
=> non-nil means, something went wrong

// Exmaples 

package main
import ("errors" "fmt")

func divide(a,b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a/b, nil
}

func main(){
	result, err := divide(10,2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// Practice Again 

package main
import ("fmt" "errors")

func divide(a,b int) (int, error){  // return two values that is result AND Error
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a/b, nil
}

func main(){
	result, err := divide(10,2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}


// Exmaples (Login System)


package main
import ("fmt" "errors")

func Login(email string) error {
	if email == "" {
		return errors.New("email can't be empty")
	}
	return nil
}

func main(){
	err := Login("")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Login Successful")
}

// ATM Example 

package main
import ("fmt" "errors")

func Withdraw(balance, amount int) (int, error) {
	if amount > balance {
		return balance, errors.New("insufficient balance")
	}
	return balance - amount, nil
}

func main(){
	newBalance, err := Withdraw(5000,7000)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Balance: ", newBalance)
}


// Examples 

package main
import ("fmt" "errors")

func CheckAge(age int) error {
	if age < 18 {
		return errors.New("you must be at least 180000 year old lol")
	}
	return nil
}

func main(){
	err := CheckAge(16)


	if err != nil {
		fmt.Println(err)
		return 
	}
	fmt.Println("Access Granted")
}


//+++++++++++++ The Golden Pattern ++++++++++++++

// you will write this thouands of times 

result, err := someFunction()

if err != nil {
	return err
}


