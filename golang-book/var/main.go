package main

import "fmt"

func main(){
	fmt.Println("Enter a number: ")
	var input float64
	fmt.Scan("%f", &input)

	output := input * 2

	fmt.Println(output)
}
