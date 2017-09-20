package main 

import (
	"fmt"
	"reflect" //Manejar los tipos de datos.
)

func main() {

	nombre := "Antony"
	valorVariable := reflect.ValueOf(nombre).Kind() //Alcenar el tipo de dato
	fmt.Println("Tipo de dato", valorVariable)

	//Concatenaciones

	strin := fmt.Sprintf("Mi edad %d", 21)
	fmt.Println(strin)

}