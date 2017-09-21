package main 

import "fmt"

//Estructura
type Estudiante struct {
	Nombre string
	Apellido string
	Cedula string
	Notas []int
}
//Metodos de la estructura estudiante
func (this Estudiante) getNotas()([]int) {
	return this.Notas
}
//El * antes del nombre de la estructura hace refenencia al espacio de memoria que quiero modificar.
func (this *Estudiante) addNota(nota int){
	this.Notas = append(this.Notas, nota)
}

func main() {
	
	alumno := Estudiante {
		Nombre: "Antony",
		Apellido: "Carrizo",
		Cedula: "24623346",
		Notas: []int{},
	}
	alumno.addNota(20)

	fmt.Println(alumno)
	fmt.Println("Obtener notas de ",alumno.Nombre, " : ",alumno.getNotas() )
}