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
