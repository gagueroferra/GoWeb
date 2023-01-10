package main

import "github.com/gin-gonic/gin"

/*Ejercicio 1 - Prueba de Ping
Vamos a crear una aplicación Web con el framework Gin que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON*/

func main() {
	//CREAR EL ROUTER

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Hola!")
	})
	router.Run()
}
