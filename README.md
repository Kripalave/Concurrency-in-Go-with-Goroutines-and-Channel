## Process
In computing, a process is an instance of a computer program that is being executed. A process simply having the following resources

    - Executable code
    - Operating system descriptors
    - Stack and heap memory
    - Security attributes
    - Context (Context is registers and physical memory addressing)
All records are kept stored in a block called Process Control Block(PCB). And many processes communicating with each other via an inter-process communication mechanism. A simple example is if you open python shell then it creates a process and now you able to visible python shell editor.


## Thread
Thread is simply a piece of code that will execute on the process or else you can say it process within the process is called a thread. It is also referred to as a “lightweight process”. A thread having its own

    - Thread ID
    - Program counter
    - Registers and stack
Threads communicate with each other using the Thread Control Block(TCB). An example of the thread is like you opened Python shell editor and now started to executing the python command like print(“Hello World”). So command execution is nothing but the execution of a thread.



What Is Concurrency?
“Concurrency is about dealing with lots of things at once. 
Parallelism is about doing lots of things at once.” — Rob Pike

“Share memory by communicating, don’t communicate by sharing memory.” - One of Go’s mottos


## Goroutines

A goroutine is a lightweight thread of execution in the Go programming language. It is similar to a thread in other programming languages, but it is managed by the Go runtime rather than the operating system. Goroutines allow concurrent execution of functions in a program, and they are designed to be efficient and scalable.

In Go, a program starts with a single goroutine, which executes the main function. Additional goroutines can be created using the go keyword followed by a function call. This starts a new goroutine that runs concurrently with the original goroutine.

Goroutines are far smaller that threads, they typically take around 2kB of stack space to initialize compared to a thread which takes 1Mb.

Goroutines are typically multiplexed onto a very small number of OS threads which typically mean concurrent go programs require far less resources in order to provide the same level of performance as languages such as Java. Creating a thousand goroutines would typically require one or two OS threads at most, whereas if we were to do the same thing in java it would require 1,000 full threads each taking a minimum of 1Mb of Heap space.


## Advantages of Goroutines

- Goroutines are cheaper than threads.
- Goroutine are stored in the stack and the size of the stack can grow and shrink according to the requirement of the program. But in threads, the size of the stack is fixed.
- Goroutines can communicate using the channel and these channels are specially designed to prevent race conditions when accessing shared memory using Goroutines.
- Suppose a program has one thread, and that thread has many Goroutines associated with it. If any of Goroutine blocks the thread due to resource requirement then all the remaining Goroutines will assign to a newly created OS thread. All these details are hidden from the programmers.

## Goroutines compared to Threads (Goroutines vs Threads)

- Goroutines are managed by the go runtime. 
- Goroutine are not hardware dependent.	
- Goroutines have easy communication medium known as channel.
- Goroutines are cheaper than threads.	
- They are cooperatively scheduled.
- They have faster startup time than threads.	
- Goroutine has growable segmented stacks.
- Goroutine does not have ID because go does not have Thread Local Storage.
- Due to the presence of channel one goroutine can communicate with other goroutine with low latency.


- Threads are hardware dependent.
- Operating system threads are managed by kernal.(OS threads are scheduled by the OS kernel)
- Thread does not have easy communication medium.
- The cost of threads are higher than goroutine.
- Threads have their own unique ID because they have Thread Local Storage.
- Due to lack of easy communication medium inter-threads communicate takes place with high latency.
- They are preemptively scheduled.
- They have slow startup time than goroutines.
- Threads does not have growable segmented stacks.
  

## How to create a Goroutine?

You can create your own Goroutine simply by using go keyword as a prefixing to the function or method call as shown in the below syntax:
- goroutines executes a function asynchronously
package main
import (
  "fmt"
  "time"
)

func myGoroutine() {
  fmt.Println("This is my first goroutine")
}

func main() {
  /*Then the main Goroutine terminated since there is no other code to execute and hence the myGoroutine() Goroutine did not get a chance to run.*/
  /* goroutines executes a function asynchronously */

  go myGoroutine() 
  
  fmt.Println("This is the main function")
}

Note: 
- When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call and any return values from the Goroutine are ignored.
- The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
- For any goroutines to run, the main() function must be defined and executed. When the main() function terminates, the program will be terminated, and no other goroutine will run.


## Using WaitGroup( we can fix above problem)

WaitGroup is a mechanism that can be used to syncronize the Go code. The basic usages in WaitGroup are:

