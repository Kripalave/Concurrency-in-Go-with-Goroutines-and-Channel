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