package main

import (

	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/iris-contrib/middleware/cors"

)

type Usuario struct {
	arango.Document //Hay que incluirlo siempre en todas las estructuras porque añade campos nativos de arango
	Usuario     string `json:"usuario"`
	Password    string `json:"password"`
	Email 	  string `json:"email"`
}

func main() {

	app := iris.New()

	//Manejar la comunicacion del API
	app.WrapRouter(cors.WrapNext(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
			Debug:            true,
		}))
	
	//Connect(host, user, password string, log bool) (*Session, error) {
    session,err := arango.Connect("http://192.168.0.100:8529","boanergepro","123456",false) 
    if err != nil{
        panic(err)
    }

    //Creando una colleccion si no existe
    if !session.DB("boanergepro").ColExist("usuarios") {

    	documento := arango.NewCollectionOptions("usuarios", true)
    	session.DB("boanergepro").CreateCollection(documento)

    }

    //CREAR
	app.Post("/api/usuarios", func (contexto context.Context) {
		var user Usuario
		contexto.ReadJSON(&user)

    	//Insertar documento
    	err = session.DB("boanergepro").Col("usuarios").Save(&user)
    	if err != nil {
    		contexto.StatusCode(iris.StatusInternalServerError)
    	}

		contexto.StatusCode(iris.StatusOK)
	})

	//VER TODOS
	app.Get("/api/usuarios", func(contexto context.Context) {
		
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
		
	})

	//VER
	app.Get("/api/usuarios/{key:string}", func(contexto context.Context) {
		
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
		
	})
	
	//ACTUALIZAR
	app.Put("/api/usuarios/{key:string}", func(contexto context.Context) {

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
	})
	
	//ELIMINAR
	app.Delete("/api/usuarios/{key:string}", func(contexto context.Context) {

		key_params := contexto.Params().Get("key")

		err = session.DB("boanergepro").Col("usuarios").Delete(key_params)

	    if err != nil {
	    	contexto.StatusCode(iris.StatusInternalServerError)
	    }

	    contexto.StatusCode(iris.StatusOK)

	})

    //Servidor corriendo en http://localhost:8080
    app.Run(iris.Addr(":8080"))
}