Add() defines the number of goroutines that involved.
Wait() defines the wait condition in certain goroutine.
Done() defines the finish condition in certain goroutine. It means that the operation inside goroutine is finished.

### A WaitGroup has three exported methods that we make use of. These are −

1) Add(int) – Increases the counter.

2) Wait() – Blocks the execution until the internal counter becomes 0.

3) Done() – Decreases the counter by 1.

WaitGroups essentially allow us to tackle this problem by blocking until any goroutines within that WaitGroup have successfully executed.

We first call .Add(1) on our WaitGroup to set the number of goroutines we want to wait for, and subsequently, we call .Done() within any goroutine to signal the end of its' execution.

Note - You need to ensure that you call .Add(1) before you execute your goroutine.


Example:

package main

import (
    "fmt"
    "sync"
)

//initiate the waitgroup
var wg sync.WaitGroup

func main() {
    //add the 1 goroutine (in this case the odd() function is a goroutine)
    wg.Add(1)
    fmt.Println("this prints out")
    go odd()  //goroutine
    wg.Wait() //wait until odd() function is finished
}

//prints out odd numbers
func odd() {
    for i := 1; i < 10; i += 2 {
        fmt.Println(i)
    }
    //defines that the operation or job inside this function is finished
    wg.Done()
}

## Anonymous Goroutine Functions

Anonymous functions allow us to inline code by creating a function with no name. Combined with closures, this can simplify function signatures for goroutines.

anonymous goroutine is that it is simply an anonymous function that runs on a separate goroutine from which it was invoked.

Example:
package main
import (
   "fmt"
   "time"
)
func main() {
   fmt.Println("before the anonymous goroutine")
   go func(name string) {
      fmt.Println("Welcome to", name)
   }("TutorialsPoint")
   /* we have to once again block until our anonymous goroutine
     has finished or our main() function will complete without printing
    */
   time.Sleep(time.Second)
   fmt.Println("after the anonymous goroutine")
}



package main

import (
    "fmt"
    "sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
    fmt.Println("Inside my goroutine")
    waitgroup.Done()
}

func main() {
    fmt.Println("Hello World")

    var waitgroup sync.WaitGroup
    waitgroup.Add(1)
    go myFunc(&waitgroup)
    waitgroup.Wait()

    fmt.Println("Finished Execution")
}
## Race Condition
A race condition in Go occurs when two or more goroutines have shared data and interact with it simultaneously. 

func main(){
    m := make(map[int]struct{})
    wg := &sync.WaitGroup{}
    for i:=0;i<100;i++{
        wg.Add(1)
        go updatemap(wg,m,i)
    }
    wg.Wait()
}

func updatemap(wg *sync.WaitGroup,m map[int]struct{},r int){
    defer wg.Done()
    m[r] = struct{}{}
}
Note:
We have created a map and passed it to the updateMap function updating the key. Running the code gives the below output,  fatal error: concurrent map writes because maps are not thread-safe as mentioned earlier and this is a case of race condition. 

## Fixing Race Conditions

We have created our mutex and then in updateMap function we are doing Lock and Unlock. It’s not necessary to defer the unlock.

func main(){
    m := make(map[int]struct{})
   
    wg := &sync.WaitGroup{}
    mx := &sync.Mutex{}
    for i:=0;i<100;i++{
        wg.Add(1)
        go updatemap(wg,mx,m,i)
    }
    wg.Wait()

}

func updatemap(wg *sync.WaitGroup,mx *sync.Mutex,m map[int]struct{},r int){
    defer wg.Done()
    mx.Lock()
    defer mx.Unlock()
    m[r] = struct{}{}
}

We can create the mutex like this – mx := &sync.Mutex{} . Mutex provides 2 methods, Lock and Unlock . These methods will lock the code snippets in between so basically the code between them will always be executed in sync for this mutex.

##  Mutex 

Mutex or Mutual Exclusion is a mechanism that can be used to solve a race condition in a code. Mutex is a mechanism which allows only one goroutine is running the critical section (a code that has a potential of race condition) to prevent from race condition.

Mutex: The mutex guards a critical section by allowing a single thread an exclusive access (mutual exclusion).

A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.

A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time. This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex

- Lock
- Unlock

Any code present between a call to Lock and Unlock will be executed by only one Goroutine.

mutex.Lock() 

x = x + 1 // this statement be executed
          // by only one Goroutine 
          // at any point of time  

mutex.Unlock()


main.go
package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    balance int
)

func init() {
    balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
    balance += value
    mutex.Unlock()
    wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
    balance -= value
    mutex.Unlock()
    wg.Done()
}

