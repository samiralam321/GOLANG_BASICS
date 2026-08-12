// problem!!

// simple program 

package main
import "fmt"

func task1(){
	fmt.Println("Task 1")
}

func task2(){
	fmt.Println("Task 2")
}

func main(){
	task1()
	task2()
}

// output :
// Task 1
// Task 2

// the flow : 

/*
main()
  |
  ↓
task1()
  |
  ↓
task1 finishes
  |
  ↓
task2()
  |
  ↓
task2 finishes
  |
  ↓
program ends
*/

// go waits for task1() to finish before starting task2()
// it is called sequential execution 

//+++++++++++++++++++++++++++ Goroutine ++++++++++++++++++++++++++

// a goroutine is a lightweight unit od execution manages by the Go runtime
// a goroutine lets you start a function so it can concurrently with the rest of your progrmm

// you put 'go' before a function call 


/////////////// NORMAL FUNCTION CALL

// task1()
// Run task1 and wait for it to finish

///////////// Goroutine
go task1()

// means => start task1 as a goroutine ans let the current goroutine continue


package main
import "fmt"

func task(){
	fmt.Println("Task is running")
}

func main(){
	go task()

	fmt.Println("Main is running")
}

// possible outcomes 

//1 :
// Task is running
// Main is running

// 2 : 
// Main is running

// 3 : 
// Main is running
// Task is running

//4 : Sometimes you may see
// Task is running
// Main is running


//////////// WHY>>>>>>>?????????????????????????????

// because of go task() ; the goroutine starts running independently
// meanwhile main() continues then main() finsihed, and when main function finished the entird Go program exits
// it does not wait automatically for your goroutines 


// we can temporarily use : time.Sleep()

package main
import ("fmt" "time")

func task(){
	fmt.Println("Task is running")
}

func main(){
	go task()

	time.Sleep(time.Second)  // time.Second means 1 second
	fmt.Println("Main is running")
}


// Output : 
// Task is running
// Main is running

// the program waits for one second, giving the goroutine time to execute

// time.Second - represent one second
// time.Sleep(time.Second) means pause the current gorountine for one second

// time.Sleep(2*time.Second) which means sleep for 2 seconds


// do not use time.Sleep() it's just for learning purpose

// imagine database operation takes 100ms and you sleep 5 sec , wasted 4.9 second 

// one function without Gorountines 

package main
import ("fmt" "time")

func task(){
	for i:=1; i<=5; i++ {
		fmt.Println("Task : ", i)
		time.Sleep(500* time.Millisecond)
	}
}

func main(){
	task()
	fmt.Println("Main finished")
}


// output : 
Task: 1
Task: 2
Task: 3
Task: 4
Task: 5
Main finished

// now add go 

package main

import (
	"fmt"
	"time"
)

func task() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task()

	time.Sleep(3 * time.Second)
	fmt.Println("Main finished")
}

//Now the task is running as a goroutine.
//The main goroutine is also running.


///+++++++++++++++++ Multiple Gorountines 

package main

import (
	"fmt"
	"time"
)

func numbers(){
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func letters(){
	for i := 'A'; i <= 'E'; i++ {
		fmt.Println("Letter:", string(i))
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go numbers()
	go letters()

	time.Sleep(3 * time.Second)
}

// possible output : 
Number: 1
Letter: A
Number: 2
Letter: B
Number: 3
Letter: C
Number: 4
Letter: D
Number: 5
Letter: E

OR

Letter: A
Number: 1
Letter: B
Number: 2
Number: 3
Letter: C
Number: 4
Letter: D
Number: 5
Letter: E

// exact order is not guranteed

// cause you told Go:
// go numbers()
// go letters()

// you did not say 
numbers must finish first

or letters must finish first

// both are allowed to run concurrently
// the Go runtime schedules the goroutines


//////// without goroutines 

Main
 │
 ▼
Numbers
 │
 ▼
Finished
 │
 ▼
Letters
 │
 ▼
Finished


// with gorountines

             Main
              │
       ┌──────┴──────┐
       ↓             ↓
   Numbers         Letters
       │             │
       ↓             ↓
    Running       Running
       │             │
       └──────┬──────┘
              ↓
           Finished


/// Real Backend Examples 

// Imagine a user registers and you backedn needs to :
 save data
 send welcome message
 send notifications 

 witout concurrency 

Save User
   ↓
Send Email
   ↓
Send Notification
   ↓
Response


if 

Save = 100ms
Email = 2 seconds
Notification = 1 second

totoal waiitng time can approach : 3.1 seconds 



////////////////////////// Using Goroutines 

you might have independent background tasks 

go sendEmail()
go sendNotification()

now they can execute concurrently
             Save User
                |
          ┌─────┴─────┐
          ↓           ↓
      Send Email   Notification
          ↓           ↓
          └─────┬─────┘
                ↓


//+++++++++++++ VVI : do not create Goroutine everywhere 

// goroutines are fast so i will put go before every function 

// NOOOOOOOOOOOOOOOOOO, you need to ask
// can this work safely and usefully happen concurrently

// for examples 

// save User()
// sendEmail()

// may be they can run independently

but 

createUser()
deleteUser()

// may have ordering dependency
if deleteUser() starts before creteUser() finishes, you may have a problem 

// Concurrency requires thinking about : 
Dependencies
Shared data
Ordering
Errors
Cancellation
Synchronization


//////////////////// Concurrency VS Parallesim 

/*
Concurrency

You have:

Task A
Task B

and you're making progress on both.


Parallelism

Task A and Task B are literally executing simultaneously on different CPU cores.
So:

Concurrency = dealing with multiple tasks
Parallelism = executing multiple tasks at the same time

Go supports both, but you don't need to manually control CPU cores for normal goroutine usage.

*/



// What is a goroutine?

// A lightweight concurrent execution unit managed by Go's runtime.









