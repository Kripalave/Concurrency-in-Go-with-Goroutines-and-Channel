package main
import (
	"fmt"
	"sync"
)

var (
	mutex  sync.Mutex
	accountBalance int
)


func init() {
	accountBalance = 1000
}


func main() {

	fmt.Printf("initial balance %d\n", accountBalance)

	var wg sync.WaitGroup

	wg.Add(2)

	go withdraw(300, &wg)

	go deposit(500, &wg)

	wg.Wait()

	fmt.Printf("New balance %d\n", accountBalance)

}


func deposit(value int, wg *sync.WaitGroup) {

	mutex.Lock()

	fmt.Printf("Current balance: %d. Depositing %d\n", accountBalance, value)

	accountBalance += value

	mutex.Unlock()

	wg.Done()

}


func withdraw(value int, wg *sync.WaitGroup) {

	mutex.Lock()

	fmt.Printf("Current balance: %d. Withdrawing %d\n", accountBalance, value)

	accountBalance -= value

	mutex.Unlock()

	wg.Done()

}