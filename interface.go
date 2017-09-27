package main 

import "fmt"

type Mesa struct {
	Nombre string
	NumPatas int8	
}

type Silla struct {
	Nombre string
	Material string
}
type Televisor struct {
	Nombre string
	NumPulgadas float32
}
type Pizarra struct {
	Nombre string
	Color string
}

type Objeto interface {
	GetNombre() string
}

func (this Mesa) GetNombre() string{
	return this.Nombre
}
func (this Silla) GetNombre() string{
	return this.Nombre
}
func (this Televisor) GetNombre() string{
	return this.Nombre
}
func (this Pizarra) GetNombre() string{
	return this.Nombre
}
func main () {

	var s_objetos[]Objeto

	s_objetos = make([]Objeto,0)

	mesa := Mesa{
		Nombre: "Mesa",
		NumPatas: 4,
	}
	
	silla := Silla{
		Nombre: "Silla",
		Material: "Madera",
	}

	televisor := Televisor{
		Nombre: "Sony",
		NumPulgadas: 20.5,
	}

	pizarra := Pizarra{
		Nombre: "Pizarra",
		Color: "Blanco Perla",
	}

	s_objetos = append(s_objetos, mesa,silla,televisor,pizarra)

	for _,value := range s_objetos {
		fmt.Println(value.GetNombre())
	}

}
