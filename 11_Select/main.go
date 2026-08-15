//*************** Select **************

/*

select lets a goroutines basically say : 

i'm waiting for multiple channel operations.
whichever one is ready first, i'll handle it

select lets a goroutines wait for multiple channel opeations 
and execute ahichever one becomes ready first

what if you have TWO channels ?

channe 1 -> maybe some data
channel 2 -> maybe some data 

which one should we wait for ? 
this is where select comes in ? 

select says : i'll wait for both, whichever channel 
becomes ready first i will handle that one.



*/

// Syntax 

select {

case value := <- ch1 :
	fmt.Println(value)


case value := <-ch2 :
	fmt.Println(value)
}


// example

package main
import "fmt"

func main(){
	ch1 := make (chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Message from channel 1"
	}()

	go func() {
		ch2 <- "Message from channel 2"
	}()

	select {
	case message := <-ch1:
		fmt.Println(message)

	case message := <-ch2 :
		fmt.Println(message)
	}
}

// possible outputs :
Messae from channel 1
OR
Message from channel 2

coz select chooses whicever channel is ready.


// Real life examples to understand it 

/*
imagine you are waiting for two people to call you
you don't say i will wait for friend A for 10 minutes and then friend B

you say : i will answer whoever calls first


*/

// exmaple with different speeds 
// let's make one channel send after 2 seconds and another after 1 second

package main
import ("fmt" "time")

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2*time.Second)
		ch1 <- "Channel 1 is finished"
	}()

	go func() {
		time.Sleep(1*time.Second)
		ch2 <- "Channel 2 finished"
	}()

	select {

	case message := <-ch1:
		fmt.Println(message)

	case message := <-ch2:
		fmt.Println(message)
	}
}

// Output : Channel 2 finshed
// coz take only 1 second so it will be selected first

// selected does NOT wait for BOTH
// it means wait for ch1 OR wait for ch2 : whichever is ready first win


//************** What if both are ready at the same time **************

// if multiple cases are ready, Go chooses one of them


//************** select with sending ********************

case value := <-ch     which is receiving

select can also handle sending

select {
case ch1 <- 10:
	fmt.Println("Sent to ch1")

case ch2 <- 20:
	fmt.Println("Sent to ch2")
}

// this means : Send 10 to ch1 or 20 to ch2, whichever send operation 
// is ready first

so select can handle

Receive + Send

////////// Receive Example

select {
case value := <-ch1:
	fmt.Println("Recieved : ", value)

case value := <-ch2:
	fmt.Println("Recieved : ", value)
}



///////// Send Example

select {

case ch1 <- 10:
	fmt.Println("Sent 10")

case ch2 <- 20:
	fmt.Println("Sent 20")
}


//**************** The default case *****************

select {
case value := <-ch:
	fmt.Println(value)
}

// Normally if nothing is ready, the gorountine waits
// but we can add

default :
fmt.Println("Nothing available")


/////// Example

package main
import "fmt"

func main(){
	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println(value)
	
	default:
		fmt.Println("Nothing is available")
	}
}

// output : Nothing is available

// WHY ?
// cause nobody is sending into ch
// so instead of waiting forever, default runs immediately



/// Check for new Orders :

// a backend wokrer checking whether a new order is avialble 

package main
import "fmt"

func main(){
	orders := make(chan int)

	select {
	case orderID := <-orders:
		fmt.Println("Processing Order : ", orderID)

	default:
		fmt.Println("No new orders")
	}
}


// Output : No new orders 
// the worker does not get stuck waiting

////// Real life backend examples 
// suppose your backedn is waiting for responses from :
payment service
user servise

package main
import ("fmt" "time")

func main(){
	payment := make(chan string)
	user := make(chan string)

	go func() {
		time.Sleep(2*time.Second){
			payment <- "Payment succesful"
		}()
	}

	go func() {
		time.Sleep(1*time.Second)
		user <- "User data recevied"
	}()

	select {
	case result := <- payment:
		fmt.Println(result)
	case result := <-user :
		fmt.Println(result)
	}
}

// User data received
// coz User Servies responded first

//***************** Select with multiple iteration ***************

// a select execute once

select {
case value := <- ch1:
	fmt.Println(value)
case value := <- ch2:
	fmt.Println(value)
}

// after one case execute : select end

// but if you want to continusoly listen to channels, put select inside a loop


for{
	select{
	case value := <-ch1:
		fmt.Println("ch1", value)
	case value := <-ch2 :
		fmt.Println("ch2", value)
	}
}


LOOP -> SELECT -> HANDLE READY CHANNEL -> LOOP -> SELECT -> HANDLE NEXT VALUE

/// Example  " Continuously listen"


package main
import ("fmt" "time")

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			time.Sleep(1*time.Second)
			ch1 <- "New Order"
		}
	}()

	go func(){
		for {
			time.Sleep(2*time.Second)
			ch2 <- "New notification"
		}
	}()

	for {
		select {
		case message := <-ch1:
			fmt.Println(message)
		case message := <-ch2:
			fmt.Println(message)
		}
	}
}


/// Imagine a backedn service has

orders
payment
notifications
message


for {
	select {
	case order := <- orders:
		processOrder(order)
	case payment := <- payments:
		processPayment(payment)
	case notifications := <-notification:
		sendNotifications(notification)
	case message := <-messages:
		processMessage(message)
	}
}


// means : keep listening to all these channels and handle whichever one has data 


//************ The most important Syntax ***************

select {
case value := <-ch1:
	// ch1 received
case value := <-ch2:
	//ch2 received
default:
	// nothing is ready
}

AND

for {
	select {
	case value := <-ch1:
		// handle ch1
	case value := <-ch2:
		// handle ch2
	}
}


//******************** select vs WaitGroup *********************

// WaitGroup

wg.Wait() means Wait untill the goroutines are finished


// Channel

ch <- value means Send Data 

// Select

select {
case value := <-ch1:
case value := <-ch2
}

means wait for one of several channel operations 



SO

/*

WaitGroup => All workers finshed ?
Channel => Can i send/receive data?
Select => Which channel is ready ? 

*/


//**************** timeout ***************


// if you do not want your goroutines waiting forever
// you can use time.After()

select {
case result := <-ch:
	fmt.Println("Response : ", result)
case <- time.After(2*time.Second):
	fmt.Println("Request times out")
}



//********* Focus on these three patterns *********

select {
case value := <-ch1:
	fmt.Println(value)
case value := <-ch2:
	fmt.Println(value)
}


select {
case value := <-ch1:
	fmt.Println(value)
default:
	fmt.Println("Nothing avaiable")
}


AND

select {
case result := <-ch1:
	fmt.Println(result)
case <-time.After(2*time.Second):
	fmt.Println("Timeout")
}














