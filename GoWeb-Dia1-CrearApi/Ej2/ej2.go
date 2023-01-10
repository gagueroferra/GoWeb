package main

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

import (
	"github.com/gin-gonic/gin"
)

type person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func main() {
	router := gin.Default()
	router.POST("/saludo", func(c *gin.Context) {
		persona := person{
			Name:     "Gustavo",
			LastName: "Aguero",
		}
		//Guardamos el json en la variable de tipo person persona
		c.BindJSON(&persona)
		response := "Hola" + " " + persona.Name + " " + persona.LastName
		//response as string
		//c.String(200, response)
		//response as json
		c.JSON(200, gin.H{"message": response})
	})
	err := router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
