package controladores

import (
	"../db"
	"../modelos"
	arango "github.com/diegogub/aranGO"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

)

func HandlerAllUsers (contexto context.Context) {

	query := arango.NewQuery(`
			FOR usuario in usuarios
			RETURN usuario
		`)

	cursor,err := db.GetSessionDB().DB("boanergepro").Execute(query)
	if  err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}
	var resultados []modelos.Usuario

	cursor.FetchBatch(&resultados)
	contexto.StatusCode(iris.StatusOK)
	contexto.JSON(resultados)

}