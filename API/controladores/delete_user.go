package controladores

import (
	"../db"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

)

func HandlerDeleteUser (contexto context.Context) {
	key_params := contexto.Params().Get("key")

	err := db.GetSessionDB().DB("boanergepro").Col("usuarios").Delete(key_params)

	if err != nil {
		contexto.StatusCode(iris.StatusInternalServerError)
	}

	contexto.StatusCode(iris.StatusOK)
}