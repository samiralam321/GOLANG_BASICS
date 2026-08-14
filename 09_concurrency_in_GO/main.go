// problem!!

// simple program

package main

import (
	"fmt"
	"sync"
	"time"
)

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



//++++++++++++++++++++++++++ Sync waitGroups ++++++++++++++++++++++++++

// the problem we had 

package main
import "fmt"

func task(){
	fmt.Println("task Completed")
}

func main(){
	go task()

	fmt.Println("Main Finished")
}

// you might get only get 
// "Main Finished"


becuase 

main()
  ↓
start goroutine
  ↓
print Main finished
  ↓
main ends
  ↓
program ends

// the goroutines may not get enough time to execute

// we used time.Sleep() : but this is not a proper solution
// so Go gives us : sync.WaitGroup



/////////////// WaitGroups ////////////////////

// a WaitGroup lets you wait untill a group of goroutines finished

// creating a WaitGroup : var wg sync.WaitGroup


package main
import ("fmt" "sync")

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this goroutine is done
	fmt.Printf("worker %d started\n", i )
	fmt.Printf("worker %d end\n", i)
	//wg.Done() bettwe way to write this like defer.wg.Done
	// function ke just starting me likg do 
}



func main(){

	var wg sync.WaitGroup

	for i:=1; i<=3; i++ {
		wg.Add(1)  // increment the waitgroup counter 
		go worker(i, &wg)
	}

	// wait for all workers to finish
	wg.Wait()
	fmt.Println("worker task complted")

}


/// Practice WaitGroups 

package main
import ("fmt" "sync")

func worker(i int){
	defer wg.Done()   // when this worker function finshed automatically call wg.Done(); when this function exits , tell waitgroup i m done
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d end\n", i)
}

func main(){
	var wg sync.WaitGroup

	for i:=1; i<=3; i++ {
		wg.Add(1) // increase counter
		go worker(i, &wg)
	}

	wg.Wait() // wait unitll conuter becomes 0 ; wait untill everbody is done
	fmt.Println("worker task completed")
}

//+++++++++++++++++ what problem WaitGroup Solved?? ++++++++++++++++++

/*
Suppose we have 3 workers 

go worker(1)
go worker(2)
go worker(3) 

fmt.Println("All workers completed!")

The problem is go worker() means "fun this function concurrently"

so Go does not wait this function concurrently

It might print : 

All workers complted
worker 1 started
worker 2 started
worker 3 started

or the program might even finish before worker print anything 

why ? 

becuae main() can finish, and when main() finsihes, the entire Go program stops

So we need a way to tell main()

"wait untill all my goroutines are finished"

this is exactly what sync.WaitGroup does 




//////////////// The Three Important Functions 

Add() : Tell waitGroup I am going to start a gorountine

wg.Add(1)
go worker()


Done() 
Teels WaitGroup : This goroutines has finished

func worker(){
      fmt.Prinln("wokrer finished")
	  wg.Done()
}



Wait() : it means wait untill all workers finish 
ag.Wait()

*/


//++++++++++++++++++++++ A Real World Examples ++++++++++++

Imagine a backedn server needs to fetch 

user information 
orders
notifications 
recommendation 

instead of doing : 

fetch user
   ↓
fetch orders
   ↓
fetch notifications
   ↓
fetch recommendations


we can potententailly do 

          MAIN
           |
     ┌─────┼─────┬─────┐
     ↓     ↓     ↓     ↓
    User Orders Notif  Recomm.
     ↓     ↓     ↓     ↓
     └─────┼─────┴─────┘
           ↓
       WaitGroup
           ↓
       Response


//One thing you must NO do 

do not do this 

go worker(i, &wg)
wg.Add(1)

// the correct order is : 

wg.Add(1)
go worker(i, &wg)

/////////////////////

sync.WaitGroup
    ↓
Add()
    ↓
go function()
    ↓
Done()
    ↓
Wait()


//////////////// Real Life Backend Examples 

package main
import (
	"fmt"
	"sync"
	"time"
)

func processOrder(orderID int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Processing Order", orderID)
	time.Sleep(1 * time.Second)

	fmt.Println("Order", orderID, "Completed")
}

func main(){
	var wg sync.WaitGroup

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go processOrder(i, &wg)
	}
	wg.Wait()

	fmt.Println("All orders processed")

}


/// Send Emails to users 

// imagine you have to send an email to 4 users 

package main
import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(user string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Sending email to", user)
	time.Sleep(2* time.Second)

	fmt.Println("Email sent to", user)
}


func main(){
	var wg sync.WaitGroup

	users := []string {
		"Samir",
		"Rahul",
		"Aman",
		"Proya",
	}

	for _, user := range users {
		wg.Add(1)
		go sendEmail(user, &wg)
	}
	wg.Wait()

	fmt.Println("All emails sent")
}


////// Fetch Data rom multiple APIs

