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