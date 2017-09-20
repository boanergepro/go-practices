package main

import "fmt"

var(
	//Asignaciones multiples
	nombre = "Antony"
	apellido = "carrizo"

) 
func main() {
	
	fmt.Println(nombre)
	fmt.Println(apellido)
	//Recorrer un arreglo
	estudiantes := [] string {
		"israel",
		"daniel",
		"noel",
		"alirio",
		"jesus",
	}

	//Añadir un elementos a un slice
	estudiantes = append(estudiantes, "Antony")

	for key,value := range estudiantes {
		fmt.Println(key," --> ",value)
	}

	//Otra forma de usar el for
	for i := 0; i < len(estudiantes); i++ {
		fmt.Println(estudiantes[i])	
	}

	//map --> map[tipo]tipo
	var mapa map[string]string

	mapa = map[string]string{
		"edad":"21",
		"altura":"1.87",
	}
	if _, ok := mapa["edad"]; ok{
		fmt.Println(ok)
	}


}
