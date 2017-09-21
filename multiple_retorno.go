package main 

import "fmt"

//Funcion retorna mas de un valor
func valores()(int, int) {
	return 5, 8
}

func main() {
	
	a, b := valores()
	fmt.Println("La funcion de retornno multiple, retorna a", a, b)

}