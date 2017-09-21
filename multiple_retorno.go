package main 

import "fmt"

//Funcion retorna mas de un valor
func valores()(int, int) {
	return 5, 8
}

func main() {
	
	a, b := valores()
	fmt.Println("La funcion de retornno multiple, retorna a", a, "y",b)

	//Ignorando algun valor del retorno
	z,_ := valores()
	fmt.Println("Si ignoramos el segundo valor del retorno nos retorna a", z)

}