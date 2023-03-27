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
