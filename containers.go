package main

import "fmt"

//variables

var (

	number1 , number2 int = 20, 3; //multiple assignments
	b, c *int; //pointers
	d [1]int //array
)

//constants

const message string = "golang is amazing"

func main() {
	
	fmt.Println("Number 1 =", number1)
	fmt.Println("Number 2 =", number2)

	d[0] = 2017

	fmt.Println("array position 0 =", d[0])
}