package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 2 : Creando un servidor
Vamos a levantar un servidor utilizando el paquete gin en el puerto 8080. Para probar nuestros endpoints haremos uso de postman.
Crear una ruta /ping que debe respondernos con un string que contenga pong con el status 200 OK.
Crear una ruta /products que nos devuelva la lista de todos los productos en la slice.
Crear una ruta /products/:id que nos devuelva un producto por su id.
Crear una ruta /products/search que nos permita buscar por parámetro los productos cuyo precio sean mayor a un valor priceGt.*/

func main() {

	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Pong")
	})
}