// suppose your application needs : 
// user API
// Payment API
// Order API

// so we can run this task concurrently

package main
import (
	"fmt"
	"sync"
	"time"
)

func getUser(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fetching user data...")
	time.Sleep(2*time.Second)
	fmt.Println("User data received")
}

func getOrders(wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Println("Fetching Orders...")
	time.Sleep(3*time.Second)
	fmt.Println("User data received")
}

func getOrder(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fetching orders...")
	time.Sleep(3*time.Second)
	fmt.Println("Order received")
}

func getPayment(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fetching Payment")
	time.Sleep(1*time.Second())
	fmt.Println("Payment Received")
}

func main(){
	var wg sync.WaitGroup

	wg.Add(1)
	go getUser(&wg)

	wg.Add(1)
	go getOrder(&wg)

	wg.Add(1)
	go getPayment(&wg)

	wg.Wait()

	fmt.Println("All data received")
}



//////// Imagine your backend needs to perfom 3 jobs 

package main
import (
	"fmt"
	"sync"
	"time"
)

func generateReport(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating Report...")
	time.Sleep(2*time.Second)
	fmt.Println("Report Generated")
}

func sendNotification(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Sending Notifications...")
	time.Sleep(1*time.Second)
	fmt.Println("Notication sent")
}

func updateDatabase(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Updating Database")
	time.Sleep(3*time.Second)
	fmt.Println("Database updated!!")
}


func main(){
	var wg sync.WaitGroup

	wg.Add(1)
	go generateReport(&wg)

	wg.Add(2)
	go sendNotification(&wg)

	wg.Add(1)
	go updateDatabase(&wg)

	wg.Wait()

	fmt.Println("All background jobs completed")
}


//++++++++++++++++++++++++++++++ Go Channels +++++++++++++++++++++++++++++++++

let say Gorountine 1 produce some data and Gorountine needs that data 
so how does Goroutine 1 give data to Goroutine 2 ? 

This is where Channels come in.

// What is Cahnnel ? 
=> A Channel is simply a way for goroutines to communicate with each other

Think of channel like a pipe : 

Goroutine 1
     |
     | sends data
     ↓
  [ CHANNEL ]
     |
     | receives data
     ↓
Goroutine 2


// In short : A channel is communication pipe used by goroutines to send and receive data 
thats it!!!!

/// Creating a Channel

the syntax is : 

ch := make(chan int)

make(chan int) => means : create a channel that can carry int values

chan int : channels that carries integer
chan string : channels that carries integers
chan bool : channel that carries true/false values 

ch1 := make(chan int)
ch2 := make(chan string)
ch3 := make(chan bool)


// Sending data into a channel

ch <- 10 : Put 10 into channel ch
ch <- 50


// Receiving data from channel
// now suppose someone put 10 inside the channel

// we can receive it 

x := <- ch

// read it like  : take a value from channel ch and store it in x

so 

ch <- 10 : means, send 10 

and 

x := <- ch : means, receive 10 

// There are only two basic operations 

SEND 

ch <- value

RECEIVE

x := <- ch


package main
import "fmt"

func main(){
	ch := make(chan int)

	ch <- 10
	x := <- ch

	fmt.Println(x)
}

// you might get 10, but this program gets stuck cause this is an unbuffered channel


///// Unbuffered Channel

ch := make(chan int)



// first proper channel program 

package main
import "fmt"

func sendData(ch chan int){ // this functio send the value
	ch <- 10
}

func main(){
	ch := make(chan int) // creating an integer channel

	go sendData(ch) // start a gorountine
	x := <-ch // main receive the value

	fmt.Println(x) // 10
}


// Channel + Goroutine = Communication 


/////////// Real Life Example : Order Processing //////////////

// imagine an e-commerce application 
// one gorountine createOrder and second process it 

package main
import "fmt"

func createOrder(ch chan int) {
	orderID := 101

	fmt.Println("Order Created : ", orderID)
	ch <- orderID
}

func processOrder(ch chan int) {
	orderID := <-ch

	fmt.Println("Processing Order : ", orderID)
}

func main(){
	ch := make(chan int)

	go createOrder(ch)
	go processOrder(ch)

	select{}
}


//

package main
import "fmt"

func calculate(ch chan int) {
	result := 10 + 20

	ch <- result
}

func main(){
	ch := make(chan int)

	go calculate(ch)

	result := <- ch

	fmt.Println("Result : ", result) // 30

}


// Channel with string 

package main
import "fmt"

func worker(ch chan string) {
	ch <- "Task Completed"
}

func main(){
	ch := make(chan string)

	go worker(ch)

	message := <- ch

	fmt.Println(message)
}

// Channel with multiple values 

package main
import "fmt"

func worker(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30
}

func main(){
	ch := make(chan int)

	go worker(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}



////// Multiple Workers + one channel

// suppose we have 3 workers 

package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan int, wg *sync.WaitGroup){
	defer wg.Done()
	ch<- id
}


