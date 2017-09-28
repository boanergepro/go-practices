package controladores

import (
	"../db"
	"../modelos"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"

)

func HandlerCreateUser (contexto context.Context) {
	var user modelos.Usuario
	contexto.ReadJSON(&user)

	//Insertar documento
	db.GetSessionDB().DB("boanergepro").Col("usuarios").Save(&user)

	contexto.StatusCode(iris.StatusOK)
}