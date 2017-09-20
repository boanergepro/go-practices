package main 

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	mapa := map[string] string{
		"Nombre": "ANTONY",
		"Apellido": "CARRIZO",
		"Edad": "21",
	}

	app := iris.New()

	app.Get("/", func(contexto context.Context) {
		contexto.JSON(mapa)
	})

	app.Run(iris.Addr(":8080"))
}