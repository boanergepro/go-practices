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
	/*
	//VER
	app.Get("/api/usuarios/{key:string}", HandlerUser)

	//VER TODOS
	app.Get("/api/usuarios", HandlerAllUsers)
	
	//ACTUALIZAR
	app.Put("/api/usuarios/{key:string}", HandlerUpdateUser)
	
	//ELIMINAR
	app.Delete("/api/usuarios/{key:string}",HandlerDeleteUser)
	*/
    //Servidor corriendo en http://localhost:8080
    app.Run(iris.Addr(":8080"))
}

/*
func HandlerAllUsers (contexto context.Context) {

	query := arango.NewQuery(`
			FOR usuario in usuarios
			RETURN usuario
		`)

	cursor, err := session.DB("boanergepro").Execute(query)
	if  err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var resultados []Usuario

	cursor.FetchBatch(&resultados)
	contexto.StatusCode(iris.StatusOK)
	contexto.JSON(resultados)

}

func HandlerUser (contexto context.Context) {
	key_params := contexto.Params().Get("key")

	query := arango.NewQuery(`
			FOR usuario in usuarios
			FILTER usuario._key == @key
			RETURN usuario
		`)
	query.BindVars = map[string]interface{}{
		"key": key_params,
	}

	cursor, err := session.DB("boanergepro").Execute(query)
	if  err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var resultado Usuario

	cursor.FetchOne(&resultado)
	contexto.StatusCode(iris.StatusOK)
	contexto.JSON(resultado)
}

func HandlerUpdateUser (contexto context.Context) {
	key_params := contexto.Params().Get("key")

	query := arango.NewQuery(`
			FOR usuario in usuarios
			FILTER usuario._key == @key
			RETURN usuario
		`)
	query.BindVars = map[string]interface{}{
		"key": key_params,
	}

	cursor, err := session.DB("boanergepro").Execute(query)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var user_actualiar Usuario

	cursor.FetchOne(&user_actualiar)

	if contexto.PostValue("usuario") != "" {
		user_actualiar.Usuario = contexto.PostValue("usuario")
	}
	if contexto.PostValue("password") != "" {
		user_actualiar.Password = contexto.PostValue("password")
	}
	if contexto.PostValue("email") != "" {
		user_actualiar.Email = contexto.PostValue("email")
	}

	err = session.DB("boanergepro").Col("usuarios").Replace(user_actualiar.Key, user_actualiar)
	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}

func HandlerDeleteUser (contexto context.Context) {
	key_params := contexto.Params().Get("key")

	err = session.DB("boanergepro").Col("usuarios").Delete(key_params)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}
*/