func main() {
    fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
    go withdraw(700, &wg)
    go deposit(500, &wg)
    wg.Wait()

    fmt.Printf("New Balance %d\n", balance)
}


So, let’s break down what we have done here. Within both our deposit() and our withdraw() functions, we have specified the first step should be to acquire the mutex using the mutex.Lock() method.


## Using atomic

Atomic operations are used when we need to have a shared variable between different goroutines which will be updated by them. 

Atomic operations can be used to solve race condition. 


## Avoiding Deadlock

There a couple of scenarios that you need to be aware of when working with mutexes that will result in deadlock. Deadlock is a scenario within our code where nothing can progress due to every goroutine continually blocking when trying to attain a lock.

## Ensure You Call Unlock()!

if you are developing goroutines that require this lock and they can terminate in a number of different ways, then ensure that regardless of how your goroutine terminates, it always calls the Unlock() method.

If you fail to Unlock() on an error, then it is possible that your application will go into a deadlock as other goroutines will be unable to attain the lock on the mutex!


## Calling Lock() Twice
One example to keep in mind when developing with mutexes is that the Lock() method will block until it attains the lock. You need to ensure that when developing your applications you do not call the Lock() method twice on the same lock or else you will experience a deadlock scenario.

deadlock_example.go
package main

import (
	"fmt"
	"sync"
)
func main() {
	var b sync.Mutex
	b.Lock()
	b.Lock()
	fmt.Println("This never executes as we are in deadlock") 
}
output:
    go run deadlock_example.go
    fatal error: all goroutines are asleep - deadlock!
    goroutine 1 [semacquire]:


## Semaphore vs Mutex

Everything you can achieve with a Mutex can be done with a channel in Go if the size of the channel is set to 1.

However, the use case for what is known as a binary semaphore - a semaphore/channel of size 1 - is so common in the real world that it made sense to implement this exclusively in the form of a mutex.

- mutex: constrains access to a 1 single thread, to guard a critical section of code.
- semaphore: constrains access to at most N threads, to control/limit concurrent access to a shared resource.
Semaphore: The semaphore (a variable or an abstract data type) guards a critical section by allowing a N thread(s) an exclusive access (mutual exclusion).


## Channels
Channel is a pipeline for sending and receiving data. Think of it as a socket that runs inside your program. Channels provide a way for one goroutine to send structured data to another.


Working with Channels?

Channels are used when you want to transfer data between go routines

In Go language, a channel is a medium through which a goroutine communicates with another goroutine and this communication is lock-free. Or in other words, a channel is a technique which allows to let one goroutine to send data to another goroutine. By default channel is bidirectional, means the goroutines can send or receive data through the same channel.

Channels provide a way to send messages from one goroutine to another. Go channels work like sockets between goroutines within a single application. Like network sockets, they can be unidirectional or bidirectional. Channels can be short-lived or long-lived.

## Creating a Channel
    In Go language, a channel is created using chan keyword and it can only transfer data of the same type, different types of data are not allowed to transport from the same channel.

    - var Channel_name chan Type
    - channel_name:= make(chan Type)

Channels are goroutine-safe.
Channels can store and pass values between goroutines.
Channels provide FIFO semantics.
Channels cause goroutines to block and unblock, which we just learned about. 

## Buffered vs unbuffered

- Buffered channels have a limit make(chan int, 3), whereas unbuffered channels don't make(chan int)
- Buffered channels are a FIFO (first in first out) queue.
- Unbuffered channels are used for synchronous communication
- Unbuffered channels send data to other routines as soon as it receives it.
- Buffered channels have to be filled, and then the routines that consume the channel receive.
- Unbuffered Channel will block the goroutine whenever it is empty and waiting to be filled.
- Buffered Channel will also block the goroutine either when it is empty and waiting to be filled or it's on its full-capacity and there's a statement that want to fill the channel.

# Channel Buffering and Synchronization
Go provides functionality for channel buffering. To create a buffered channel, you’ll have to specify the buffer length as a second argument to the make function when declaring a channel.

channels := make(chan string, 2) // buffer capacity is 2



##Unbuffered channels

- Is the default channel type available in Go nd is the simplest one.
- An unbuffered channel has no capacity which means that only can handle one value at a time.
- A send operation on an unbuffered channel blocks the sending goroutine until another goroutine executes a receive on the same channel, now both goroutines may continue.
- If the receive operation was triggered first, the receiving goroutine is blocked until another goroutine performs a send on the same channel.

