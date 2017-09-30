package controladores

import (
	"../db"
	"../modelos"
	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

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

	cursor, err := db.GetSessionDB().DB("boanergepro").Execute(query)
	if  err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	var resultado modelos.Usuario

	cursor.FetchOne(&resultado)
	contexto.StatusCode(iris.StatusOK)
	contexto.JSON(resultado)
}