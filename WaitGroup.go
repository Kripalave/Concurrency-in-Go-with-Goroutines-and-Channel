package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetch(waitgroup *sync.WaitGroup, URL string) {
	//defer waitgroup.Done()
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	waitgroup.Done()
	fmt.Printf("[ %s ] -> %d \n", URL, resp.StatusCode)
}

func main() {
	fmt.Println("Start consuming API endpoints")
	var waitgroup sync.WaitGroup

	URLs := []string{
		"https://www.golinuxcloud.com",
		"https://www.google.com",
		"https://go.dev",
		"https://aws.amazon.com",
		"https://www.microsoft.com",
	}

	waitgroup.Add(len(URLs))
	//if u waitgroup.Add(1) it fetch  https://www.golinuxcloud.com and terminates
	for _, URL := range URLs {
		//waitgroup.Add(1)

		go fetch(&waitgroup, URL)
	}
	waitgroup.Wait()


	fmt.Println("Done consuming API endpoints")

}