Example:
package main
import (
    "sync"
    "fmt"
    "time"
)
func main(){

    chanelex:= make(chan string)
    var wg sync.WaitGroup
    wg.Add(1)

    go func(){
        fmt.Println("Preparing to send")
        time.Sleep(3 * time.Second)
        chanelex <- "hello channel"
    }()
    go func(){

        fmt.Println("Waiting for a value")
        val := <-chanelex
        fmt.Println("Value is",val)
        wg.Done()
    }()
    wg.Wait()
}
//Output:
    Waiting for a value
    Preparing to send
    Value is hello channel


## Buffered Channel

Buffered Channel has a capacity to store messages inside it. Buffered Channel could be filled up to its defined capacity, not only one message.

- Using Buffered channel send multiple messages
- The buffer channel is a kind of message queue under the FIFO configuration, that means the firs message to go to the channel is the first one that will be read.

## How the buffered channels work:

- When the buffered channel is empty:
    Receiving messages is blocked until a message is sent across the channel.
- When the buffered channel is full:
    Sending messages in the channel is blocked until at last one channel is received from the channel, thus making space for new messages to be queued on the channel.
- When the buffered channel is almost filled, that is, neither full nor empty:
    Either sending or receiving messages on the channel is unblocked and the communication is instantaneous.



Example:
package main
import (
	"fmt"
	"time"
)

func access(ch chan int) {
	time.Sleep(time.Second)
	fmt.Println("start accessing channel\n")

	for i := range ch {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// only modify this line to defined the capacity
	ch := make(chan int, 3)
	defer close(ch)

	go access(ch)

	for i := 0; i < 9; i++ {
		ch <- i
		fmt.Println("Filled")
	}

	time.Sleep(3 * time.Second)
}


//Output:
Filled
Filled
Filled
start accessing channel
0
Filled
1
Filled
2

## Buffered vs Unbuffered 
Unbuffered Channel has no capacity initially, but Buffered Channel has a capacity.

Unbuffered Channel will block the goroutine whenever it is empty and waiting to be filled. But Buffered Channel will also block the goroutine either when it is empty and waiting to be filled or it's on its full-capacity and there's a statement that want to fill the channel.

## Capacity

unbuffered := make(chan int)

buffered := make(chan int, 30)


## For-Select blocks 

 - select {..} can be used to determine which channel's value needs to be received.
 - The for-select is used to coordinate channels and allow the communication each other. Help us to keep states between channels and goroutines and also could help us to handle with errors.

Example:
func main() {
    //create a channels
    frontend := make(chan string)
    backend := make(chan string)
    quit := make(chan string)

    //send some values to channels with goroutine
    go send(frontend, backend, quit)

    //receive some values from channels
    receive(frontend, backend, quit)
}

func send(f, b, q chan<- string) {
    data := []string{"React", "NodeJS", "Vue", "Flask", "Angular", "Laravel"}
    for i := 0; i < len(data); i++ {
        if i%2 == 0 {
            //send value to channel f
            f <- data[i]
        } else {
            //send value to channel b
            b <- data[i]
        }
    }
    //send value to channel q
    q <- "finished"
}

func receive(f, b, q <-chan string) {
    for {
        //using select to choose certain channel
        // that the value need to be received
        select {
        //if the value comes from channel called "f"
        //then execute the code
        case v := <-f:
            fmt.Println("Front End Dev:", v)
        //if the value comes from channel called "b"
        //then execute the code
        case v := <-b:
            fmt.Println("Back End Dev:", v)
        //if the value comes from channel called "q"
        //then execute the code
        case v := <-q:
            fmt.Println("This program is", v)
            return //finish the execution
        }
    }
}

1)select statement is a  blocking statement. This means that one case in the select statement will have to wait until one of the operations becomes unblocked.
2) If several cases can proceed, a single one case will be chosen to execute at random.
3)The default case will be executed if all cases are blocked.

package main
 
import (
   "fmt"
   "time"
)
 
func goroutine1(ch1 chan string) {
   time.Sleep(time.Second)
   for i := 0; i < 3; i++ {
       ch1 <- fmt.Sprintf("%d ==> Channel 1 message", i)
   }
}
 
func goroutine2(ch2 chan string) {
   for i := 0; i < 3; i++ {
       ch2 <- fmt.Sprintf("%d ==> Channel 2 message", i)
   }
 
}
 
func main() {
   ch1 := make(chan string)
   ch2 := make(chan string)
 
   go goroutine1(ch1)
   go goroutine2(ch2)
 
   for {
       time.Sleep(time.Second * 3)
       select {
 
       case value1 := <-ch1:
           fmt.Println(value1)
       case value2 := <-ch2:
           fmt.Println(value2)
       default:
           fmt.Println("All channels are blocking")
       }
   }
 
}




