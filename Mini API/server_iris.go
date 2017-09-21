package main 

import (
	//"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)
//Estructuras
type Usuario struct {
	Usuario string
	Password string
	Email string
	Permisos []Permiso
}
type Permiso struct {
	Nombre string
}

//Inicializacion
var usuarios [] Usuario

func main () {

	usuarios = make([]Usuario,0)

	//Lenado de usuarios
	app := iris.New()
	
	app.RegisterView(iris.HTML("./", ".html"))
	

	app.Get("/registrar", func(contexto context.Context) {

		contexto.View("index.html")
	})

	app.Post("/registrar", func (contexto context.Context) {

		usuario := contexto.FormValue("usuario")
		password := contexto.FormValue("password")
		email := contexto.FormValue("email")
		permiso := contexto.FormValue("permisos")

		usuarios = append(usuarios, Usuario{ 
				Usuario: usuario, 
				Password: password,
				Email: email,
				Permisos: [] Permiso {
					Permiso {
						Nombre: permiso,
					},
				},
			})
		
		contexto.Redirect("/usuarios")
	})
	app.Get("/usuarios", func(contexto context.Context) {

		contexto.JSON(usuarios)
	})

	

	// Start the server using a network address.
    app.Run(iris.Addr(":8080"))
}
