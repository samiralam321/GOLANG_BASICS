//********************* Worker Pool **********************


/* The Problem : 

Imagine your backend receives 100 jobs :
One approach is 

for i:=1; i<=100; i++ {
    go processJob(i)
}

This creates 100 goroutines, which is obv unnecessary, instead we can create fix number of workers: 
Like : 100 jobs and 3 workers : Only 3 workers process jobs concurrently

This is Worker Pool


=> suppose a restaurent gets 100 orders
you don't hire 100 chefs

instead : 

100 orders -> kitchen queeu -> chef 1 chef 2 chef 3

whenever a chef finshes one order :
Chef 1 -> takes next order
Chef 2 -> takes next order
Chef 3 -> takes next order

this is exactly our workers pool does 


// a basic worker pool uses : Goroutines + Channel + WaitGroup*/

--------------------------------------------------------------------------

// job queue

jobs := make(chan int) 
this channel carries jobs 

// create a worker 
// a worker continuously receives jobs 

fucn worker(id int, jobs <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job : range jobs {
		fmt.Println("Worker", id, "processing job", job)
	}
}


// note that 

jobs <-chan int // this is receive-only ; the worker should only receive jobs 

for job := range jobs // keep receiving jobs until the jobs channel is closed

// so a worker basically says : 

give me a job 
process it
give me another job
process it
give me another job
process it 
.
.
.
.

channel closed
worker stops 


// complete basic worker pool 

package main
import ("fmt" "sync")

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id , "processing job", job)
	}
}

func main(){
	var wg sync.WaitGroup

	jobs := make(chan int) // create job channel  ; this is our communication pipe

	for i:=1; i<=3; i++ {  // create 3 workers
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i:=1; i<=10; i++ {   // send jobs into the job channel, workers takes them 
		jobs <- i
	}
	close(jobs) // closing the job channel, after sending all jobs , this mean there will be no more jobs 

	wg.Wait() // do not finish main() untill all workers have stopped

	fmt.Println("All jobs are completed")
}

//////////  Backend Example ///// image processing 

// imagine your backend receives image-processing jobs 

resize image 
compress image 
generate thumbnial 

you have 1000 images 
you could use 
100 goroutines 

or 10 workers ; these workers continously take jobs and process them 


package main
import ("fmt" "sync" "time")

func worker(id int, jobs <- chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for jos := range jobs {
		fmt.Println("Worker", id, "started job", job)
		time.Sleep(1*time.Second)

		fmt.Println("Worker", id, "completed job", job)
	}
}

func main(){
	var wg sync.WaitGroup

	jobs := make(chan int)

	for i:=1; i<=3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	for i:=1; i<=6; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs completed!!")
}



//*************** The most imp worker-pool patterns *************

jobs := make(chan int)

for i:=1; i<=2; i++ {   // it control the number of workers 
	wg.Add(1)
	go worker(i, jobs, &wg)
}

for i:=1; i<=10; i++ {
	jobs <- i
}

close(jobs)
wg.Wait()


// And the worker 

func worker(id int, jobs <- chan int, wg &sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// process the job 
	}
}






























