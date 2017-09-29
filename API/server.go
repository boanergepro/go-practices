package main

import (
	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"./controladores"
	"./db"
)

func main() {


	//Creando una colleccion si no existe
	if !db.GetSessionDB().DB("boanergepro").ColExist("usuarios") {

		documento := arango.NewCollectionOptions("usuarios", true)
		db.GetSessionDB().DB("boanergepro").CreateCollection(documento)

	}

	app := iris.New()

	//Manejar la comunicacion del API
	app.WrapRouter(cors.WrapNext(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
	}))

	//Index del API
	app.Get("/api", controladores.HandleIndex)

	//CREAR
	app.Post("/api/usuarios", controladores.HandlerCreateUser)
	
	//VER
	app.Get("/api/usuarios/{key:string}", controladores.HandlerUser)

	//VER TODOS
	app.Get("/api/usuarios", controladores.HandlerAllUsers)
	
	//ACTUALIZAR
	app.Put("/api/usuarios/{key:string}", controladores.HandlerUpdateUser)

	//ELIMINAR
	app.Delete("/api/usuarios/{key:string}", controladores.HandlerDeleteUser)

	//Servidor corriendo en http://localhost:8080
	app.Run(iris.Addr(":8080"))
}