## Pipelines

Pipeline is a pattern where we breaking down complicated task to smaller sub tasks, and the output of first sub task, will be input for next sub task, we will repeat this process until all sub tasks has been completed.



Piplelie Idea is you can break a logical functionality into stages. Each stage does its own processing and passes the output to the next stage to get processed. You can modify stages independent of one another, rate limit the stages and so on and forth.

Example1:
package main

import "fmt"

func Multiply(value ,multiplier int) int {
	return value*multiplier
}
func Add(value,additive int) int {
	return value+additive
}

func main(){
	ints := []int{1,2,3,4}

	for _,v := range ints {
		fmt.Println(Multiply(Add(Multiply(v,2),1),2))
	}

}
/*6
10
14
18
*/


## Concurrent Pipelines
This simplistic model can now be extended to utilize go’s channels and goroutines to perform the processing of the stages concurrently. Before we can do that we must have following entities in our pipeline.

- Generator, which would be responsible for producing the input for pipeline processing. You think of this as first stage of the pipeline.
- Stages, where the actual processing can be performed.
- Canceller, a mechanism to signal cancellation or end of processing by pipeline.

Example:
package main

import "fmt"

func main(){

	done := make(chan interface{})
	intStream := generator(done, 1, 2, 3, 4)

	pipeline := add(done, multiply(done, intStream, 2), 5)

	for i := range pipeline {
		fmt.Println(i)
	}
}

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done: return
			case intStream <- i:
			}
		}

	}()
	return intStream
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) chan int {
	multipliedStream := make(chan int)
	go func() {
		defer close(multipliedStream)
		for i := range intStream {
			select {
			case <-done: return
			case multipliedStream <- i * multiplier:
			}
		}
	}()
	return multipliedStream
}


func add(done <-chan interface{}, intStream <-chan int, adder int) chan int {
	addedStream := make(chan int)
	go func() {
		defer close(addedStream)
		for i := range intStream {
			select {
			case <-done:return
			case addedStream <- i + adder:
			}
		}
	}()
	return addedStream
}

/*output:

7
9
11
13
 */


 ## Worker Pool

 A Worker pools are a “concurrency pattern” in which a fixed number of workers runs parallely in order to work in a number of task that are holding in a queue.

 In golang we use goroutines and channels to build this pattern. Usually will the workers be defined by a goroutine that is holding until it get data through a channel which is the responsible to coordinate the workers and the task in the queue (usually a buffered channel).


 ## without Worker pool

 package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)
type city struct {
	name string
	location string
}
func createCity(record city){
	time.Sleep(10 * time.Millisecond)
}
func main(){

	startTime:= time.Now();

	csvFile,err := os.Open("worldcities.csv")
	if err != nil{
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines,err := csv.NewReader(csvFile).ReadAll()
	if err != nil{
		fmt.Println(err)
	}
	counter := 0
	for _,line  := range csvLines{
		counter++
		createCity(city{
			name:line[0],
			location:line[9],

		})
		//fmt.Println(line[0])
		//fmt.Println(line[7])
	}


	fmt.Println("record saved:",counter)
	fmt.Println("total time:",time.Since(startTime))

}
/*
Output:
record saved: 12
total time: 143.387668ms
*/



## With Worker Pool it takes less time 

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

)
type city struct {
	name string
	location string
}
func createCity(record city){
	time.Sleep(10 * time.Millisecond)
}
func readData(cityChan chan []city){
	var cities []city
	csvFile,err := os.Open("worldcities.csv")
	if err != nil{
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines,err := csv.NewReader(csvFile).ReadAll()
	if err != nil{
		fmt.Println(err)
	}
	for _,line  := range csvLines{

		cities = append(cities,city{
			name:line[0],
			location:line[9],
		})

	}
	cityChan <- cities
}
func worker(cityChan chan city){
	for val := range cityChan{
		createCity(val)
	}
}
func main(){

	startTime:= time.Now();

	cities := make(chan []city)

	go readData(cities)

	const workers = 5

	jobs := make(chan city,1000)
	for w:= 1;w <= workers;w++{
		go worker(jobs)
	}
	counter := 0
	for _,val  := range <- cities{
		counter++
		jobs <- val
	}

	fmt.Println("record saved:",counter)
	fmt.Println("total time:",time.Since(startTime))

}
/*
Output:
record saved: 12
total time: 618.039µs
 */
