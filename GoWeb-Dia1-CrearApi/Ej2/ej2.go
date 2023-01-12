package main

import "github.com/gin-gonic/gin"

/*Ejercicio 2 - Manipulando el body

Vamos a crear un endpoint llamado /saludo. Con una pequeña estructura con nombre y apellido que al pegarle deberá responder en texto “Hola + nombre + apellido”

El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hola Andrea Rivas”
La estructura deberá ser como esta:
{
		“nombre”: “Andrea”,
		“apellido”: “Rivas”
}
*/

func main() {

	router := gin.Default()

	router.POST("/saludar", func(ctx *gin.Context) {
		type Persona struct {
			Nombre   string `json:"name"`
			Apellido string `json:"apellido"`
		}
		persona := Persona{
			Nombre:   "Gustavo",
			Apellido: "Aguero",
		}
		ctx.BindJSON(&persona)
		saludo := "Hola" + " " + persona.Nombre + " " + persona.Apellido
		ctx.JSON(200, gin.H{
			"message": saludo,
		})
	})
	router.Run()
}
