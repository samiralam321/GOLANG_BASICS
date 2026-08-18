/*
imagine your backend receives an HTTP request 

Client -> your go server -> Database

supppose the database operations takes 10 seconds

but after 2 seconds, the client leaves :

client -> request -> Go server -> database -> still working....

=> why should your server continue doing expensive work if nobody wants the result anymore

we want to able to say : "stop this work , the request has been cancelled"

This is one of the main job of context.Context

The context can communicate things like : 
"cancel the work"
"Time is over"
"The request has ended"
*/


// creating a context
ctx := context.Background()

// you need to import "context"

package main
import ("fmt" "context")

func main(){
	ctx := context.Background()  // context.Background() is basicallt your starting/ root context
	fmt.Println(ctx)
}

//********** Child Context ************

ctx := context.Background()

// we can create another context from it

ctx, cancel := context.WithCancel(ctx)

// cancel() can cancel the child

// so now we have parent context and child context

// FIrst Cancellation program 

package main
import (
	"context"
	"time"
	"fmt"
)

func worker(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():     // ctx.Done() gives us a channel, when the context is cancelled , channel is closed
			fmt.Println("Worker cancelled")
			return 
		default:
			fmt.Println("worker is working")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main(){
	ctx, cancel := context.WithCancel(context.Background())
	// it create two things 
	// ctx -> context
	// cancel -> function that cancels the context
	// so cancel() means calcel this context

	go worker(ctx)
	time.Sleep(2*time.Second)
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("Main finished")
}

// possible outputs :

worker is working
worker is working
worker is working
worker is working
worker is working
main finished


//*************** why use ctx.Done() with select ?

select {
case <-ctx.Done():
	return;

case job := <-jobs:
	process(job)
}

// it means : either process a job or stop if the context is cancelled 


//************ withTimeout **************

// instead of manually calling : cancel()

// we can say : this operation is allowed to run for at most 2 seconds 
// so use 

context.WithTimeout()

// Example : 

ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()


// Compelte timeout example

package main
import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <- ctx.Done():
		fmt.Println("worker stopped:", ctx.Err())
		return

		default:
		fmt.Println("worker working...")
		time.Sleep(500* time.Millisecond)
		}
	}
}

func main() {
	ctx,cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	go worker(ctx)

	time.Sleep(3*time.Second)

	fmt.Println("Main function")
}


//**************** ctx.Err() ********************

ctx.Err()

// it tells that why the context ended

// common result include : 

context cancelled 

or

context deadline exceeded

so :

//********************* WithCancel VS WithTimeout

// WithCancel : you decide when to stop : 

ctx, cancel := context.WithCancel(parent)
cancel()

// Manual stop


// WithTimeout : the context automatically stops after a duration:

ctx, cancel : context.WithTimeout(parent, 2*time.Second,)
// it menas : stop after 2 seconds 


// WithDeadline

context.WithDeadline()


// instead of give me 2 seconds, you say : stop at this specific time 

deadline := time.Now().Add(2* time.Second)

ctx, cancel := context.WithDeadline(
	context.Background(),
	deadline,
)

defer cancel()


//******************* Things you should remember ****************

WithCancel -> manual cancellation
WithTimeout -> cancel after duration
WithDeadline -> cancel at specific time


//********************* Real Backend Example : API timeout ******************

// imagine your server calls a payment service.

Your Backend -------request-------> Payment API

// you don't want to wait forever

ctx, cancel := context.WithTimeout(
	context.Background(),
	3*time.Second,
)

defer cacel()

result := callPaymentAPI(ctx)



// if the payment API does not repond within 3 seconds : 

// 3 seconds -> context cancelled -> stop waiting -> return timeout


//******************** Context + Worker Pool *******************

Earlier : 

Worker Pool
  jobs 
   |
Channel
   |
Worker 1
Worker 2
Worker 3


// now we can add context in workers 

// worker 

func worker(ctx context.Context, jobs <- chan int,) {
	for {
		select {
		case <- ctx.Done():
		return

		case job := <-jobs:
			fmt.Println("Processing", jobs)
		}
	}
}


// now workers can stop when the context is cancelled 


//**************** defer cancel() pattern ***************


ctx, cancel := context.WithTimeout(.....)
defer cancel()


ctx, cancel := context.WithTimeout(.....)
defer cancel()


ctx, cancel := contect.WithCancel(......)
defer cancel()



//******************** NOTES *******************

/* do not use context for everything

Context is mainly for : 

Cancellation
Timeouts
Deadlines
Request-scoped information 

*/

//************ Practice ************

worker starts 
works for 3 seconds 
main calls cancel()
worker notices cancellation
worker stops

package main
import ("fmt" "context" "time")

// worker receives a context , the context allow the worker to know when it should stop 
func worker(ctx context.Context) {
	for {
		select {
			case <- ctx.Done();  // this case become ready when cancel() is called 
			fmt.Println("worker stopped")
			return

		default:
			fmt.Println("worker working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}


func main(){
	ctx, cancel := context.WithCancel(contect.Background())

	// start the worker goroutine
	go worker(ctx)
	time.Sleep(3*time.Second)

	// tell the worker, you should stop now
	cancel()

	/// give the worker a little time to notice
	// the cancellation signal and print its message
	time.Sleep(500*time.Millisecond)
	fmt.Println("Main finished")
}



///

ctx, cancel := context.WithCancel(context.Background())

// you get two things

ctx -> context
cacel -> funtion that cancel the context



//******************* Look at how your previous topics are connecting ***************

Goroutines 
    |
"Run work concurrently"
    |
WaitGroup
	|
"wait for goroutines"
    |
Channel
	|
"send data between goroutines"
    |
 select
    |
"wait for multiple channel operations"
    |
Worker Pool
	|
"limits cuncurrent workers"
    |
Context
    |
"cancel/timeout the work"


// The most important Context Pattern 

ctx, cancel := context.WithTimeOut(context.Background(), 5*time.Second,)
defer cancel()

// then inside long running work : 

select {
case <- ctx.Done():
	return

case job := <-jobs:
	// do work
}





















































