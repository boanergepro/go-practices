package controladores

import (
	"../db"
	"../modelos"
	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

)

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

	cursor, err := db.GetSessionDB().DB("boanergepro").Execute(query)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var user_actualiar modelos.Usuario

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

	err = db.GetSessionDB().DB("boanergepro").Col("usuarios").Replace(user_actualiar.Key, user_actualiar)
	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}