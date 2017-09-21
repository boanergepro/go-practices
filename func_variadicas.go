package main 

import "fmt"

func ver (numeros ...int) {

	total := 0

	if numeros != nil {
		for _, num := range numeros {
			total += num
		}
		fmt.Println("La suma de los parametros pasados es:", total)
	} else {
		fmt.Println("No se ha proporcionado ningun parametro")
	}

}

func main () {
	fmt.Println("Funciones Variadicas.")
	ver()
}