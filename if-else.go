package main

import "fmt"

func main() {

	//basic

	if 7%2 == 0 {
		fmt.Println("7 es par")
	} else {
		fmt.Println("7 es impar")
	}

	//if without else

	if 8%4 == 0 {
		fmt.Println("8 es divisible entre 4")
	}

	//conditions with declarations

	if num := 9; num < 0 {
		
		fmt.Println(num, "es negativo")
	
	} else if num < 10 {

		fmt.Println(num, "solo tiene un digito")

	} else {
		fmt.Println(num, "tiene multiples digitos")
	}
	
}