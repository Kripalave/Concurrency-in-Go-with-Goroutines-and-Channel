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