func main(){
	var wg sync.WaitGroup

	ch := make(chan int)

	for i:=1; i<=3; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received : ", value)
	}
}


///////////// Recap : WaitGroup Vs Channel ///////////////

WaitGroup : Wait until goroutines finish

wg.Add(1)
go worker(&wg)
wg.Wait()


Channel : Used for Send data between gorountines 

ch := make(chan int)
go worker(ch)
result := <-ch


So, 

WaitGroup
    ↓
Synchronization / waiting

Channel
    ↓
Communication / data transfer


/// Note that 

ch := make(chan int)   this creates an UNBUFFERED CHANNEL

////++++++++++++++++++++++++  BUFFERED CHANNEL +++++++++++++++++++++++

ch : make(chan int, 3) means channel with capacity 3 which can hold 3 values

3 is the capcity , so this channel can temporarily hold 3 values like a box with 3 spaces



// unbuffered 

ch : make(chan int)

no storage, sender and receiver need to coordinate

// buffered 

ch := make(chan int, 3)

can store 3 values 
sender can send without an immediate receiver, as long as buffer is not full 


// let see an unbuffered channel 

package main
import "fmt"

func main(){
	ch := make(cha int)

	ch <- 10
	fmt.Println("DONE")
}


this program get stuck coz ch <- 10 is trying to send 10 but there is no receive
so it will wait until somebody receives this values but nobody does therefore the program is deadlock 

/// now buffered channel 

package main
import "fmt"

func main(){
	ch := make(chan int, 3)
	ch <- 10

	fmt.Println("DONE")
}

this channel has room i mean 3 space so the sender can put 10 inside
then println "DONE" this is output


// The important Differnce is 

Unbuffered : send -> needs receiver

Buffered : send -> can stroe value if buffer has space


////////// Sending Multiple Values

package main
import "fmt"

func main(){
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("All values sent")
}


// so this channel looks like 

--------------
 10 | 20 | 30|
--------------


if we send one more 

ch <- 40
as there is no space for 40 so the sender blocks, it waits until somebody receives a value

if we recieve 

x := <- ch

now x  = 10

and channel becomes 

20 30 ___

then x := <- ch
gives 20
then 30 

************** so CHANNEL BEHAVES LIKE A QUEUEU ************************

// First in First Out // FIFO Order

// let's create a produces 

package main
import "fmt"

func producer (ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println("Producer finished")

}


func main(){
	ch := make(chan int, 3)
	go producer(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


// output 

10
20
30

the producer puts values into the buffer
then main receives them 


/////////////// RECAP ///////////

with an unbuffered channel :

worker -> wait

with a buffered channel : 

worker -> channel buffer
workers can temporarily put results into the buffer


//**************** Common Confusion ******************

ch := make(chan int, 5)

the 5 means : The channel can hold 5 values 

it does NOT mean :
Create 5 goroutines 

These are seperate concepts 


//////********************* Practice  ///////////////////

// buffered channel with 3 integers 

package main
import "fmt"

func main(){
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/// Real life job queue

package main
import "fmt"

func main(){
	jobs := make(chan string, 5)

	jobs <- "send email"
	jobs <- "process payment"
	jobs <- "generate invoice"
	jobs <- "upload files"
	jobs <- "create report"

	fmt.Println("Jobs in queue")

	fmt.Println(<-jobs)
	fmt.Println(<-jobs)
	fmt.Println(<-jobs)
	fmt.Println(<-jobs)
	fmt.Println(<-jobs)
}

// so till here we have covered 

Goroutines -> WaitGroup -> Channels -> Unbuffered channels -> Buffered Channels

//********************************* Closing Channels + Range **********************************

// it ans a very common qns

// How does the receiver know that the sender has finished sending data ?????

suppose a worker sends 3 numbers 

ch <- 10
ch <- 20
ch <- 30

the receiver can do : 

fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)

but what if we do not know how many values the worker will send ??? 
may be 10, 15 ,20, .....

how does the receiver know when to stop ????

so here "close(ch)" comes in 

close(ch) => when the sender is completely finished sending values, it can say : 
close (ch)


func worker(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)  // closing does not mean delete channel it means : NO more value will be sent
}

// Example 

package main
import "fmt"


func main(){
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//***************** need of close() ***************

cause otherwise the receiver does not know "Are there more values coming ? "
The receiver does not know whether the worker is finished
Closing tells the receiver
Now receivers knows there wont be any more values



//*************** range over a channel ************

for value := range ch {
	fmt.Println(value)
}

it means : keep receiving values from the channel until the channel is closed 

package main
import "fmt"

func worker(ch chan int) {
	ch <- 10
	ch <- 20
	ch <- 30

	close(ch)
}

func main(){
	ch := make(chan int)
	go worker(ch)

	for value := range ch {
		fmt.Println(value)
	}
}


// Output : 
10
20
30







































































